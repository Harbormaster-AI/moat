
// Define collection and schema for Adjuster
export interface Adjuster {
    firstName:
	type : string
    lastName:
	type : string
    licenseNumber:
	type : string
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    ServiceProviders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ServiceProvider' }]
    AdjusterType:
 	type : String
#
    collection: 'adjusters'
}
