
// Define collection and schema for InventoryTransaction
export interface InventoryTransaction {
    transactionNumber:
	type : string
    quantity:
	type : Quantity
    transactionDateTime:
	type : Date
    referenceDocument:
	type : string
    Item:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    WorkOrder:
	type : Schema.Types.ObjectId
    PurchaseOrder:
	type : Schema.Types.ObjectId
    SalesOrder:
	type : Schema.Types.ObjectId
    TransactionType:
 	type : String
#
    collection: 'inventoryTransactions'
}
