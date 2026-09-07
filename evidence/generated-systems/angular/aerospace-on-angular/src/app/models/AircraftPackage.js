
// Define collection and schema for AircraftPackage
export interface AircraftPackage {
    name:
	type : string
    Options:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftOption' }]
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    PackageType:
 	type : String
#
    collection: 'aircraftPackages'
}
