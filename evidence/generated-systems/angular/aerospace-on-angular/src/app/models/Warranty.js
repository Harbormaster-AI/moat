
// Define collection and schema for Warranty
export interface Warranty {
    coverageMonths:
	type : number
    Aircraft:
	type : Schema.Types.ObjectId
    WarrantyType:
 	type : String
#
    collection: 'warrantys'
}
