package election

type Service interface {
	GetCurrentRegElectionsByIdList(IdentityListRequest) ([]Election, error)
	GetCurrentVotingElectionsByIdList(IdentityListRequest) ([]Election, error)
	GetRegElectionById(IdentityRequest) (*Election, error)
	GetVotingElectionById(IdentityRequest) (*Election, error)
	GetElectionById(IdentityRequest) (*Election, error)
}
