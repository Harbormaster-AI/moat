
// Define collection and schema for Diagnosis
export interface Diagnosis {
    code:
	type : string
    description:
	type : string
    onsetDate:
	type : Date
    Encounter:
	type : Schema.Types.ObjectId
    Patient:
	type : Schema.Types.ObjectId
    Certainty:
 	type : String
#
    collection: 'diagnosiss'
}
