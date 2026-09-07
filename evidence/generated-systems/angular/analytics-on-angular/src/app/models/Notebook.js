
// Define collection and schema for Notebook
export interface Notebook {
    title:
	type : string
    repository:
	type : RepositoryRef
    Workspace:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Experiments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Experiment' }]
    Queries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BIQuery' }]
    Language:
 	type : String
#
    collection: 'notebooks'
}
