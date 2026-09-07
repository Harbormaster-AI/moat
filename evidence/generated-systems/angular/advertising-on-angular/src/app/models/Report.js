
// Define collection and schema for Report
export interface Report {
    reportName:
	type : string
    generatedAt:
	type : Date
    fileUrl:
	type : URL
    AdAccount:
	type : Schema.Types.ObjectId
    Campaign:
	type : Schema.Types.ObjectId
    LineItem:
	type : Schema.Types.ObjectId
    ReportType:
 	type : String
#
    collection: 'reports'
}
