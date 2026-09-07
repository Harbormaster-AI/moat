
// Define collection and schema for QualityRule
export interface QualityRule {
    name:
	type : string
    threshold:
	type : Threshold
    targetField:
	type : string
    Dataset:
	type : Schema.Types.ObjectId
    Checks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'QualityCheck' }]
    Dimension:
 	type : String
    Operator:
 	type : String
#
    collection: 'qualityRules'
}
