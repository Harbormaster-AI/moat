
// Define collection and schema for Insurer
export interface Insurer {
    name:
	type : string
    legalName:
	type : string
    domicileCountry:
	type : string
    naicNumber:
	type : string
    website:
	type : string
    Products:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsuranceProduct' }]
    DistributionPartners:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Distributor' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    ReinsuranceAgreements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ReinsuranceAgreement' }]
#
    collection: 'insurers'
}
