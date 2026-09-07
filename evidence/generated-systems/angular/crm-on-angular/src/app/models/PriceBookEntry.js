
// Define collection and schema for PriceBookEntry
export interface PriceBookEntry {
    unitPrice:
	type : Money
    effectiveDate:
	type : Date
    expirationDate:
	type : Date
    asActive:
	type : boolean
    PriceBook:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
#
    collection: 'priceBookEntrys'
}
