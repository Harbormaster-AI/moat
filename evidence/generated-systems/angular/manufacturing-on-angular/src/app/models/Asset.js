
// Define collection and schema for Asset
export interface Asset {
    assetTag:
	type : string
    assetName:
	type : string
    commissioningDate:
	type : Date
    Plant:
	type : Schema.Types.ObjectId
    WorkCenter:
	type : Schema.Types.ObjectId
    MaintenanceOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceOrder' }]
    MaintenancePlans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenancePlan' }]
    AssetStatus:
 	type : String
#
    collection: 'assets'
}
