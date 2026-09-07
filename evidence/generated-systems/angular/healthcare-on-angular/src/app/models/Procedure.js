
// Define collection and schema for Procedure
export interface Procedure {
    procedureCode:
	type : string
    startDateTime:
	type : Date
    endDateTime:
	type : Date
    Encounter:
	type : Schema.Types.ObjectId
    Performer:
	type : Schema.Types.ObjectId
    ProcedureOrder:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'procedures'
}
