
// Define collection and schema for SalaryComponent
export interface SalaryComponent {
    amount:
	type : Money
    recurring:
	type : boolean
    CompensationPackage:
	type : Schema.Types.ObjectId
    ComponentType:
 	type : String
#
    collection: 'salaryComponents'
}
