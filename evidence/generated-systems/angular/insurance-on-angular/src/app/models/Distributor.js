
// Define collection and schema for Distributor
export interface Distributor {
    name:
	type : string
    licenseNumber:
	type : string
    region:
	type : string
    Insurers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Insurer' }]
    Agents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Agent' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    DistributorType:
 	type : String
#
    collection: 'distributors'
}
