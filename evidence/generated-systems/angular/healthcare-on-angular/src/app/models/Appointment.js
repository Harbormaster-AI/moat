
// Define collection and schema for Appointment
export interface Appointment {
    appointmentDate:
	type : Date
    reason:
	type : string
    Patient:
	type : Schema.Types.ObjectId
    Clinician:
	type : Schema.Types.ObjectId
    Facility:
	type : Schema.Types.ObjectId
    Encounter:
	type : Schema.Types.ObjectId
    Status:
 	type : String
    Priority:
 	type : String
#
    collection: 'appointments'
}
