
// Define collection and schema for Beneficiary
export interface Beneficiary {
    name:
	type : string
    share:
	type : Percentage
    Policy:
	type : Schema.Types.ObjectId
    Customer:
	type : Schema.Types.ObjectId
    Relationship:
 	type : String
#
    collection: 'beneficiarys'
}
