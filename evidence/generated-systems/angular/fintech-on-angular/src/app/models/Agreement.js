
// Define collection and schema for Agreement
export interface Agreement {
    agreementNumber:
	type : string
    effectiveDate:
	type : Date
    Customer:
	type : Schema.Types.ObjectId
    ProductOffering:
	type : Schema.Types.ObjectId
    AgreementType:
 	type : String
    Status:
 	type : String
#
    collection: 'agreements'
}
