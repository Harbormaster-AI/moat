
// Define collection and schema for MaintenanceAppointment
export interface MaintenanceAppointment {
    appointmentDate:
	type : Date
    Aircraft:
	type : Schema.Types.ObjectId
    MroFacility:
	type : Schema.Types.ObjectId
    WorkOrder:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'maintenanceAppointments'
}
