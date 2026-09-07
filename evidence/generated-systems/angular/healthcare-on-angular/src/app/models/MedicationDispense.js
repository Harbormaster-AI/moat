
// Define collection and schema for MedicationDispense
export interface MedicationDispense {
    dispenseNumber:
	type : string
    quantity:
	type : String
    whenPrepared:
	type : Date
    MedicationOrder:
	type : Schema.Types.ObjectId
    Pharmacy:
	type : Schema.Types.ObjectId
    Patient:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'medicationDispenses'
}
