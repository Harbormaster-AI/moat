
// Define collection and schema for RecordsRepository
export interface RecordsRepository {
    name:
	type : string
    location:
	type : string
    ownerDepartment:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    Systems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'System_' }]
    RetentionSchedules:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RetentionSchedule' }]
    LegalHolds:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LegalHold' }]
    RepositoryType:
 	type : String
#
    collection: 'recordsRepositorys'
}
