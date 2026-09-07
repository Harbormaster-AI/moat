
// Define collection and schema for CarePlan
export interface CarePlan {
    planNumber:
	type : string
    goalSummary:
	type : string
    Patient:
	type : Schema.Types.ObjectId
    Encounters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Encounter' }]
    Tasks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CareTask' }]
    CareTeam:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'carePlans'
}
