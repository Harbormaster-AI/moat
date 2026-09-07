
// Define collection and schema for AircraftFamily
export interface AircraftFamily {
    name:
	type : string
    familyCode:
	type : string
    Program:
	type : Schema.Types.ObjectId
    AircraftModels:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftModel' }]
#
    collection: 'aircraftFamilys'
}
