
// Define collection and schema for ServiceBulletin
export interface ServiceBulletin {
    bulletinNumber:
	type : string
    WorkOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceWorkOrder' }]
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    Category:
 	type : String
#
    collection: 'serviceBulletins'
}
