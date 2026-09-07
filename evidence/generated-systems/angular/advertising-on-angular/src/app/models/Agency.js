
// Define collection and schema for Agency
export interface Agency {
    name:
	type : string
    legalName:
	type : string
    headquartersCountry:
	type : string
    website:
	type : string
    Advertisers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Advertiser' }]
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    InsertionOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsertionOrder' }]
#
    collection: 'agencys'
}
