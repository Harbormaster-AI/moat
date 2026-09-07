
// Define collection and schema for DemandSignal
export interface DemandSignal {
    externalReference:
	type : string
    requestedDate:
	type : Date
    quantity:
	type : String
    Sku:
	type : Schema.Types.ObjectId
    Reservations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Reservation' }]
    DemandType:
 	type : String
#
    collection: 'demandSignals'
}
