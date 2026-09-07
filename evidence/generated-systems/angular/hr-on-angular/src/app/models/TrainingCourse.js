
// Define collection and schema for TrainingCourse
export interface TrainingCourse {
    code:
	type : string
    title:
	type : string
    durationHours:
	type : String
    Prerequisites:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingCourse' }]
    Enrollments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingEnrollment' }]
    JobProfiles:
 	type : [{ type: Schema.Types.ObjectId, ref: 'JobProfile' }]
    DeliveryMethod:
 	type : String
#
    collection: 'trainingCourses'
}
