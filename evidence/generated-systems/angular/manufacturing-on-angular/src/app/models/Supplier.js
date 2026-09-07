
// Define collection and schema for Supplier
export interface Supplier {
    name:
	type : string
    supplierCode:
	type : string
    address:
	type : Address
    Enterprises:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Enterprise' }]
    Items:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Item' }]
    PurchaseOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PurchaseOrder' }]
    SupplierTier:
 	type : String
    PaymentTerms:
 	type : String
#
    collection: 'suppliers'
}
