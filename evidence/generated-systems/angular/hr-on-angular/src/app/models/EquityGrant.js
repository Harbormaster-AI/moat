
// Define collection and schema for EquityGrant
export interface EquityGrant {
    grantId:
	type : string
    grantedUnits:
	type : number
    vestingStart:
	type : Date
    CompensationPackage:
	type : Schema.Types.ObjectId
    GrantType:
 	type : String
#
    collection: 'equityGrants'
}
