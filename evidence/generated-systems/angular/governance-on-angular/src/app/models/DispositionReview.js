
// Define collection and schema for DispositionReview
export interface DispositionReview {
    reviewDate:
	type : Date
    reviewer:
	type : string
    notes:
	type : string
    Record:
	type : Schema.Types.ObjectId
    RetentionSchedule:
	type : Schema.Types.ObjectId
    Outcome:
 	type : String
#
    collection: 'dispositionReviews'
}
