
// Define collection and schema for AdAccount
export interface AdAccount {
    name:
	type : string
    accountCode:
	type : string
    defaultCurrency:
	type : string
    defaultTimezone:
	type : string
    Advertiser:
	type : Schema.Types.ObjectId
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    BillingProfile:
	type : Schema.Types.ObjectId
    Dsp:
	type : Schema.Types.ObjectId
    PerformanceMetrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PerformanceMetric' }]
#
    collection: 'adAccounts'
}
