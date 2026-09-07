
// Define collection and schema for MedicalSupplier
export interface MedicalSupplier {
    name:
	type : string
    website:
	type : string
    Facilities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Facility' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    SupplierTier:
 	type : String
#
    collection: 'medicalSuppliers'
}
