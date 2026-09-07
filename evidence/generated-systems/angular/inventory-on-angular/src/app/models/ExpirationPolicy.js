
// Define collection and schema for ExpirationPolicy
export interface ExpirationPolicy {
    rejectIfDaysToExpireLessThan:
	type : number
    autoQuarantineDaysToExpire:
	type : number
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    RotationMethod:
 	type : String
#
    collection: 'expirationPolicys'
}
