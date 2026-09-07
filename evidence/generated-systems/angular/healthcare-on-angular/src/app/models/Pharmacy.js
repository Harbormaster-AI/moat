
// Define collection and schema for Pharmacy
export interface Pharmacy {
    name:
	type : string
    Facility:
	type : Schema.Types.ObjectId
    MedicationDispenses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicationDispense' }]
    MedicationOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MedicationOrder' }]
#
    collection: 'pharmacys'
}
