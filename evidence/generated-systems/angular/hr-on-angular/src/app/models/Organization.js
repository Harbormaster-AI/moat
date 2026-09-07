
// Define collection and schema for Organization
export interface Organization {
    name:
	type : string
    legalName:
	type : string
    registrationCountry:
	type : string
    website:
	type : string
    Departments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Department' }]
    Locations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Location' }]
    JobFamilies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'JobFamily' }]
    BenefitPlans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BenefitPlan' }]
    CostCenters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CostCenter' }]
    PayrollCalendars:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PayrollCalendar' }]
#
    collection: 'organizations'
}
