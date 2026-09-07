
// Define collection and schema for QualityCheck
export interface QualityCheck {
    checkedAt:
	type : Date
    observedValue:
	type : String
    sampleSize:
	type : number
    Rule:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'qualityChecks'
}
