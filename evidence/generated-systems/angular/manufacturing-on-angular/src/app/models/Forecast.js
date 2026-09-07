
// Define collection and schema for Forecast
export interface Forecast {
    forecastNumber:
	type : string
    forecastHorizonStart:
	type : Date
    forecastHorizonEnd:
	type : Date
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ForecastLine' }]
    Method:
 	type : String
#
    collection: 'forecasts'
}
