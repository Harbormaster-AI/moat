
// Define collection and schema for Beneficiary
export interface Beneficiary {
    name:
	type : string
    accountIdentifier:
	type : AccountIdentifier
    iban:
	type : IBAN
    bic:
	type : BIC
    address:
	type : Address
    Customer:
	type : Schema.Types.ObjectId
#
    collection: 'beneficiarys'
}
