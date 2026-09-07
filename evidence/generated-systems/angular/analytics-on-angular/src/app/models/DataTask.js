
// Define collection and schema for DataTask
export interface DataTask {
    name:
	type : string
    command:
	type : string
    retries:
	type : number
    Pipeline:
	type : Schema.Types.ObjectId
    InputDatasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    OutputDatasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    TaskType:
 	type : String
#
    collection: 'dataTasks'
}
