
// Define collection and schema for ProductionCertificate
export interface ProductionCertificate {
    certificateNumber:
	type : string
    authority:
	type : string
    Manufacturer:
	type : Schema.Types.ObjectId
#
    collection: 'productionCertificates'
}
