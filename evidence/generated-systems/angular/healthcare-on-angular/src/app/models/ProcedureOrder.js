
// Define collection and schema for ProcedureOrder
export interface ProcedureOrder {
    procedureCode:
	type : string
    consentObtained:
	type : boolean
    Order:
	type : Schema.Types.ObjectId
    Facility:
	type : Schema.Types.ObjectId
    Procedure:
	type : Schema.Types.ObjectId
    AnesthesiaType:
 	type : String
#
    collection: 'procedureOrders'
}
