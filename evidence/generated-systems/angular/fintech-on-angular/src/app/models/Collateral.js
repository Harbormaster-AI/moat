
// Define collection and schema for Collateral
export interface Collateral {
    description:
	type : string
    value:
	type : Money
    Loan:
	type : Schema.Types.ObjectId
    CollateralType:
 	type : String
#
    collection: 'collaterals'
}
