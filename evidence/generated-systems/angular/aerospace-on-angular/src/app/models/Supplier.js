
// Define collection and schema for Supplier
export interface Supplier {
    name:
	type : string
    Manufacturers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AerospaceManufacturer' }]
    Components:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Component_' }]
    EngineTypes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EngineType' }]
    AvionicsSuites:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AvionicsSuite' }]
    Apus:
 	type : [{ type: Schema.Types.ObjectId, ref: 'APU' }]
    LandingGears:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LandingGear' }]
    SupplierType:
 	type : String
    ApprovalStatus:
 	type : String
#
    collection: 'suppliers'
}
