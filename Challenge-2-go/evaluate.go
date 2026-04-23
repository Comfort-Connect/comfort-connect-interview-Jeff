package decisioning

// Evaluate scores an application against all partner rules and returns eligible
// offers ranked by APR (ascending), then monthly payment as tiebreaker.
//
// Returns a non-nil error if the application has invalid fields
// (credit score <= 0, loan amount <= 0, empty equipment type or state).
//
// TODO: Implement the following:
//  1. For each partner, check eligibility:
//     - Credit score >= partner's minimum
//     - Loan amount <= partner's maximum
//     - Equipment type is in partner's supported list
//     - Applicant's state is NOT in partner's excluded states
//  2. For eligible partners, calculate monthly payment using amortization:
//     M = P * [r(1+r)^n] / [(1+r)^n - 1]
//     where P = loan amount, r = monthly rate (APR/12/100), n = term in months
//  3. Return offers sorted by: lowest APR first, then lowest monthly payment as tiebreaker
//  4. Return an empty slice if no partner is eligible
func Evaluate(rules []PartnerRule, app Application) ([]Offer, error) {
	// TODO: implement decisioning logic
	return []Offer{}, nil
}
