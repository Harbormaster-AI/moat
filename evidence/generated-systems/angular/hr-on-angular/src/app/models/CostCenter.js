
// Define collection and schema for CostCenter
export interface CostCenter {
    code:
	type : string
    name:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Departments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Department' }]
    Positions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    Employees:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Employee' }]
#
    collection: 'costCenters'
}
