
// Define collection and schema for WorkAuthorization
export interface WorkAuthorization {
    country:
	type : string
    expirationDate:
	type : Date
    Employee:
	type : Schema.Types.ObjectId
    Documents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Document' }]
    Status:
 	type : String
#
    collection: 'workAuthorizations'
}
