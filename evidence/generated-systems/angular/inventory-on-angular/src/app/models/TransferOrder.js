
// Define collection and schema for TransferOrder
export interface TransferOrder {
    orderNumber:
	type : string
    requestedShipDate:
	type : Date
    requestedReceiveDate:
	type : Date
    shippedDate:
	type : Date
    receivedDate:
	type : Date
    OriginWarehouse:
	type : Schema.Types.ObjectId
    DestinationWarehouse:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TransferOrderLine' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryTransaction' }]
    Status:
 	type : String
#
    collection: 'transferOrders'
}
