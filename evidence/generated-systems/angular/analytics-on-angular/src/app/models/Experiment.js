
// Define collection and schema for Experiment
export interface Experiment {
    name:
	type : string
    objective:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    TrainingRuns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingRun' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    Notebooks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Notebook' }]
    Status:
 	type : String
#
    collection: 'experiments'
}
