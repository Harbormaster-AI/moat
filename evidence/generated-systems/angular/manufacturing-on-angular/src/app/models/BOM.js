
// Define collection and schema for BOM
export interface BOM {
    bomNumber:
	type : string
    revision:
	type : string
    effectivityStart:
	type : Date
    effectivityEnd:
	type : Date
    ParentItem:
	type : Schema.Types.ObjectId
    BomItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BOMItem' }]
    Status:
 	type : String
#
    collection: 'bOMs'
}
