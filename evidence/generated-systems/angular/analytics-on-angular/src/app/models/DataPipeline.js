
// Define collection and schema for DataPipeline
export interface DataPipeline {
    name:
	type : string
    schedule:
	type : CronSchedule
    Workspace:
	type : Schema.Types.ObjectId
    Tasks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataTask' }]
    Sources:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSource' }]
    Outputs:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    LineageNode:
	type : Schema.Types.ObjectId
    TriggerType:
 	type : String
    Status:
 	type : String
#
    collection: 'dataPipelines'
}
