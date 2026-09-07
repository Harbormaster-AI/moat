
// Define collection and schema for SoftwareLoad
export interface SoftwareLoad {
    version:
	type : string
    ConnectedAircraft:
	type : Schema.Types.ObjectId
    AvionicsSuite:
	type : Schema.Types.ObjectId
    LoadType:
 	type : String
#
    collection: 'softwareLoads'
}
