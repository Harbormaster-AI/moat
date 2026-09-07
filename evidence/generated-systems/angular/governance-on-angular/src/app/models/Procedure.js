
// Define collection and schema for Procedure
export interface Procedure {
    title:
	type : string
    versionLabel:
	type : string
    Policy:
	type : Schema.Types.ObjectId
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Status:
 	type : String
#
    collection: 'procedures'
}
