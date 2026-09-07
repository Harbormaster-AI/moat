
// Define collection and schema for SalesOrderLine
export interface SalesOrderLine {
    lineNumber:
	type : number
    quantity:
	type : Quantity
    unitPrice:
	type : Money
    dueDate:
	type : Date
    SalesOrder:
	type : Schema.Types.ObjectId
    Item:
	type : Schema.Types.ObjectId
#
    collection: 'salesOrderLines'
}
