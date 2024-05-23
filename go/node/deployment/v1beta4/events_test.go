package v1beta4

import (
	"crypto/sha256"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "pkg.akt.dev/go/node/deployment/v1"
	"pkg.akt.dev/go/sdkutil"
)

var (
	keyAcc, _         = sdk.AccAddressFromBech32("akash1qtqpdszzakz7ugkey7ka2cmss95z26ygar2mgr")
	errWildcard       = errors.New("wildcard string error can't be matched")
	tmpSum            = sha256.Sum256([]byte(keyAcc))
	deploymentVersion = encodeHex(tmpSum[:])
)

type testEventParsing struct {
	msg    sdkutil.Event
	expErr error
}

func (tep testEventParsing) testMessageType() func(t *testing.T) {
	_, err := ParseEvent(tep.msg)
	return func(t *testing.T) {
		// if the error expected is errWildcard to catch untyped errors, don't fail the test, the error was expected.
		if errors.Is(tep.expErr, errWildcard) {
			require.Error(t, err)
		} else {
			require.Equal(t, tep.expErr, err)
		}
	}
}

var TEPS = []testEventParsing{
	{
		msg: sdkutil.Event{
			Type: "nil",
		},
		expErr: sdkutil.ErrUnknownType,
	},
	{
		msg: sdkutil.Event{
			Type: sdkutil.EventTypeMessage,
		},
		expErr: sdkutil.ErrUnknownModule,
	},

	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
		},
		expErr: sdkutil.ErrUnknownAction,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: "nil",
		},
		expErr: sdkutil.ErrUnknownModule,
	},

	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: "nil",
		},
		expErr: sdkutil.ErrUnknownAction,
	},

	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentCreated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
				{
					Key:   evVersionKey,
					Value: string(deploymentVersion),
				},
			},
		},
		expErr: nil,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentCreated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "abc",
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentCreated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentCreated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
			},
		},
		expErr: errWildcard,
	},

	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentUpdated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
				{
					Key:   evVersionKey,
					Value: string(deploymentVersion),
				},
			},
		},
		expErr: nil,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentUpdated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionGroupClosed,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
				{
					Key:   evGSeqKey,
					Value: "1",
				},
			},
		},
		expErr: nil,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentClosed,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
			},
		},
		expErr: nil,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentClosed,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "abc",
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionGroupClosed,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionGroupClosed,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
				{
					Key:   evGSeqKey,
					Value: "1",
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentUpdated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: "neh",
				},
				{
					Key:   evDSeqKey,
					Value: "5",
				},
				{
					Key:   evVersionKey,
					Value: string(deploymentVersion),
				},
			},
		},
		expErr: errWildcard,
	},
	{
		msg: sdkutil.Event{
			Type:   sdkutil.EventTypeMessage,
			Module: v1.ModuleName,
			Action: evActionDeploymentUpdated,
			Attributes: []sdk.Attribute{
				{
					Key:   evOwnerKey,
					Value: keyAcc.String(),
				},
			},
		},
		expErr: errWildcard,
	},
}

func TestEventParsing(t *testing.T) {
	for i, test := range TEPS {
		t.Run(strconv.Itoa(i), test.testMessageType())
	}
}

func TestVersionEncoding(t *testing.T) {
	versionHex := encodeHex(tmpSum[:])
	assert.Len(t, versionHex, encodedVersionHexLen)
	decodedVersion, err := decodeHex(versionHex)
	assert.NoError(t, err)
	assert.Equal(t, tmpSum[:], decodedVersion)
}
