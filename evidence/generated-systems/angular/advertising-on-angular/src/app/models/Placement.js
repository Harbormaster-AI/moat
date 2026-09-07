
// Define collection and schema for Placement
export interface Placement {
    name:
	type : string
    flight:
	type : DateRange
    goalImpressions:
	type : number
    LineItem:
	type : Schema.Types.ObjectId
    AdSlot:
	type : Schema.Types.ObjectId
    Deal:
	type : Schema.Types.ObjectId
#
    collection: 'placements'
}
