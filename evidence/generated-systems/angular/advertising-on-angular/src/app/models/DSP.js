
// Define collection and schema for DSP
export interface DSP {
    name:
	type : string
    website:
	type : string
    region:
	type : string
    AdAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AdAccount' }]
#
    collection: 'dSPs'
}
