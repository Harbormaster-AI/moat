
// Define collection and schema for ImagingReport
export interface ImagingReport {
    reportNumber:
	type : string
    impression:
	type : string
    reportedDate:
	type : Date
    ImagingOrder:
	type : Schema.Types.ObjectId
    Clinician:
	type : Schema.Types.ObjectId
    Encounter:
	type : Schema.Types.ObjectId
    ImagingCenter:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'imagingReports'
}
