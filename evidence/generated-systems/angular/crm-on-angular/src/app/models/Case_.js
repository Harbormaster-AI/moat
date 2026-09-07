
// Define collection and schema for Case_
export interface Case_ {
    caseNumber:
	type : string
    subject:
	type : string
    description:
	type : string
    slaDue:
	type : Date
    Organization:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Contact:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Team:
	type : Schema.Types.ObjectId
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    CaseComments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Note' }]
    Emails:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmailMessage' }]
    RelatedOpportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    Status:
 	type : String
    Priority:
 	type : String
    Origin:
 	type : String
    Severity:
 	type : String
#
    collection: 'case_s'
}
