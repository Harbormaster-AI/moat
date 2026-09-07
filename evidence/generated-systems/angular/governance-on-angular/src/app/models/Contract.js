
// Define collection and schema for Contract
export interface Contract {
    title:
	type : string
    effectiveDate:
	type : Date
    expiryDate:
	type : Date
    repositoryUrl:
	type : URL
    ThirdParty:
	type : Schema.Types.ObjectId
    Obligations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Obligation' }]
    DataProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    Matter:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'contracts'
}
