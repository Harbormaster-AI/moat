
// Define collection and schema for Activity
export interface Activity {
    subject:
	type : string
    dueDate:
	type : Date
    startAt:
	type : Date
    endAt:
	type : Date
    location:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Contact:
	type : Schema.Types.ObjectId
    Lead:
	type : Schema.Types.ObjectId
    Opportunity:
	type : Schema.Types.ObjectId
    Case:
	type : Schema.Types.ObjectId
    Campaign:
	type : Schema.Types.ObjectId
    ActivityType:
 	type : String
    Status:
 	type : String
    Priority:
 	type : String
#
    collection: 'activitys'
}
