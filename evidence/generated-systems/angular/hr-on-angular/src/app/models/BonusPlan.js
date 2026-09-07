
// Define collection and schema for BonusPlan
export interface BonusPlan {
    name:
	type : string
    targetPercentage:
	type : Percentage
    CompensationPackages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CompensationPackage' }]
#
    collection: 'bonusPlans'
}
