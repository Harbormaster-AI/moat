
// Define collection and schema for ImagingCenter
export interface ImagingCenter {
    name:
	type : string
    Facility:
	type : Schema.Types.ObjectId
    ImagingOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingOrder' }]
    ImagingReports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingReport' }]
#
    collection: 'imagingCenters'
}
