
// Define collection and schema for Application
export interface Application {
    applicationNumber:
	type : string
    submissionDate:
	type : Date
    Customer:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
    Distributor:
	type : Schema.Types.ObjectId
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    SelectedQuote:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'applications'
}
