package wallet

type FeeStrategy interface {
	Calculate(amount float64) float64
}

type BitcoinFeeStrategy struct{}
type EthereumFeeStrategy struct{}



// gives ablity to chose according to COIN
func (BitcoinFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.02 // 2%
}

func (EthereumFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.01 // 1%
}
