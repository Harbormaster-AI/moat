
// Define collection and schema for CabinLayout
export interface CabinLayout {
    layoutCode:
	type : string
    totalSeats:
	type : number
    classLayout:
	type : string
    Variant:
	type : Schema.Types.ObjectId
    Aircraft:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Aircraft' }]
    Options:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftOption' }]
#
    collection: 'cabinLayouts'
}
