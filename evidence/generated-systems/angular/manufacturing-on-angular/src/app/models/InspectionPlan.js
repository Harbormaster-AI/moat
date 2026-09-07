
// Define collection and schema for InspectionPlan
export interface InspectionPlan {
    planNumber:
	type : string
    revision:
	type : string
    Item:
	type : Schema.Types.ObjectId
    Characteristics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InspectionCharacteristic' }]
    SamplingPlan:
 	type : String
    Status:
 	type : String
#
    collection: 'inspectionPlans'
}
