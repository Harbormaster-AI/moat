
// Define collection and schema for InferenceEndpoint
export interface InferenceEndpoint {
    name:
	type : string
    endpointUrl:
	type : string
    trafficShare:
	type : Percentage
    ModelVersion:
	type : Schema.Types.ObjectId
    Workspace:
	type : Schema.Types.ObjectId
    Predictions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Prediction' }]
    Mode:
 	type : String
#
    collection: 'inferenceEndpoints'
}
