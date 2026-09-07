
// Define collection and schema for Allergy
export interface Allergy {
    substance:
	type : string
    reaction:
	type : string
    Patient:
	type : Schema.Types.ObjectId
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'allergys'
}
