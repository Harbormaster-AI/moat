
// Define collection and schema for FraudSignal
export interface FraudSignal {
    name:
	type : string
    ruleLogic:
	type : string
    Scenario:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
    ModelVersion:
	type : Schema.Types.ObjectId
    SignalType:
 	type : String
#
    collection: 'fraudSignals'
}
