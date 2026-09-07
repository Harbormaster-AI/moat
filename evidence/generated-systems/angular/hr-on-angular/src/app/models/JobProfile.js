
// Define collection and schema for JobProfile
export interface JobProfile {
    title:
	type : string
    jobCode:
	type : string
    JobFamily:
	type : Schema.Types.ObjectId
    Competencies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Competency' }]
    TrainingRecommendations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingCourse' }]
    Positions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    JobLevel:
 	type : String
    ExemptStatus:
 	type : String
#
    collection: 'jobProfiles'
}
