
// Define collection and schema for DataSet
export interface DataSet {
    name:
	type : string
    schemaVersion:
	type : string
    refreshSchedule:
	type : CronSchedule
    Sensitive:
	type : boolean
    Workspace:
	type : Schema.Types.ObjectId
    Sources:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSource' }]
    Pipelines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataPipeline' }]
    SemanticModels:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SemanticModel' }]
    Dimensions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dimension' }]
    Measures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Measure' }]
    Metrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Metric' }]
    QualityRules:
 	type : [{ type: Schema.Types.ObjectId, ref: 'QualityRule' }]
    LineageNode:
	type : Schema.Types.ObjectId
    Tags:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Tag' }]
    DataFormat:
 	type : String
#
    collection: 'dataSets'
}
