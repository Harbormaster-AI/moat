
// Define collection and schema for JobFamily
export interface JobFamily {
    name:
	type : string
    description:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    JobProfiles:
 	type : [{ type: Schema.Types.ObjectId, ref: 'JobProfile' }]
#
    collection: 'jobFamilys'
}
