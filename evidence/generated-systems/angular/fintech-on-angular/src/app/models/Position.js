
// Define collection and schema for Position
export interface Position {
    quantity:
	type : String
    averageCost:
	type : Money
    marketValue:
	type : Money
    Portfolio:
	type : Schema.Types.ObjectId
    Security:
	type : Schema.Types.ObjectId
#
    collection: 'positions'
}
