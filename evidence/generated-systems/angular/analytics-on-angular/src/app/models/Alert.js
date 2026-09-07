
// Define collection and schema for Alert
export interface Alert {
    title:
	type : string
    createdAt:
	type : Date
    Metric:
	type : Schema.Types.ObjectId
    Dashboard:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
    Rule:
	type : Schema.Types.ObjectId
    Anomalies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Anomaly' }]
    Subscribers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Subscriber' }]
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'alerts'
}
