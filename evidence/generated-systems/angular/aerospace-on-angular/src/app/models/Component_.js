
// Define collection and schema for Component_
export interface Component_ {
    partNumber:
	type : string
    name:
	type : string
    Supplier:
	type : Schema.Types.ObjectId
    ComponentCategory:
 	type : String
    SerializationMethod:
 	type : String
#
    collection: 'component_s'
}
