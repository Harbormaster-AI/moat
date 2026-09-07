
// Define collection and schema for CareTeam
export interface CareTeam {
    name:
	type : string
    Department:
	type : Schema.Types.ObjectId
    Clinicians:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Clinician' }]
    Patients:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Patient' }]
    CareSetting:
 	type : String
#
    collection: 'careTeams'
}
