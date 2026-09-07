
// Define collection and schema for Facility
export interface Facility {
    name:
	type : string
    facilityCode:
	type : string
    address:
	type : Address
    HealthSystem:
	type : Schema.Types.ObjectId
    Departments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Department' }]
    CareTeams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CareTeam' }]
    Laboratories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Laboratory' }]
    ImagingCenters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingCenter' }]
    Pharmacies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Pharmacy' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    FacilityType:
 	type : String
#
    collection: 'facilitys'
}
