
// Define collection and schema for FraudScenario
export interface FraudScenario {
    name:
	type : string
    riskAppetite:
	type : string
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Alert' }]
    Signals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FraudSignal' }]
    DetectionType:
 	type : String
#
    collection: 'fraudScenarios'
}
