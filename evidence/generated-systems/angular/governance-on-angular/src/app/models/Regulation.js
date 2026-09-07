
// Define collection and schema for Regulation
export interface Regulation {
    name:
	type : string
    citation:
	type : string
    jurisdiction:
	type : string
    publicationUrl:
	type : URL
    Obligations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Obligation' }]
    CompliancePrograms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceProgram' }]
#
    collection: 'regulations'
}
