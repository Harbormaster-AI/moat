
// Define collection and schema for PurchaseOrderLine
export interface PurchaseOrderLine {
    lineNumber:
	type : number
    quantity:
	type : Quantity
    unitPrice:
	type : Money
    dueDate:
	type : Date
    PurchaseOrder:
	type : Schema.Types.ObjectId
    Item:
	type : Schema.Types.ObjectId
#
    collection: 'purchaseOrderLines'
}
