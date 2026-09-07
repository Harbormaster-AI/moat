
// Define collection and schema for Contract
export interface Contract {
    contractNumber:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    renewalTermMonths:
	type : number
    autoRenew:
	type : boolean
    Organization:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Order' }]
    Cases:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Case_' }]
    Status:
 	type : String
#
    collection: 'contracts'
}
