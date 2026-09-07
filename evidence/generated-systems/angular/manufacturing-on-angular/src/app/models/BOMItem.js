
// Define collection and schema for BOMItem
export interface BOMItem {
    lineNumber:
	type : number
    quantity:
	type : Quantity
    scrapPercent:
	type : Percentage
    Bom:
	type : Schema.Types.ObjectId
    Component:
	type : Schema.Types.ObjectId
#
    collection: 'bOMItems'
}
