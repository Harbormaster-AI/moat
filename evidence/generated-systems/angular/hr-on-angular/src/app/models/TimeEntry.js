
// Define collection and schema for TimeEntry
export interface TimeEntry {
    entryDate:
	type : Date
    hoursWorked:
	type : String
    Timesheet:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    CostCenter:
	type : Schema.Types.ObjectId
    EntryType:
 	type : String
#
    collection: 'timeEntrys'
}
