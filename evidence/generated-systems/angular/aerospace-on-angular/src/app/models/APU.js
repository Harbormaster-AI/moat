
// Define collection and schema for APU
export interface APU {
    model_:
	type : string
    Supplier:
	type : Schema.Types.ObjectId
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
#
    collection: 'aPUs'
}
