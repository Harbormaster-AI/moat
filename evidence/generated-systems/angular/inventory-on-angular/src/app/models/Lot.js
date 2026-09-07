
// Define collection and schema for Lot
export interface Lot {
    batchNumber:
	type : BatchNumber
    manufactureDate:
	type : Date
    expirationDate:
	type : Date
    Sku:
	type : Schema.Types.ObjectId
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    LotStatus:
 	type : String
#
    collection: 'lots'
}
