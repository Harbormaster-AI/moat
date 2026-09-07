
// Define collection and schema for RecommendationScenario
export interface RecommendationScenario {
    name:
	type : string
    objective:
	type : string
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Experiments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Experiment' }]
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Alert' }]
    RecommendationType:
 	type : String
#
    collection: 'recommendationScenarios'
}
