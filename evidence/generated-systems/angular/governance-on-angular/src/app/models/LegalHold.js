
// Define collection and schema for LegalHold
export interface LegalHold {
    name:
	type : string
    reason:
	type : string
    issuedDate:
	type : Date
    releaseDate:
	type : Date
    Repositories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RecordsRepository' }]
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    Matter:
	type : Schema.Types.ObjectId
    HoldStatus:
 	type : String
#
    collection: 'legalHolds'
}
