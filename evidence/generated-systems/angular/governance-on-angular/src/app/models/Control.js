
// Define collection and schema for Control
export interface Control {
    name:
	type : string
    objective:
	type : string
    ownerDepartment:
	type : string
    Policy:
	type : Schema.Types.ObjectId
    ControlTests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ControlTest_' }]
    Evidence:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Evidence' }]
    Risks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Risk' }]
    Obligations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Obligation' }]
    Procedures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Procedure' }]
    Issues:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Issue' }]
    ControlType:
 	type : String
    Frequency:
 	type : String
    Status:
 	type : String
#
    collection: 'controls'
}
