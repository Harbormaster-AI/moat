
// Define collection and schema for WorkCenter
export interface WorkCenter {
    name:
	type : string
    capability:
	type : string
    ProductionLine:
	type : Schema.Types.ObjectId
#
    collection: 'workCenters'
}
