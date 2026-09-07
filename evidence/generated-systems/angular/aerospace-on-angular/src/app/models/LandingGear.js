
// Define collection and schema for LandingGear
export interface LandingGear {
    supplierPartNumber:
	type : string
    Supplier:
	type : Schema.Types.ObjectId
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    GearType:
 	type : String
#
    collection: 'landingGears'
}
