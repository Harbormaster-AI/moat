
// Define collection and schema for HealthSystem
export interface HealthSystem {
    name:
	type : string
    legalName:
	type : string
    headquartersCountry:
	type : string
    website:
	type : string
    Facilities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Facility' }]
    Suppliers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicalSupplier' }]
#
    collection: 'healthSystems'
}
