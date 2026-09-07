
// Define collection and schema for Organization
export interface Organization {
    name:
	type : string
    defaultCurrency:
	type : string
    defaultLocale:
	type : _Locale
    website:
	type : URL
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    Territories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Territory' }]
    Products:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Product' }]
    PriceBooks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PriceBook' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
#
    collection: 'organizations'
}
