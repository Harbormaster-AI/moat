
// Define collection and schema for TaxWithholding
export interface TaxWithholding {
    taxId:
	type : TaxId
    allowances:
	type : number
    additionalAmount:
	type : Money
    Employee:
	type : Schema.Types.ObjectId
    FilingStatus:
 	type : String
#
    collection: 'taxWithholdings'
}
