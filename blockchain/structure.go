package blockchain

type BlockchainConfig struct {
	Login           BchLogin `valid:"required~Required"`
	Address         string   `valid:"required~Required"`
	VotingEventBody interface{}
}

type BchLogin struct {
	Login    string `valid:"required~Required"`
	Password string `valid:"required~Required"`
}
