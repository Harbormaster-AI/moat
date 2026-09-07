
// Define collection and schema for Anomaly
export interface Anomaly {
    occurredAt:
	type : Date
    details:
	type : string
    TimeSeries:
	type : Schema.Types.ObjectId
    Alert:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
    AnomalyType:
 	type : String
    Severity:
 	type : String
#
    collection: 'anomalys'
}
