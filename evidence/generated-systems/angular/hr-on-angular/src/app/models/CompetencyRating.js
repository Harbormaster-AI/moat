
// Define collection and schema for CompetencyRating
export interface CompetencyRating {
    comment:
	type : string
    Review:
	type : Schema.Types.ObjectId
    Competency:
	type : Schema.Types.ObjectId
    Rating:
 	type : String
#
    collection: 'competencyRatings'
}
