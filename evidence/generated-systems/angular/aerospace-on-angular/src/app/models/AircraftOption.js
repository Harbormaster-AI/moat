
// Define collection and schema for AircraftOption
export interface AircraftOption {
    code:
	type : string
    name:
	type : string
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    Packages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftPackage' }]
    OptionCategory:
 	type : String
#
    collection: 'aircraftOptions'
}
