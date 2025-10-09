package wallet

// FeeStrategy calculates a fee for a given amount.
type FeeStrategy interface {
	Calculate(amount float64) float64
}

// BitcoinFeeStrategy charges 2%.
type BitcoinFeeStrategy struct{}

// EthereumFeeStrategy charges 1%.
type EthereumFeeStrategy struct{}

// Calculate implements FeeStrategy for Bitcoin fees.
func (BitcoinFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.02 // 2%
}

// Calculate implements FeeStrategy for Ethereum fees.
func (EthereumFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.01 // 1%
}
