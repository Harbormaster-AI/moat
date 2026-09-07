
// Define collection and schema for BusinessUnit
export interface BusinessUnit {
    name:
	type : string
    code:
	type : string
    Enterprise:
	type : Schema.Types.ObjectId
    Items:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Item' }]
    Plants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Plant' }]
    Category:
 	type : String
#
    collection: 'businessUnits'
}
