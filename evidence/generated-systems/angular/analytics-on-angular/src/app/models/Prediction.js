
// Define collection and schema for Prediction
export interface Prediction {
    referenceKey:
	type : string
    predictedAt:
	type : Date
    score:
	type : String
    Endpoint:
	type : Schema.Types.ObjectId
    ModelVersion:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
#
    collection: 'predictions'
}
