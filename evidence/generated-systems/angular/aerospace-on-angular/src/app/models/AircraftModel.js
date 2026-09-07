
// Define collection and schema for AircraftModel
export interface AircraftModel {
    name:
	type : string
    modelDesignation:
	type : string
    Family:
	type : Schema.Types.ObjectId
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    EngineTypes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EngineType' }]
    AircraftType:
 	type : String
#
    collection: 'aircraftModels'
}
