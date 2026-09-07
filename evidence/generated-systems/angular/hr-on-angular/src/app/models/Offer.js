
// Define collection and schema for Offer
export interface Offer {
    offerNumber:
	type : string
    proposedStartDate:
	type : Date
    baseSalary:
	type : Money
    signOnBonus:
	type : Money
    Requisition:
	type : Schema.Types.ObjectId
    Candidate:
	type : Schema.Types.ObjectId
    ApprovedBy:
	type : Schema.Types.ObjectId
    Contract:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'offers'
}
