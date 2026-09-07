
// Define collection and schema for PayrollItem
export interface PayrollItem {
    amount:
	type : Money
    taxable:
	type : boolean
    PayrollRun:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    ItemType:
 	type : String
#
    collection: 'payrollItems'
}
