# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)


5.times do |i|
  Organization.create( name:"test string for name", defaultCurrency:"test string for defaultCurrency", defaultLocale:"test value", website:"test value" )
  User.create( username:"test string for username", fullName:"test string for fullName", email:"test value", locale:"test value", Role:0, Status:0 )
  Team.create( name:"test string for name", TeamType:0 )
  Territory.create( name:"test string for name", region:"test string for region", TerritoryType:0 )
  Account.create( name:"test string for name", accountNumber:"test string for accountNumber", industry:"test string for industry", billingAddress:"test value", shippingAddress:"test value", website:"test value", phone:"test value", asActive:true, AccountType:0, LifecycleStage:0 )
  Contact.create( firstName:"test string for firstName", lastName:"test string for lastName", title:"test string for title", email:"test value", phone:"test value", mobile:"test value", mailingAddress:"test value", PreferredContactMethod:0 )
  Lead.create( firstName:"test string for firstName", lastName:"test string for lastName", company:"test string for company", email:"test value", phone:"test value", converted:true, Status:0, Source:0, Rating:0 )
  Opportunity.create( name:"test string for name", amount:"test value", closeDate:1.week.ago, probability:"test value", description:"test string for description", Stage:0, Type:0, ForecastCategory:0 )
  OpportunityLineItem.create( quantity:"test value", unitPrice:"test value", discountPercent:"test value", totalPrice:"test value" )
  OpportunityStageHistory.create( changedAt:1.week.ago, comment:"test string for comment", FromStage:0, ToStage:0 )
  Product.create( sku:"test string for sku", name:"test string for name", asActive:true, standardPrice:"test value", description:"test string for description", ProductType:0, Uom:0 )
  PriceBook.create( name:"test string for name", asActive:true, description:"test string for description" )
  PriceBookEntry.create( unitPrice:"test value", effectiveDate:1.week.ago, expirationDate:1.week.ago, asActive:true )
  Quote.create( quoteNumber:"test string for quoteNumber", validityStart:1.week.ago, validityEnd:1.week.ago, totalAmount:"test value", discountPercent:"test value", taxAmount:"test value", shippingAmount:"test value", Status:0 )
  QuoteLineItem.create( quantity:"test value", unitPrice:"test value", discountAmount:"test value", taxAmount:"test value", totalAmount:"test value" )
  Order.create( orderNumber:"test string for orderNumber", orderDate:1.week.ago, totalAmount:"test value", taxAmount:"test value", shippingAmount:"test value", Status:0 )
  OrderItem.create( quantity:"test value", unitPrice:"test value", discountAmount:"test value", taxAmount:"test value", totalAmount:"test value" )
  Contract.create( contractNumber:"test string for contractNumber", startDate:1.week.ago, endDate:1.week.ago, renewalTermMonths:100, autoRenew:true, Status:0 )
  Case_.create( caseNumber:"test string for caseNumber", subject:"test string for subject", description:"test string for description", slaDue:1.week.ago, Status:0, Priority:0, Origin:0, Severity:0 )
  Activity.create( subject:"test string for subject", dueDate:1.week.ago, startAt:1.week.ago, endAt:1.week.ago, location:"test string for location", ActivityType:0, Status:0, Priority:0 )
  Campaign.create( name:"test string for name", startDate:1.week.ago, endDate:1.week.ago, budget:"test value", actualCost:"test value", expectedRevenue:"test value", Status:0, Type:0 )
  CampaignMember.create( responded:true, Status:0, MemberType:0 )
  Note.create( title:"test string for title", content:"test string for content", createdAt:1.week.ago, updatedAt:1.week.ago )
  EmailMessage.create( subject:"test string for subject", body:"test string for body", sentAt:1.week.ago, messageId:"test string for messageId", Direction:0, Status:0 )
end
