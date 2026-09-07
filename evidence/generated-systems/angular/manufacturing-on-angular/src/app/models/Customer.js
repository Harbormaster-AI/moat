
// Define collection and schema for Customer
export interface Customer {
    name:
	type : string
    customerCode:
	type : string
    address:
	type : Address
    Enterprises:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Enterprise' }]
    SalesOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SalesOrder' }]
    CustomerType:
 	type : String
#
    collection: 'customers'
}
