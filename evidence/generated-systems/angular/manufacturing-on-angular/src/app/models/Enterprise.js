
// Define collection and schema for Enterprise
export interface Enterprise {
    name:
	type : string
    legalName:
	type : string
    registrationCountry:
	type : string
    website:
	type : string
    taxId:
	type : string
    BusinessUnits:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessUnit' }]
    Plants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Plant' }]
    Suppliers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Supplier' }]
    Customers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Customer' }]
#
    collection: 'enterprises'
}
