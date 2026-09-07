
// Define collection and schema for Branch
export interface Branch {
    name:
	type : string
    branchCode:
	type : string
    address:
	type : Address
    Institution:
	type : Schema.Types.ObjectId
#
    collection: 'branchs'
}
