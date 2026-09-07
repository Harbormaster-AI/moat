
// Define collection and schema for Aircraft
export interface Aircraft {
    msn:
	type : MSN
    deliveryDate:
	type : Date
    Variant:
	type : Schema.Types.ObjectId
    Operator:
	type : Schema.Types.ObjectId
    Registration:
	type : Schema.Types.ObjectId
    Warranty:
	type : Schema.Types.ObjectId
    MaintenanceRecords:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceWorkOrder' }]
    ConnectedAircraft:
	type : Schema.Types.ObjectId
    CabinLayout:
	type : Schema.Types.ObjectId
#
    collection: 'aircrafts'
}
