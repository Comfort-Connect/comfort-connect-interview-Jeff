# Challenge 2: Partner Decisioning

## Context

When a homeowner applies for financing, Comfort Connect evaluates the application against multiple lending partner rule sets and returns ranked, eligible offers. Each partner has different eligibility criteria (min credit score, max loan amount, supported equipment types) and different pricing (base rate, term options, promotional adjustments).

Partner rules change frequently (monthly rate sheet updates, new promos, new partners onboarded). The logic must be config-driven, not hardcoded.

## Task

Implement `Evaluate` so that it:

1. **Takes a slice of partner rules** (`[]PartnerRule`) and an `Application` — the test file builds the rule set directly in code; no JSON parsing required
2. **Filters** to the partners whose rules the application satisfies
3. **Returns eligible offers** ranked by lowest APR, then lowest monthly payment as tiebreaker
4. **Calculates monthly payment** using standard amortization: `P * [r(1+r)^n] / [(1+r)^n - 1]` where P = principal, r = monthly rate, n = term months
5. **Handles edge cases**: no eligible partners (empty slice), invalid application fields (return an error)

## Types

### `evaluate.go`
```go
func Evaluate(rules []PartnerRule, app Application) ([]Offer, error)
```

### `application.go`
```go
type Application struct {
    ApplicationID string
    CreditScore   int
    LoanAmount    float64
    EquipmentType string // "HVAC", "Plumbing", "Generator", "Solar"
    State         string // US state abbreviation
}
```

### `offer.go`
```go
type Offer struct {
    PartnerID      string
    PartnerName    string
    APR            float64
    TermMonths     int
    MonthlyPayment float64
    TotalCost      float64
    PromoLabel     string // empty string if no promo applies
}
```

## Expected Behavior

```go
app := Application{ApplicationID: "APP-001", CreditScore: 720, LoanAmount: 15000.0, EquipmentType: "HVAC", State: "MD"}
offers, _ := Evaluate(rules, app)

// Should return offers from eligible partners, sorted by APR:
// 1. Partner4 (4.99%, 60mo, $283.07/mo) — HVAC eligible, credit >= 700
// 2. Premier (5.99%, 60mo, $289.99/mo) — HVAC eligible, credit >= 680
// 3. Partner3 (6.49%, 72mo, $252.39/mo) — HVAC eligible, credit >= 650
// 4. Partner1 (7.49%, 48mo, $362.69/mo) — HVAC eligible, credit >= 600
// NOT: Partner2 — only supports Solar equipment

lowCredit := Application{ApplicationID: "APP-002", CreditScore: 580, LoanAmount: 8000.0, EquipmentType: "Plumbing", State: "TX"}
offers2, _ := Evaluate(rules, lowCredit)
// Should return empty slice — no partners accept credit < 600
```

## Provided Files

- Starter source files at the package root (`application.go`, `offer.go`, `partner_rule.go`, `evaluate.go`)
- Test file `evaluate_test.go` — builds the partner rule slice in code and calls `Evaluate`

## Running

No external libraries required — Go standard library only.

```bash
cd Challenge-2-go
go test ./...
```

## Evaluation Criteria

- Correct eligibility filtering
- Accurate monthly payment calculation
- Proper ranking
- Clean separation of concerns
- Error handling for bad input / malformed rules
- The provided test cases should pass
