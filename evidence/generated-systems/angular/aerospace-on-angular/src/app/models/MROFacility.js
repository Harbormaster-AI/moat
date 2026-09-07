
// Define collection and schema for MROFacility
export interface MROFacility {
    name:
	type : string
    approvalScope:
	type : string
    address:
	type : Address
    Appointments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceAppointment' }]
    WorkOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceWorkOrder' }]
#
    collection: 'mROFacilitys'
}
