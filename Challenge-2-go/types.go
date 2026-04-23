package decisioning

// Application is a homeowner's financing application.
type Application struct {
	ApplicationID string
	CreditScore   int
	LoanAmount    float64
	EquipmentType string // "HVAC", "Plumbing", "Generator", "Solar"
	State         string // US state abbreviation
}

// Offer is a financing offer from a lending partner.
type Offer struct {
	PartnerID      string
	PartnerName    string
	APR            float64
	TermMonths     int
	MonthlyPayment float64
	TotalCost      float64
	PromoLabel     string // empty string if no promo applies
}

// PartnerRule is the configuration for a lending partner's eligibility and pricing rules.
type PartnerRule struct {
	PartnerID          string
	PartnerName        string
	MinCreditScore     int
	MaxLoanAmount      float64
	SupportedEquipment []string
	ExcludedStates     []string
	BaseAPR            float64
	DefaultTermMonths  int
	PromoLabel         string
}
