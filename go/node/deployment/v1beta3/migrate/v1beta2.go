package migrate

import (
	"github.com/akash-network/akash-api/go/node/deployment/v1beta2"
	"github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	amigrate "github.com/akash-network/akash-api/go/node/types/v1beta3/migrate"
)

func ResourceUnitFromV1Beta2(id uint32, from v1beta2.Resource) v1beta3.ResourceUnit {
	return v1beta3.ResourceUnit{
		Resources: amigrate.ResourcesFromV1Beta2(id, from.Resources),
		Count:     from.Count,
		Price:     from.Price,
	}
}

func ResourcesUnitsFromV1Beta2(from []v1beta2.Resource) v1beta3.ResourceUnits {
	res := make(v1beta3.ResourceUnits, 0, len(from))

	for i, oval := range from {
		res = append(res, ResourceUnitFromV1Beta2(uint32(i+1), oval)) // nolint: gosec
	}

	return res
}

func GroupIDFromV1Beta2(from v1beta2.GroupID) v1beta3.GroupID {
	return v1beta3.GroupID{
		Owner: from.Owner,
		DSeq:  from.DSeq,
		GSeq:  from.GSeq,
	}
}

func GroupSpecFromV1Beta2(from v1beta2.GroupSpec) v1beta3.GroupSpec {
	return v1beta3.GroupSpec{
		Name:         from.Name,
		Requirements: amigrate.PlacementRequirementsFromV1Beta2(from.Requirements),
		Resources:    ResourcesUnitsFromV1Beta2(from.Resources),
	}
}

func GroupFromV1Beta2(from v1beta2.Group) v1beta3.Group {
	return v1beta3.Group{
		GroupID:   GroupIDFromV1Beta2(from.GroupID),
		State:     v1beta3.Group_State(from.State),
		GroupSpec: GroupSpecFromV1Beta2(from.GroupSpec),
		CreatedAt: from.CreatedAt,
	}
}
