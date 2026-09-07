
// Define collection and schema for BankAccount
export interface BankAccount {
    accountHolder:
	type : string
    bankName:
	type : string
    iban:
	type : string
    bic:
	type : string
    accountNumber:
	type : string
    routingNumber:
	type : string
#
    collection: 'bankAccounts'
}
