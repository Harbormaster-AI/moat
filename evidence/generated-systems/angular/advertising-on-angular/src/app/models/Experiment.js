
// Define collection and schema for Experiment
export interface Experiment {
    name:
	type : string
    hypothesis:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    Campaign:
	type : Schema.Types.ObjectId
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ExperimentVariant' }]
    Status:
 	type : String
#
    collection: 'experiments'
}
