
// Define collection and schema for ControlTest_
export interface ControlTest_ {
    name:
	type : string
    testPeriodStart:
	type : Date
    testPeriodEnd:
	type : Date
    sampleSize:
	type : number
    Control:
	type : Schema.Types.ObjectId
    Evidence:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Evidence' }]
    Engagement:
	type : Schema.Types.ObjectId
    TestType:
 	type : String
    Effectiveness:
 	type : String
    Status:
 	type : String
#
    collection: 'controlTest_s'
}
