
// Define collection and schema for OpportunityStageHistory
export interface OpportunityStageHistory {
    changedAt:
	type : Date
    comment:
	type : string
    Opportunity:
	type : Schema.Types.ObjectId
    ChangedBy:
	type : Schema.Types.ObjectId
    FromStage:
 	type : String
    ToStage:
 	type : String
#
    collection: 'opportunityStageHistorys'
}
