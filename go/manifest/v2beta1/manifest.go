package v2beta1

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	k8svalidation "k8s.io/apimachinery/pkg/util/validation"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta2"
	types "github.com/akash-network/akash-api/go/node/types/v1beta2"
)

var (
	serviceNameValidationRegex = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	hostnameMaxLen             = 255
)

// Manifest store list of groups
type Manifest []Group

// GetGroups returns a manifest with groups list
func (m Manifest) GetGroups() []Group {
	return m
}

// ValidateManifest does validation for manifest
func ValidateManifest(m Manifest) error {
	if len(m) == 0 {
		return fmt.Errorf("%w: manifest is empty", ErrInvalidManifest)
	}
	return validateManifestGroups(m.GetGroups())
}

type validateManifestGroupsHelper struct {
	hostnames          map[string]int // used as a set
	globalServiceCount uint
}

func validateManifestGroups(groups []Group) error {
	helper := validateManifestGroupsHelper{
		hostnames: make(map[string]int),
	}
	names := make(map[string]int) // used as a set
	for _, group := range groups {
		if err := validateManifestGroup(group, &helper); err != nil {
			return err
		}
		if _, exists := names[group.GetName()]; exists {
			return fmt.Errorf("%w: duplicate group %q", ErrInvalidManifest, group.GetName())
		}

		names[group.GetName()] = 0 // Value stored is not used
	}
	if helper.globalServiceCount == 0 {
		return fmt.Errorf("%w: zero global services", ErrInvalidManifest)
	}
	return nil
}

func validateManifestGroup(group Group, helper *validateManifestGroupsHelper) error {
	if 0 == len(group.Services) {
		return fmt.Errorf("%w: group %q contains no services", ErrInvalidManifest, group.GetName())
	}

	if err := dtypes.ValidateResourceList(&group); err != nil {
		return err
	}
	for _, s := range group.Services {
		if err := validateManifestService(s, helper); err != nil {
			return err
		}
	}
	return nil
}

func validateManifestService(service Service, helper *validateManifestGroupsHelper) error {
	if len(service.Name) == 0 {
		return fmt.Errorf("%w: service name is empty", ErrInvalidManifest)
	}

	serviceNameValid := serviceNameValidationRegex.MatchString(service.Name)
	if !serviceNameValid {
		return fmt.Errorf("%w: service %q name is invalid", ErrInvalidManifest, service.Name)
	}

	if len(service.Image) == 0 {
		return fmt.Errorf("%w: service %q has empty image name", ErrInvalidManifest, service.Name)
	}

	for _, envVar := range service.Env {
		idx := strings.Index(envVar, "=")
		if idx == 0 {
			return fmt.Errorf("%w: service %q defines an env. var. with an empty name", ErrInvalidManifest, service.Name)
		}

		var envVarName string
		if idx > 0 {
			envVarName = envVar[0:idx]
		} else {
			envVarName = envVar
		}

		if 0 != len(k8svalidation.IsEnvVarName(envVarName)) {
			return fmt.Errorf("%w: service %q defines an env. var. with an invalid name %q", ErrInvalidManifest, service.Name, envVarName)
		}

	}

	for _, serviceExpose := range service.Expose {
		if err := validateServiceExpose(service.Name, serviceExpose, helper); err != nil {
			return err
		}
	}

	return nil
}

func validateServiceExpose(serviceName string, serviceExpose ServiceExpose, helper *validateManifestGroupsHelper) error {
	if serviceExpose.Port == 0 || serviceExpose.Port > math.MaxUint16 {
		return fmt.Errorf("%w: service %q port value must be 0 < value <= 65535 ", ErrInvalidManifest, serviceName)
	}

	switch serviceExpose.Proto {
	case TCP, UDP:
		break
	default:
		return fmt.Errorf("%w: service %q protocol %q unknown", ErrInvalidManifest, serviceName, serviceExpose.Proto)
	}

	if serviceExpose.Global {
		helper.globalServiceCount++
	}

	for _, host := range serviceExpose.Hosts {
		if !isValidHostname(host) {
			return fmt.Errorf("%w: service %q has invalid hostname %q", ErrInvalidManifest, serviceName, host)
		}

		_, exists := helper.hostnames[host]
		if exists {
			return errors.Errorf("hostname %q is duplicated, this is not allowed", host)
		}
		helper.hostnames[host] = 0 // Value stored does not matter
	}

	return nil
}

func isValidHostname(hostname string) bool {
	return len(hostname) <= hostnameMaxLen && 0 == len(k8svalidation.IsDNS1123Subdomain(hostname))
}

func ValidateManifestWithGroupSpecs(m *Manifest, gspecs []*dtypes.GroupSpec) error {
	rlists := make([]types.ResourceGroup, 0, len(gspecs))
	for _, gspec := range gspecs {
		rlists = append(rlists, gspec)
	}
	return validateManifestDeploymentGroups(m.GetGroups(), rlists)
}

func ValidateManifestWithDeployment(m *Manifest, dgroups []dtypes.Group) error {
	rgroups := make([]types.ResourceGroup, 0, len(dgroups))
	for _, dgroup := range dgroups {
		rgroups = append(rgroups, dgroup)
	}

	return validateManifestDeploymentGroups(m.GetGroups(), rgroups)
}

func validateManifestDeploymentGroups(mgroups []Group, dgroups []types.ResourceGroup) error {
	if len(mgroups) != len(dgroups) {
		return errors.Errorf("invalid manifest: group count mismatch (%v != %v)", len(mgroups), len(dgroups))
	}

	dgroupByName := make(map[string]types.ResourceGroup)

	for _, dgroup := range dgroups {
		dgroupByName[dgroup.GetName()] = dgroup
	}

	for _, mgroup := range mgroups {
		dgroup, dgroupExists := dgroupByName[mgroup.GetName()]

		if !dgroupExists {
			return errors.Errorf("invalid manifest: unknown deployment group ('%v')", mgroup.GetName())
		}

		if err := validateManifestDeploymentGroup(mgroup, dgroup); err != nil {
			return err
		}
	}

	return nil
}

func validateManifestDeploymentGroup(mgroup Group, dgroup types.ResourceGroup) error {
	mlist := make([]types.Resources, len(mgroup.GetResources()))
	copy(mlist, mgroup.GetResources())

	httpOnlyEndpointsCountForDeploymentGroup := 0
	otherEndpointsCountForDeploymentGroup := 0

	// Iterate over all deployment groups
deploymentGroupLoop:
	for _, drec := range dgroup.GetResources() {
		for _, endpoint := range drec.Resources.Endpoints {
			switch endpoint.Kind {
			case types.Endpoint_SHARED_HTTP:
				httpOnlyEndpointsCountForDeploymentGroup++
			case types.Endpoint_RANDOM_PORT:
				otherEndpointsCountForDeploymentGroup++
			}
		}
		// Find a matching manifest group
		for idx := range mlist {
			mrec := mlist[idx]

			// Check that this manifest group is not yet exhausted
			if mrec.Count == 0 {
				continue
			}

			if !drec.Resources.CPU.Equal(mrec.Resources.CPU) ||
				!drec.Resources.Memory.Equal(mrec.Resources.Memory) ||
				!drec.Resources.Storage.Equal(mrec.Resources.Storage) {
				continue
			}

			// If the manifest group contains more resources than the deployment group, then
			// fulfill the deployment group entirely
			if mrec.Count >= drec.Count {
				mrec.Count -= drec.Count
				drec.Count = 0
			} else {
				// Partially fulfill the deployment group since the manifest group contains less
				drec.Count -= mrec.Count
				mrec.Count = 0
			}

			// Update the value stored in the list
			mlist[idx] = mrec

			// If the deployment group is fulfilled then break out and
			// move to the next deployment
			if drec.Count == 0 {
				continue deploymentGroupLoop
			}
		}
		// If this point is reached then the deployment group cannot be fully matched
		// against the given manifest groups
		return fmt.Errorf("%w: underutilized deployment group %q", ErrManifestCrossValidation, dgroup.GetName())
	}

	// Search for any manifest groups which are not fully satisfied
	for _, mrec := range mlist {
		if mrec.Count > 0 {
			return fmt.Errorf("%w: manifest resources %q is not fully matched with deployment groups", ErrManifestCrossValidation, mgroup.GetName())
		}
	}

	httpOnlyEndpointCount := 0
	otherEndpointCount := 0

	for _, service := range mgroup.Services {
		for _, serviceExpose := range service.Expose {
			if serviceExpose.Global {
				if IsIngress(serviceExpose) {
					httpOnlyEndpointCount++
				} else {
					otherEndpointCount++
				}
			}
		}
	}

	if otherEndpointCount != otherEndpointsCountForDeploymentGroup {
		return errors.Errorf("invalid manifest: mismatch on number of endpoints %d != %d", otherEndpointCount, otherEndpointsCountForDeploymentGroup)
	}

	if httpOnlyEndpointCount != httpOnlyEndpointsCountForDeploymentGroup {
		return errors.Errorf("invalid manifest: mismatch on number of HTTP only endpoints %d != %d", httpOnlyEndpointCount, httpOnlyEndpointsCountForDeploymentGroup)
	}

	return nil
}

func IsIngress(expose ServiceExpose) bool {
	return expose.Proto == TCP && expose.Global && 80 == ExposeExternalPort(expose)
}

func ExposeExternalPort(expose ServiceExpose) int32 {
	if expose.ExternalPort == 0 {
		return int32(expose.Port) // nolint: gosec
	}
	return int32(expose.ExternalPort) // nolint: gosec
}
