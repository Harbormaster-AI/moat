
// Define collection and schema for Laboratory
export interface Laboratory {
    name:
	type : string
    cliaNumber:
	type : string
    Facility:
	type : Schema.Types.ObjectId
    LaboratoryOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LaboratoryOrder' }]
    LabResults:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LabResult' }]
#
    collection: 'laboratorys'
}
