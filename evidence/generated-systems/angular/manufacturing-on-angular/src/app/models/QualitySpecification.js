
// Define collection and schema for QualitySpecification
export interface QualitySpecification {
    specCode:
	type : string
    name:
	type : string
    version:
	type : string
    Item:
	type : Schema.Types.ObjectId
#
    collection: 'qualitySpecifications'
}
