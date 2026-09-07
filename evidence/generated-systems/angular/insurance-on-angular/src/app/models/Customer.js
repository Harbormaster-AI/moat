
// Define collection and schema for Customer
export interface Customer {
    firstName:
	type : string
    lastName:
	type : string
    organizationName:
	type : string
    taxId:
	type : string
    dateOfBirth:
	type : Date
    primaryAddress:
	type : Address
    Applications:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Application' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    Agents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Agent' }]
    Beneficiaries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Beneficiary' }]
    CustomerType:
 	type : String
#
    collection: 'customers'
}
