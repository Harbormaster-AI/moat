
// Define collection and schema for Encounter
export interface Encounter {
    encounterNumber:
	type : string
    startDateTime:
	type : Date
    endDateTime:
	type : Date
    Patient:
	type : Schema.Types.ObjectId
    Clinician:
	type : Schema.Types.ObjectId
    Facility:
	type : Schema.Types.ObjectId
    Appointment:
	type : Schema.Types.ObjectId
    Diagnoses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Diagnosis' }]
    Procedures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Procedure' }]
    Observations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Observation' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ClinicalOrder' }]
    Admission:
	type : Schema.Types.ObjectId
    Discharge:
	type : Schema.Types.ObjectId
    Status:
 	type : String
    EncounterType:
 	type : String
#
    collection: 'encounters'
}
