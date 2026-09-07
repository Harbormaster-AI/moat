
// Define collection and schema for Operator
export interface Operator {
    name:
	type : string
    icaoDesignator:
	type : string
    AircraftOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftOrder' }]
    OperatedAircraft:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Aircraft' }]
    SalesRegion:
	type : Schema.Types.ObjectId
    OperatorType:
 	type : String
#
    collection: 'operators'
}
