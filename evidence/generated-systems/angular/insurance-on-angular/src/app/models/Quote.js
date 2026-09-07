
// Define collection and schema for Quote
export interface Quote {
    quoteNumber:
	type : string
    totalPremium:
	type : Money
    ratingDate:
	type : Date
    asBound:
	type : boolean
    Application:
	type : Schema.Types.ObjectId
    UnderwritingDecisions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'UnderwritingDecision' }]
    Policy:
	type : Schema.Types.ObjectId
#
    collection: 'quotes'
}
