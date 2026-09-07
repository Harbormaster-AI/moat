
// Define collection and schema for DeviceCriterion
export interface DeviceCriterion {
    TargetingProfile:
	type : Schema.Types.ObjectId
    DeviceType:
 	type : String
    PlatformType:
 	type : String
    Operator:
 	type : String
#
    collection: 'deviceCriterions'
}
