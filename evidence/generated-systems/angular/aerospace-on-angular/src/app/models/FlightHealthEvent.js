
// Define collection and schema for FlightHealthEvent
export interface FlightHealthEvent {
    eventCode:
	type : string
    ConnectedAircraft:
	type : Schema.Types.ObjectId
    Severity:
 	type : String
#
    collection: 'flightHealthEvents'
}
