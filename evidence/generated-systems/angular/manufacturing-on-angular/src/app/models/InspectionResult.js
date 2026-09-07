
// Define collection and schema for InspectionResult
export interface InspectionResult {
    resultValue:
	type : Measurement
    recordedOn:
	type : Date
    notes:
	type : string
    InspectionLot:
	type : Schema.Types.ObjectId
    Characteristic:
	type : Schema.Types.ObjectId
    ResultStatus:
 	type : String
#
    collection: 'inspectionResults'
}
