# Direct-to-Consumer Financing Flow

Comfort Connect is a fintech platform that connects homeowners with financing options for home comfort equipment (HVAC, plumbing, generators) through a network of lending partners. Today, all applications originate through a dealer/contractor portal. We want to build a Direct-to-Consumer (D2C) channel where homeowners can visit a website, apply for financing on their own, and get matched with both a financing offer and a local contractor for installation.

## Requirements

Design a system that supports the following end-to-end flow:

1. A homeowner submits personal and project information (name, address, SSN, loan amount, equipment type) through a web application.
2. The system performs a soft credit pull against one or more credit bureaus to assess eligibility without impacting the applicant's credit score.
3. Based on the credit profile, the system runs decisioning logic to determine eligibility across multiple lending partners — one of which is an internal "Premier Program" that has its own pricing engine (rate sheets, buydown tiers, promotional terms). The system should return ranked offers from all eligible partners.
4. Once the homeowner selects an offer, the system must allocate a qualified contractor based on the homeowner's location, the contractor's service area, availability, and certification status, and create a dispatch/appointment for the installation.

## Design Considerations

Your design should address:

- **API boundaries** between these stages
- **Partner-specific decisioning rules** that change frequently
- **Sensitive PII/credit data** management
- **Contractor matching performance** at scale across thousands of contractors and service territories

  