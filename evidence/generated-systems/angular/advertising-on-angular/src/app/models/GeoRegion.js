
// Define collection and schema for GeoRegion
export interface GeoRegion {
    code:
	type : string
    name:
	type : string
    Parent:
	type : Schema.Types.ObjectId
    Children:
 	type : [{ type: Schema.Types.ObjectId, ref: 'GeoRegion' }]
    RegionType:
 	type : String
#
    collection: 'geoRegions'
}
