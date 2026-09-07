
// Define collection and schema for LabResult
export interface LabResult {
    resultCode:
	type : string
    issuedDate:
	type : Date
    LaboratoryOrder:
	type : Schema.Types.ObjectId
    Observations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Observation' }]
    Laboratory:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'labResults'
}
