
// Define collection and schema for AvionicsSuite
export interface AvionicsSuite {
    suiteName:
	type : string
    softwareBaseline:
	type : string
    Supplier:
	type : Schema.Types.ObjectId
    Variants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftVariant' }]
    SoftwareLoads:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SoftwareLoad' }]
#
    collection: 'avionicsSuites'
}
