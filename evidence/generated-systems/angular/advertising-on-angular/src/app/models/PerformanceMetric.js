
// Define collection and schema for PerformanceMetric
export interface PerformanceMetric {
    date:
	type : Date
    value:
	type : String
    AdAccount:
	type : Schema.Types.ObjectId
    Campaign:
	type : Schema.Types.ObjectId
    LineItem:
	type : Schema.Types.ObjectId
    Placement:
	type : Schema.Types.ObjectId
    CreativeAsset:
	type : Schema.Types.ObjectId
    MetricType:
 	type : String
#
    collection: 'performanceMetrics'
}
