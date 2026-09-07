
// Define collection and schema for DirectDebitMandate
export interface DirectDebitMandate {
    mandateId:
	type : string
    signedAt:
	type : DateTime
    Account:
	type : Schema.Types.ObjectId
    Creditor:
	type : Schema.Types.ObjectId
    Scheme:
 	type : String
    Status:
 	type : String
#
    collection: 'directDebitMandates'
}
