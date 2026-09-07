
// Define collection and schema for UoMConversion
export interface UoMConversion {
    factor:
	type : String
    precision:
	type : number
    Sku:
	type : Schema.Types.ObjectId
    FromUnit:
 	type : String
    ToUnit:
 	type : String
#
    collection: 'uoMConversions'
}
