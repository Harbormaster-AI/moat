
// Define collection and schema for Admission
export interface Admission {
    admitDateTime:
	type : Date
    bed:
	type : string
    Encounter:
	type : Schema.Types.ObjectId
    Facility:
	type : Schema.Types.ObjectId
    AdmissionType:
 	type : String
#
    collection: 'admissions'
}
