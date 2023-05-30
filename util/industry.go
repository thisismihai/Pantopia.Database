package util

// Constants for all supported currencies
const (
	EDUCATION    = "EDUCATION"
	CONSTRUCTION = "CONSTRUCTION"
	HEALTHCARE   = "HEALTHCARE"
	FINANCE      = "FINANCE"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedIndustry(industry string) bool {
	switch industry {
	case EDUCATION, CONSTRUCTION, HEALTHCARE, FINANCE:
		return true
	}
	return false
}
