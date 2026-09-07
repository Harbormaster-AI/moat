
// Define collection and schema for AircraftOrder
export interface AircraftOrder {
    orderNumber:
	type : string
    totalAmount:
	type : Money
    Operator:
	type : Schema.Types.ObjectId
    Variant:
	type : Schema.Types.ObjectId
    Quote:
	type : Schema.Types.ObjectId
    PurchaseAgreement:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'aircraftOrders'
}
