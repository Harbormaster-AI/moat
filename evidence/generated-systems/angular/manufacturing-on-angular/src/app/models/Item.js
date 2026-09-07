
// Define collection and schema for Item
export interface Item {
    itemNumber:
	type : string
    name:
	type : string
    standardCost:
	type : Money
    weight:
	type : Measurement
    asSerialControlled:
	type : boolean
    BusinessUnit:
	type : Schema.Types.ObjectId
    Boms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BOM' }]
    Routings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Routing' }]
    Suppliers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Supplier' }]
    QualitySpecifications:
 	type : [{ type: Schema.Types.ObjectId, ref: 'QualitySpecification' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    ItemType:
 	type : String
    ProcurementType:
 	type : String
    UnitOfMeasure:
 	type : String
    LifecycleStatus:
 	type : String
#
    collection: 'items'
}
