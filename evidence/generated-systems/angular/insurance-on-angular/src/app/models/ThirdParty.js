
// Define collection and schema for ThirdParty
export interface ThirdParty {
    name:
	type : string
    taxId:
	type : string
    address:
	type : Address
    Subrogations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SubrogationRecovery' }]
    PartyType:
 	type : String
#
    collection: 'thirdPartys'
}
