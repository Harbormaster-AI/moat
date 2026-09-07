
// Define collection and schema for AircraftProgram
export interface AircraftProgram {
    name:
	type : string
    programCode:
	type : string
    entryIntoServiceYear:
	type : number
    Manufacturer:
	type : Schema.Types.ObjectId
    AircraftFamilies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AircraftFamily' }]
    TypeCertificate:
	type : Schema.Types.ObjectId
    KeySuppliers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Supplier' }]
    Status:
 	type : String
#
    collection: 'aircraftPrograms'
}
