
// Define collection and schema for PolicyAcknowledgement
export interface PolicyAcknowledgement {
    acknowledgementDate:
	type : Date
    Policy:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'policyAcknowledgements'
}
