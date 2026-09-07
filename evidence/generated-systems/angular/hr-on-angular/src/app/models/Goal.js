
// Define collection and schema for Goal
export interface Goal {
    title:
	type : string
    description:
	type : string
    targetDate:
	type : Date
    weight:
	type : Percentage
    Employee:
	type : Schema.Types.ObjectId
    Cycle:
	type : Schema.Types.ObjectId
    ParentGoal:
	type : Schema.Types.ObjectId
    ChildGoals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Goal' }]
    Status:
 	type : String
#
    collection: 'goals'
}
