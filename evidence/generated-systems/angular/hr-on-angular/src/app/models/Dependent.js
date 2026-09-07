
// Define collection and schema for Dependent
export interface Dependent {
    firstName:
	type : string
    lastName:
	type : string
    birthDate:
	type : Date
    BenefitEnrollment:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    Relationship:
 	type : String
#
    collection: 'dependents'
}
