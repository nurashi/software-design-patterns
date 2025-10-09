package wallet

// FeeStrategy defines interface for transaction fee calculation
type FeeStrategy interface {
	Calculate(amount float64) float64
}

// BitcoinFeeStrategy applies BTC-specific fees
type BitcoinFeeStrategy struct{}

func (b BitcoinFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.02 // 2% fee
}

// EthereumFeeStrategy applies ETH-specific fees
type EthereumFeeStrategy struct{}

func (e EthereumFeeStrategy) Calculate(amount float64) float64 {
	return amount * 0.015 // 1.5% fee
}
