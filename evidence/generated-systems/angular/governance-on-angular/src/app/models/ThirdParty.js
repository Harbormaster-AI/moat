
// Define collection and schema for ThirdParty
export interface ThirdParty {
    name:
	type : string
    country:
	type : string
    contactEmail:
	type : EmailAddress
    Organization:
	type : Schema.Types.ObjectId
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    Assessments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ThirdPartyAssessment' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contract' }]
    Obligations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Obligation' }]
    DataBreaches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataBreach' }]
    ThirdPartyType:
 	type : String
    Criticality:
 	type : String
#
    collection: 'thirdPartys'
}
