
// Define collection and schema for ProductionOrder
export interface ProductionOrder {
    orderNumber:
	type : string
    Variant:
	type : Schema.Types.ObjectId
    Plant:
	type : Schema.Types.ObjectId
    AircraftOrder:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'productionOrders'
}
