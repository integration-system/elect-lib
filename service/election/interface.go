package election

type Service interface {
	GetCurrentRegElectionsByIdList(IdentityListRequest) ([]Election, error)
	GetCurrentVotingElectionsByIdList(IdentityListRequest) ([]Election, error)
}
