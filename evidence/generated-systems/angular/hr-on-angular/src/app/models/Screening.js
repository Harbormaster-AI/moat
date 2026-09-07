
// Define collection and schema for Screening
export interface Screening {
    name:
	type : string
    completedDate:
	type : Date
    Application:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'screenings'
}
