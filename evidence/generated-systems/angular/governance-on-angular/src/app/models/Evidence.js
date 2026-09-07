
// Define collection and schema for Evidence
export interface Evidence {
    title:
	type : string
    locationUrl:
	type : URL
    receivedDate:
	type : Date
    ControlTest:
	type : Schema.Types.ObjectId
    Control:
	type : Schema.Types.ObjectId
    Obligation:
	type : Schema.Types.ObjectId
    Workpaper:
	type : Schema.Types.ObjectId
    EvidenceType:
 	type : String
#
    collection: 'evidences'
}
