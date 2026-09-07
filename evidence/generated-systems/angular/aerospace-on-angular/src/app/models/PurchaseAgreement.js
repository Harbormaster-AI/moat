
// Define collection and schema for PurchaseAgreement
export interface PurchaseAgreement {
    agreementNumber:
	type : string
    effectiveDate:
	type : Date
    AircraftOrder:
	type : Schema.Types.ObjectId
#
    collection: 'purchaseAgreements'
}
