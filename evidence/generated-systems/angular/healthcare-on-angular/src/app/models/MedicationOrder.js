
// Define collection and schema for MedicationOrder
export interface MedicationOrder {
    medicationCode:
	type : string
    dose:
	type : Dose
    frequency:
	type : string
    duration:
	type : string
    Order:
	type : Schema.Types.ObjectId
    Pharmacy:
	type : Schema.Types.ObjectId
    Dispenses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicationDispense' }]
    Route:
 	type : String
#
    collection: 'medicationOrders'
}
