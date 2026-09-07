
// Define collection and schema for Plant
export interface Plant {
    name:
	type : string
    plantCode:
	type : string
    address:
	type : Address
    timeZone:
	type : string
    Enterprise:
	type : Schema.Types.ObjectId
    ProductionLines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductionLine' }]
    WorkCenters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkCenter' }]
    Warehouses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Warehouse' }]
    Assets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Asset' }]
    ProductionSchedules:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductionSchedule' }]
#
    collection: 'plants'
}
