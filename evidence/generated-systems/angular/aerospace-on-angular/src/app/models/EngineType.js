
// Define collection and schema for EngineType
export interface EngineType {
    engineModelCode:
	type : string
    maxThrustKn:
	type : String
    Supplier:
	type : Schema.Types.ObjectId
    CompatibleModels:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftModel' }]
    Category:
 	type : String
#
    collection: 'engineTypes'
}
