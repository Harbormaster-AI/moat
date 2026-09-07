
// Define collection and schema for Clinician
export interface Clinician {
    firstName:
	type : string
    lastName:
	type : string
    licenseNumber:
	type : string
    CareTeams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CareTeam' }]
    Appointments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Appointment' }]
    Encounters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Encounter' }]
    Procedures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Procedure' }]
    ImagingReports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingReport' }]
    ClinicianType:
 	type : String
    Specialty:
 	type : String
#
    collection: 'clinicians'
}
