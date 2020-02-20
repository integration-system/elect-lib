package blockchain

type BchConfig struct {
	Login   Login  `valid:"required~Required"`
	Address string `valid:"required~Required"`
}

type Login struct {
	Login    string `valid:"required~Required"`
	Password string `valid:"required~Required"`
}
