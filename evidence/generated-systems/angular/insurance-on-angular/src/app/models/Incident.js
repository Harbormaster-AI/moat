
// Define collection and schema for Incident
export interface Incident {
    location:
	type : Address
    description:
	type : string
    Claim:
	type : Schema.Types.ObjectId
    InsuredObjects:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsuredObject' }]
    IncidentType:
 	type : String
#
    collection: 'incidents'
}
