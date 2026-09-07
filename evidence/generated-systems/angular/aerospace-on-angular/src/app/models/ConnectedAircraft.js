
// Define collection and schema for ConnectedAircraft
export interface ConnectedAircraft {
    communicationsProvider:
	type : string
    Aircraft:
	type : Schema.Types.ObjectId
    FlightHealthEvents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FlightHealthEvent' }]
    SoftwareLoads:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SoftwareLoad' }]
    ConnectivityStatus:
 	type : String
#
    collection: 'connectedAircrafts'
}
