
// Define collection and schema for Account
export interface Account {
    accountNumber:
	type : AccountNumber
    iban:
	type : IBAN
    bic:
	type : BIC
    openedDate:
	type : Date
    currency:
	type : string
    balance:
	type : Money
    availableBalance:
	type : Money
    Customer:
	type : Schema.Types.ObjectId
    Institution:
	type : Schema.Types.ObjectId
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    Cards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentCard' }]
    Statements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AccountStatement' }]
    Mandates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DirectDebitMandate' }]
    AccountType:
 	type : String
    Status:
 	type : String
#
    collection: 'accounts'
}
