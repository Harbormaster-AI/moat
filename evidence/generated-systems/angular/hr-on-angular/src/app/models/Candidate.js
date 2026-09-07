
// Define collection and schema for Candidate
export interface Candidate {
    name:
	type : PersonName
    email:
	type : Email
    phone:
	type : PhoneNumber
    Applications:
 	type : [{ type: Schema.Types.ObjectId, ref: 'JobApplication' }]
    Interviews:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Interview' }]
    Offers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Offer' }]
    Documents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Document' }]
    Source:
 	type : String
#
    collection: 'candidates'
}
