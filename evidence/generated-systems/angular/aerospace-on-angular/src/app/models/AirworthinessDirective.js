
// Define collection and schema for AirworthinessDirective
export interface AirworthinessDirective {
    directiveNumber:
	type : string
    title:
	type : string
    WorkOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceWorkOrder' }]
#
    collection: 'airworthinessDirectives'
}
