
// Define collection and schema for ClaimReserve
export interface ClaimReserve {
    amount:
	type : Money
    setDate:
	type : Date
    Claim:
	type : Schema.Types.ObjectId
    Exposure:
	type : Schema.Types.ObjectId
    ReserveType:
 	type : String
    Status:
 	type : String
#
    collection: 'claimReserves'
}
