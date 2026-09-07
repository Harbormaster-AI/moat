
// Define collection and schema for InspectionCharacteristic
export interface InspectionCharacteristic {
    characteristicCode:
	type : string
    name:
	type : string
    lowerSpecLimit:
	type : Measurement
    upperSpecLimit:
	type : Measurement
    target:
	type : Measurement
    InspectionPlan:
	type : Schema.Types.ObjectId
    MeasurementType:
 	type : String
#
    collection: 'inspectionCharacteristics'
}
