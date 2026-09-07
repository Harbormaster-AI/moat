
// Define collection and schema for OnboardingTask
export interface OnboardingTask {
    taskNumber:
	type : string
    name:
	type : string
    dueDate:
	type : Date
    Employee:
	type : Schema.Types.ObjectId
    AssignedTo:
	type : Schema.Types.ObjectId
    Dependencies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OnboardingTask' }]
    RelatedOffer:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'onboardingTasks'
}
