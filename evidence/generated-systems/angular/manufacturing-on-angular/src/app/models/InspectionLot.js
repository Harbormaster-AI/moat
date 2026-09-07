
// Define collection and schema for InspectionLot
export interface InspectionLot {
    lotNumber:
	type : string
    quantity:
	type : Quantity
    sampleSize:
	type : number
    createdOn:
	type : Date
    Item:
	type : Schema.Types.ObjectId
    WorkOrder:
	type : Schema.Types.ObjectId
    GoodsReceipt:
	type : Schema.Types.ObjectId
    Results:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InspectionResult' }]
    InspectionType:
 	type : String
    Status:
 	type : String
#
    collection: 'inspectionLots'
}
