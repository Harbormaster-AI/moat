
// Define collection and schema for SerialNumber
export interface SerialNumber {
    serial:
	type : SerialCode
    activationDate:
	type : Date
    Sku:
	type : Schema.Types.ObjectId
    CurrentInventoryItem:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'serialNumbers'
}
