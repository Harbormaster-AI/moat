
// Define collection and schema for ForecastLine
export interface ForecastLine {
    period:
	type : Date
    quantity:
	type : Quantity
    confidence:
	type : Percentage
    Forecast:
	type : Schema.Types.ObjectId
    Item:
	type : Schema.Types.ObjectId
#
    collection: 'forecastLines'
}
