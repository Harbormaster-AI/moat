
// Define collection and schema for LaboratoryOrder
export interface LaboratoryOrder {
    testCode:
	type : string
    fastingRequired:
	type : boolean
    Order:
	type : Schema.Types.ObjectId
    Laboratory:
	type : Schema.Types.ObjectId
    Results:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LabResult' }]
    SpecimenType:
 	type : String
#
    collection: 'laboratoryOrders'
}
