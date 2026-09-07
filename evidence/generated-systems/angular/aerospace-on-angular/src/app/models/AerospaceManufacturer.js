
// Define collection and schema for AerospaceManufacturer
export interface AerospaceManufacturer {
    name:
	type : string
    legalName:
	type : string
    headquartersCountry:
	type : string
    website:
	type : string
    Programs:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftProgram' }]
    Plants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Plant' }]
    Suppliers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Supplier' }]
    ProductionCertificates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductionCertificate' }]
#
    collection: 'aerospaceManufacturers'
}
