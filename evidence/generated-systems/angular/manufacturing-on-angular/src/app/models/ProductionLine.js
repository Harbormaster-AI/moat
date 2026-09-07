
// Define collection and schema for ProductionLine
export interface ProductionLine {
    name:
	type : string
    lineCode:
	type : string
    Plant:
	type : Schema.Types.ObjectId
    WorkCenters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkCenter' }]
    LineType:
 	type : String
#
    collection: 'productionLines'
}
