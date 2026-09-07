
// Define collection and schema for DataCategory
export interface DataCategory {
    name:
	type : string
    description:
	type : string
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    DataBreaches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataBreach' }]
    Classification:
 	type : String
#
    collection: 'dataCategorys'
}
