
// Define collection and schema for ClinicalOrder
export interface ClinicalOrder {
    orderNumber:
	type : string
    Patient:
	type : Schema.Types.ObjectId
    Encounter:
	type : Schema.Types.ObjectId
    OrderingClinician:
	type : Schema.Types.ObjectId
    MedicationOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicationOrder' }]
    LaboratoryOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LaboratoryOrder' }]
    ImagingOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingOrder' }]
    ProcedureOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProcedureOrder' }]
    Authorizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Authorization' }]
    Status:
 	type : String
    OrderType:
 	type : String
    Priority:
 	type : String
#
    collection: 'clinicalOrders'
}
