
// Define collection and schema for AccountStatement
export interface AccountStatement {
    statementNumber:
	type : string
    periodStart:
	type : Date
    periodEnd:
	type : Date
    openingBalance:
	type : Money
    closingBalance:
	type : Money
    generatedAt:
	type : DateTime
    Account:
	type : Schema.Types.ObjectId
#
    collection: 'accountStatements'
}
