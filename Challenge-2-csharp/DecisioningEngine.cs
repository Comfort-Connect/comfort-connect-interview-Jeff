// Challenge 2: Partner Decisioning Engine
//
// Implement DecisioningEngine.Evaluate so it:
//  1. Filters partners whose rules the application satisfies:
//     - CreditScore >= MinCreditScore
//     - LoanAmount  <= MaxLoanAmount
//     - EquipmentType is in SupportedEquipment
//     - State is NOT in ExcludedStates
//  2. Calculates monthly payment using amortization:
//       M = P * [r(1+r)^n] / [(1+r)^n - 1]
//     where P = LoanAmount, r = APR/12/100, n = TermMonths
//  3. Returns offers sorted by APR ascending, then MonthlyPayment ascending
//  4. Returns an empty list if no partner is eligible
//  5. Throws ArgumentException for invalid applications
//     (CreditScore <= 0, LoanAmount <= 0, empty EquipmentType/State)

// --- Inline test harness (top-level statements must come first) ---

var rules = new List<PartnerRule>
{
    new("premier",  "Premier Program",          680, 50_000m, new[] { "HVAC", "Plumbing", "Generator" },           Array.Empty<string>(),  5.99m, 60, "0% for first 12 months"),
    new("partner1", "Partner 1 Finance",        600, 30_000m, new[] { "HVAC", "Plumbing" },                        new[] { "CA", "NY" },   7.49m, 48, null),
    new("partner2", "Partner 2 Leasing",        550, 25_000m, new[] { "Solar" },                                   Array.Empty<string>(),  9.99m, 36, "No credit check required"),
    new("partner3", "Partner 3 Financial",      650, 40_000m, new[] { "HVAC", "Plumbing", "Generator", "Solar" },  new[] { "FL" },         6.49m, 72, "First payment deferred 90 days"),
    new("partner4", "Partner 4 Energy Finance", 700, 60_000m, new[] { "HVAC", "Solar", "Generator" },              Array.Empty<string>(),  4.99m, 60, "Energy efficiency bonus"),
};

var engine = new DecisioningEngine(rules);

var goodCredit = new Application("APP-001", 720, 15_000m, "HVAC", "MD");
var offers = engine.Evaluate(goodCredit);
Console.WriteLine($"Good credit HVAC: {offers.Count} offers (expected 4)");
foreach (var o in offers)
{
    Console.WriteLine($"  {o.PartnerId}: APR={o.Apr}%, {o.TermMonths}mo, ${o.MonthlyPayment}/mo");
}

var lowCredit = new Application("APP-002", 580, 8_000m, "Plumbing", "TX");
Console.WriteLine($"\nLow credit: {engine.Evaluate(lowCredit).Count} offers (expected 0)");

// --- Types ---

public record Application(
    string ApplicationId,
    int CreditScore,
    decimal LoanAmount,
    string EquipmentType, // "HVAC", "Plumbing", "Generator", "Solar"
    string State          // US state abbreviation
);

public record Offer(
    string PartnerId,
    string PartnerName,
    decimal Apr,
    int TermMonths,
    decimal MonthlyPayment,
    decimal TotalCost,
    string? PromoLabel    // null if no promo applies
);

public record PartnerRule(
    string PartnerId,
    string PartnerName,
    int MinCreditScore,
    decimal MaxLoanAmount,
    IReadOnlyList<string> SupportedEquipment,
    IReadOnlyList<string> ExcludedStates,
    decimal BaseApr,
    int DefaultTermMonths,
    string? PromoLabel
);

public class DecisioningEngine
{
    private readonly IReadOnlyList<PartnerRule> _rules;

    public DecisioningEngine(IReadOnlyList<PartnerRule> rules)
    {
        _rules = rules ?? throw new ArgumentNullException(nameof(rules));
    }

    public IReadOnlyList<Offer> Evaluate(Application application)
    {
        // TODO: implement decisioning logic
        return Array.Empty<Offer>();
    }
}
