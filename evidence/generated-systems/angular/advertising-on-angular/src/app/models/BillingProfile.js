
// Define collection and schema for BillingProfile
export interface BillingProfile {
    billingName:
	type : string
    taxId:
	type : string
    billingAddress:
	type : Address
    Advertiser:
	type : Schema.Types.ObjectId
    PaymentMethods:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentMethod' }]
    AdAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AdAccount' }]
    PaymentTerms:
 	type : String
#
    collection: 'billingProfiles'
}
