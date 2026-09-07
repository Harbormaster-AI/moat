
// Define collection and schema for DataSource
export interface DataSource {
    name:
	type : string
    connection:
	type : ConnectionInfo
    Streaming:
	type : boolean
    Workspace:
	type : Schema.Types.ObjectId
    ProducedDatasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Pipelines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataPipeline' }]
    SourceType:
 	type : String
    Format:
 	type : String
#
    collection: 'dataSources'
}
