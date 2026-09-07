
// Define collection and schema for Patient
export interface Patient {
    firstName:
	type : string
    lastName:
	type : string
    mrn:
	type : MRN
    dateOfBirth:
	type : Date
    address:
	type : Address
    primaryLanguage:
	type : string
    Appointments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Appointment' }]
    Encounters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Encounter' }]
    CarePlans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CarePlan' }]
    Allergies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Allergy' }]
    Conditions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Condition' }]
    MedicationOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicationOrder' }]
    LabOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LaboratoryOrder' }]
    ImagingOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingOrder' }]
    Coverages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Coverage' }]
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    Devices:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicalDevice' }]
    Observations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Observation' }]
    SexAtBirth:
 	type : String
    BloodType:
 	type : String
#
    collection: 'patients'
}
