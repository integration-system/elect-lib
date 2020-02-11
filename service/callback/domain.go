package callback

import (
	"github.com/integration-system/elect-lib/service/registry"
)

type VoterChangedEvent struct {
	FromState registry.VoterState
	ToState   registry.VoterState
	Voter     registry.Voter
	Citizen   registry.Citizen
}
