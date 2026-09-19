# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)


5.times do |i|
  Insurer.create( name:"test string for name", legalName:"test string for legalName", domicileCountry:"test string for domicileCountry", naicNumber:"test string for naicNumber", website:"test string for website" )
  InsuranceProduct.create( name:"test string for name", productCode:"test string for productCode", LineOfBusiness:0 )
  CoverageDefinition.create( name:"test string for name", defaultLimit:"test value", defaultDeductible:"test value", asMandatory:true, CoverageType:0 )
  Distributor.create( name:"test string for name", licenseNumber:"test string for licenseNumber", region:"test string for region", DistributorType:0 )
  Agent.create( firstName:"test string for firstName", lastName:"test string for lastName", licenseId:"test string for licenseId", Status:0 )
  Customer.create( firstName:"test string for firstName", lastName:"test string for lastName", organizationName:"test string for organizationName", taxId:"test string for taxId", dateOfBirth:1.week.ago, primaryAddress:"test value", CustomerType:0 )
  Application.create( applicationNumber:"test string for applicationNumber", submissionDate:1.week.ago, Status:0 )
  Quote.create( quoteNumber:"test string for quoteNumber", totalPremium:"test value", ratingDate:1.week.ago, asBound:true )
  UnderwritingDecision.create( notes:"test string for notes", decisionDate:1.week.ago, Decision:0 )
  Underwriter.create( firstName:"test string for firstName", lastName:"test string for lastName", employeeId:"test string for employeeId", authorityLimit:"test value" )
  Policy.create( policyNumber:"test value", effectivePeriod:1.week.ago, totalPremium:"test value", Status:0, PaymentPlan:0 )
  Endorsement.create( endorsementNumber:"test string for endorsementNumber", effectiveDate:1.week.ago, description:"test string for description" )
  PolicyCoverage.create( limit:"test value", deductible:"test value", premium:"test value", CoverageType:0 )
  InsuredObject.create( description:"test string for description", serialOrId:"test string for serialOrId", primaryAddress:"test value", ObjectType:0 )
  Beneficiary.create( name:"test string for name", share:"test value", Relationship:0 )
  BillingAccount.create( accountNumber:"test string for accountNumber", balance:"test value", Status:0 )
  Invoice.create( invoiceNumber:"test string for invoiceNumber", dueDate:1.week.ago, totalDue:"test value", Status:0 )
  Payment.create( paymentReference:"test string for paymentReference", amount:"test value", paymentDate:1.week.ago, Method:0, Status:0 )
  Claim.create( claimNumber:"test value", noticeDate:1.week.ago, lossDate:1.week.ago, reportedBy:"test string for reportedBy", Status:0, LossCause:0 )
  Incident.create( location:"test value", description:"test string for description", IncidentType:0 )
  Exposure.create( ExposureType:0, Status:0 )
  Adjuster.create( firstName:"test string for firstName", lastName:"test string for lastName", licenseNumber:"test string for licenseNumber", AdjusterType:0 )
  ClaimReserve.create( amount:"test value", setDate:1.week.ago, ReserveType:0, Status:0 )
  ClaimPayment.create( paymentNumber:"test string for paymentNumber", amount:"test value", paymentDate:1.week.ago, PayeeType:0, Method:0, Status:0 )
  ServiceProvider.create( name:"test string for name", taxId:"test string for taxId", ProviderType:0, NetworkStatus:0 )
  ReinsuranceAgreement.create( agreementNumber:"test string for agreementNumber", effectivePeriod:1.week.ago, retention:"test value", limit:"test value", cessionPercentage:"test value", ReinsuranceType:0, TreatyType:0 )
  SubrogationRecovery.create( recoveryReference:"test string for recoveryReference", amount:"test value", recoveryDate:1.week.ago, Status:0 )
  ThirdParty.create( name:"test string for name", taxId:"test string for taxId", address:"test value", PartyType:0 )
  Document.create( fileName:"test string for fileName", uploadedDate:1.week.ago, DocumentType:0 )
end
