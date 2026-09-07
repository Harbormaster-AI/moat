
// Define collection and schema for PerformanceReview
export interface PerformanceReview {
    reviewNumber:
	type : string
    reviewDate:
	type : Date
    reviewerComments:
	type : string
    Employee:
	type : Schema.Types.ObjectId
    Reviewer:
	type : Schema.Types.ObjectId
    Cycle:
	type : Schema.Types.ObjectId
    CompetencyRatings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CompetencyRating' }]
    Goals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Goal' }]
    Rating:
 	type : String
    Status:
 	type : String
#
    collection: 'performanceReviews'
}
