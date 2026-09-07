
// Define collection and schema for Registration
export interface Registration {
    tailNumber:
	type : TailNumber
    registryCountry:
	type : string
    Aircraft:
	type : Schema.Types.ObjectId
#
    collection: 'registrations'
}
