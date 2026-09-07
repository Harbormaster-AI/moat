
// Define collection and schema for TrainingEnrollment
export interface TrainingEnrollment {
    enrollmentNumber:
	type : string
    completionDate:
	type : Date
    score:
	type : String
    Course:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    Instructor:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'trainingEnrollments'
}
