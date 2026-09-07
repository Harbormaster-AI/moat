
// Define collection and schema for MedicalDevice
export interface MedicalDevice {
    udi:
	type : string
    manufacturer:
	type : string
    Patient:
	type : Schema.Types.ObjectId
    Observations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Observation' }]
    SoftwareUpdates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SoftwareUpdate' }]
    DeviceType:
 	type : String
    ConnectivityStatus:
 	type : String
#
    collection: 'medicalDevices'
}
