
// Define collection and schema for Forecast
export interface Forecast {
    name:
	type : string
    horizon:
	type : number
    ModelVersion:
	type : Schema.Types.ObjectId
    TimeSeries:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Granularity:
 	type : String
#
    collection: 'forecasts'
}
