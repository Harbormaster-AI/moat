
// Define collection and schema for TimeSeries
export interface TimeSeries {
    name:
	type : string
    timezone:
	type : string
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Forecasts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Forecast' }]
    Anomalies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Anomaly' }]
    Granularity:
 	type : String
#
    collection: 'timeSeriess'
}
