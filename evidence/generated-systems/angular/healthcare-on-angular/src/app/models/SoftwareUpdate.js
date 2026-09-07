
// Define collection and schema for SoftwareUpdate
export interface SoftwareUpdate {
    version:
	type : string
    appliedDate:
	type : Date
    Device:
	type : Schema.Types.ObjectId
    UpdateType:
 	type : String
#
    collection: 'softwareUpdates'
}
