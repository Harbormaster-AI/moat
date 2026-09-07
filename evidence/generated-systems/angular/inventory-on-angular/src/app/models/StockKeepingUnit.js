
// Define collection and schema for StockKeepingUnit
export interface StockKeepingUnit {
    skuCode:
	type : SKU
    name:
	type : string
    weight:
	type : String
    weightUnit:
	type : string
    volume:
	type : String
    volumeUnit:
	type : string
    shelfLifeDays:
	type : number
    hazardousMaterial:
	type : boolean
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    UomConversions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'UoMConversion' }]
    ReplenishmentPolicies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ReplenishmentPolicy' }]
    Lots:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Lot' }]
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    ItemType:
 	type : String
    UnitOfMeasure:
 	type : String
#
    collection: 'stockKeepingUnits'
}
