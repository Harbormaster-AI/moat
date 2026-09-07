
// Define collection and schema for Nonconformance
export interface Nonconformance {
    ncNumber:
	type : string
    description:
	type : string
    containmentAction:
	type : string
    Item:
	type : Schema.Types.ObjectId
    WorkOrder:
	type : Schema.Types.ObjectId
    InspectionLot:
	type : Schema.Types.ObjectId
    CorrectiveAction:
	type : Schema.Types.ObjectId
    NcType:
 	type : String
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'nonconformances'
}
