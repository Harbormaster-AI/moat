
// Define collection and schema for InventoryThresholdAlert
export interface InventoryThresholdAlert {
    alertNumber:
	type : string
    detectedAt:
	type : Date
    message:
	type : string
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    RelatedPolicy:
	type : Schema.Types.ObjectId
    AlertType:
 	type : String
    Status:
 	type : String
#
    collection: 'inventoryThresholdAlerts'
}
