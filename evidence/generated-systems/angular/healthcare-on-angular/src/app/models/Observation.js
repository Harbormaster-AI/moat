
// Define collection and schema for Observation
export interface Observation {
    code:
	type : string
    value:
	type : string
    unit:
	type : string
    effectiveDateTime:
	type : Date
    Encounter:
	type : Schema.Types.ObjectId
    Patient:
	type : Schema.Types.ObjectId
    Device:
	type : Schema.Types.ObjectId
    LabResult:
	type : Schema.Types.ObjectId
    Interpretation:
 	type : String
#
    collection: 'observations'
}
