
// Define collection and schema for Matter
export interface Matter {
    matterName:
	type : string
    leadCounsel:
	type : string
    LegalHolds:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LegalHold' }]
    Organization:
	type : Schema.Types.ObjectId
    DataBreaches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataBreach' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contract' }]
    MatterType:
 	type : String
    Status:
 	type : String
#
    collection: 'matters'
}
