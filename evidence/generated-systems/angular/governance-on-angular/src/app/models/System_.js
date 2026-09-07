
// Define collection and schema for System_
export interface System_ {
    name:
	type : string
    ownerDepartment:
	type : string
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    RecordsRepositories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RecordsRepository' }]
    SystemType:
 	type : String
#
    collection: 'system_s'
}
