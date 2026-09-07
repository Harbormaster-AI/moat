
// Define collection and schema for KPI
export interface KPI {
    targetValue:
	type : String
    Campaign:
	type : Schema.Types.ObjectId
    MetricType:
 	type : String
#
    collection: 'kPIs'
}
