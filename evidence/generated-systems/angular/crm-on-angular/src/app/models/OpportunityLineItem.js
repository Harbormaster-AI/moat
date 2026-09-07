
// Define collection and schema for OpportunityLineItem
export interface OpportunityLineItem {
    quantity:
	type : String
    unitPrice:
	type : Money
    discountPercent:
	type : String
    totalPrice:
	type : Money
    Opportunity:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
    PriceBookEntry:
	type : Schema.Types.ObjectId
#
    collection: 'opportunityLineItems'
}
