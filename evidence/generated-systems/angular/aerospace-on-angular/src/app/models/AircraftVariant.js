
// Define collection and schema for AircraftVariant
export interface AircraftVariant {
    variantCode:
	type : string
    rangeNm:
	type : number
    maxTakeoffWeightKg:
	type : String
    Model_:
	type : Schema.Types.ObjectId
    EngineType:
	type : Schema.Types.ObjectId
    AvionicsSuite:
	type : Schema.Types.ObjectId
    Apu:
	type : Schema.Types.ObjectId
    LandingGear:
	type : Schema.Types.ObjectId
    CabinLayouts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CabinLayout' }]
    Options:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftOption' }]
    Packages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftPackage' }]
#
    collection: 'aircraftVariants'
}
