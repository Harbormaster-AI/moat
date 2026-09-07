
// Define collection and schema for Quote
export interface Quote {
    quoteNumber:
	type : string
    totalAmount:
	type : Money
    AircraftOrder:
	type : Schema.Types.ObjectId
#
    collection: 'quotes'
}
