
// Define collection and schema for Discharge
export interface Discharge {
    dischargeDateTime:
	type : Date
    Encounter:
	type : Schema.Types.ObjectId
    Disposition:
 	type : String
#
    collection: 'discharges'
}
