
// Define collection and schema for Competency
export interface Competency {
    name:
	type : string
    category:
	type : string
    JobProfiles:
 	type : [{ type: Schema.Types.ObjectId, ref: 'JobProfile' }]
    CompetencyRatings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CompetencyRating' }]
#
    collection: 'competencys'
}
