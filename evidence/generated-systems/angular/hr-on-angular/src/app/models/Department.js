
// Define collection and schema for Department
export interface Department {
    name:
	type : string
    code:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Manager:
	type : Schema.Types.ObjectId
    Positions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    Employees:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Employee' }]
    CostCenter:
	type : Schema.Types.ObjectId
#
    collection: 'departments'
}
