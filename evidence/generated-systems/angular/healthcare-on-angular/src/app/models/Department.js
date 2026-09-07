
// Define collection and schema for Department
export interface Department {
    name:
	type : string
    Facility:
	type : Schema.Types.ObjectId
    CareTeams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CareTeam' }]
    DepartmentType:
 	type : String
#
    collection: 'departments'
}
