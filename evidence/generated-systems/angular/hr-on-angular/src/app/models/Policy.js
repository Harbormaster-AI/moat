
// Define collection and schema for Policy
export interface Policy {
    policyNumber:
	type : string
    name:
	type : string
    effectiveDate:
	type : Date
    description:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Acknowledgements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PolicyAcknowledgement' }]
#
    collection: 'policys'
}
