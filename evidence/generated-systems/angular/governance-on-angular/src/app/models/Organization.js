
// Define collection and schema for Organization
export interface Organization {
    name:
	type : string
    legalName:
	type : string
    jurisdiction:
	type : string
    industrySector:
	type : string
    GovernanceBodies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'GovernanceBody' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Risks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Risk' }]
    ThirdParties:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ThirdParty' }]
    RecordsRepositories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RecordsRepository' }]
    DataProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    CompliancePrograms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceProgram' }]
    AuditPrograms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditProgram' }]
    BusinessUnits:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessUnit' }]
    Matters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Matter' }]
    DataBreaches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataBreach' }]
#
    collection: 'organizations'
}
