
// Define collection and schema for InventoryTransaction
export interface InventoryTransaction {
    transactionNumber:
	type : string
    quantity:
	type : String
    unitCost:
	type : Money
    transactionDate:
	type : Date
    reasonCode:
	type : string
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    RelatedReservation:
	type : Schema.Types.ObjectId
    TransferOrder:
	type : Schema.Types.ObjectId
    Adjustment:
	type : Schema.Types.ObjectId
    CycleCount:
	type : Schema.Types.ObjectId
    TransactionType:
 	type : String
    UnitOfMeasure:
 	type : String
    Status:
 	type : String
#
    collection: 'inventoryTransactions'
}
