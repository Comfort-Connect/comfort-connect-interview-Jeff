package decisioning

import "testing"

var rules = []PartnerRule{
	{
		PartnerID: "premier", PartnerName: "Premier Program",
		MinCreditScore: 680, MaxLoanAmount: 50000,
		SupportedEquipment: []string{"HVAC", "Plumbing", "Generator"},
		BaseAPR:            5.99, DefaultTermMonths: 60,
		PromoLabel: "0% for first 12 months",
	},
	{
		PartnerID: "partner1", PartnerName: "Partner 1 Finance",
		MinCreditScore: 600, MaxLoanAmount: 30000,
		SupportedEquipment: []string{"HVAC", "Plumbing"},
		ExcludedStates:     []string{"CA", "NY"},
		BaseAPR:            7.49, DefaultTermMonths: 48,
	},
	{
		PartnerID: "partner2", PartnerName: "Partner 2 Leasing",
		MinCreditScore: 550, MaxLoanAmount: 25000,
		SupportedEquipment: []string{"Solar"},
		BaseAPR:            9.99, DefaultTermMonths: 36,
		PromoLabel: "No credit check required",
	},
	{
		PartnerID: "partner3", PartnerName: "Partner 3 Financial",
		MinCreditScore: 650, MaxLoanAmount: 40000,
		SupportedEquipment: []string{"HVAC", "Plumbing", "Generator", "Solar"},
		ExcludedStates:     []string{"FL"},
		BaseAPR:            6.49, DefaultTermMonths: 72,
		PromoLabel: "First payment deferred 90 days",
	},
	{
		PartnerID: "partner4", PartnerName: "Partner 4 Energy Finance",
		MinCreditScore: 700, MaxLoanAmount: 60000,
		SupportedEquipment: []string{"HVAC", "Solar", "Generator"},
		BaseAPR:            4.99, DefaultTermMonths: 60,
		PromoLabel: "Energy efficiency bonus",
	},
}

// Credit score 720, HVAC, $15k loan, MD
// Expected: 4 offers (partner4, premier, partner3, partner1) — NOT partner2 (Solar only)
func TestGoodCreditHVAC(t *testing.T) {
	app := Application{ApplicationID: "APP-001", CreditScore: 720, LoanAmount: 15000.0, EquipmentType: "HVAC", State: "MD"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(offers) != 4 {
		t.Fatalf("should return 4 offers — got %d", len(offers))
	}
	if offers[0].PartnerID != "partner4" {
		t.Errorf("first offer should be partner4 (lowest APR) — got %q", offers[0].PartnerID)
	}
	if offers[1].PartnerID != "premier" {
		t.Errorf("second offer should be premier — got %q", offers[1].PartnerID)
	}
}

// Credit score 580 — below all partner minimums.
func TestLowCreditNoOffers(t *testing.T) {
	app := Application{ApplicationID: "APP-002", CreditScore: 580, LoanAmount: 8000.0, EquipmentType: "Plumbing", State: "TX"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(offers) != 0 {
		t.Errorf("should return 0 offers for low credit — got %d", len(offers))
	}
}

// Solar equipment, credit 560. Expected: only partner2.
func TestSolarOnlyPartner2(t *testing.T) {
	app := Application{ApplicationID: "APP-003", CreditScore: 560, LoanAmount: 10000.0, EquipmentType: "Solar", State: "MD"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(offers) != 1 {
		t.Fatalf("should return 1 offer — got %d", len(offers))
	}
	if offers[0].PartnerID != "partner2" {
		t.Errorf("should be partner2 — got %q", offers[0].PartnerID)
	}
}

// Applicant in California — partner1 excludes CA.
func TestExcludedState(t *testing.T) {
	app := Application{ApplicationID: "APP-004", CreditScore: 620, LoanAmount: 15000.0, EquipmentType: "HVAC", State: "CA"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, o := range offers {
		if o.PartnerID == "partner1" {
			t.Fatalf("partner1 should be excluded for CA applicants")
		}
	}
}

// Loan amount $55k — exceeds every partner's max except partner4 ($60k).
func TestLoanAmountExceedsMax(t *testing.T) {
	app := Application{ApplicationID: "APP-005", CreditScore: 720, LoanAmount: 55000.0, EquipmentType: "HVAC", State: "MD"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(offers) != 1 {
		t.Fatalf("only partner4 should accept $55k loan — got %d offers", len(offers))
	}
	if offers[0].PartnerID != "partner4" {
		t.Errorf("should be partner4 — got %q", offers[0].PartnerID)
	}
}

// Zero-value Application has invalid fields (credit 0, loan 0) and should error.
func TestInvalidApplicationErrors(t *testing.T) {
	if _, err := Evaluate(rules, Application{}); err == nil {
		t.Fatalf("should have returned an error for zero-value application")
	}
}

// Verify offers are ranked by APR ascending.
func TestOffersRankedByApr(t *testing.T) {
	app := Application{ApplicationID: "APP-006", CreditScore: 750, LoanAmount: 20000.0, EquipmentType: "HVAC", State: "MD"}

	offers, err := Evaluate(rules, app)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 1; i < len(offers); i++ {
		if offers[i].APR < offers[i-1].APR {
			t.Fatalf("offers not sorted by APR: %v should be <= %v", offers[i-1].APR, offers[i].APR)
		}
	}
}
