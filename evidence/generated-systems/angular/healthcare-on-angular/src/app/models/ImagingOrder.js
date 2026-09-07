
// Define collection and schema for ImagingOrder
export interface ImagingOrder {
    bodySite:
	type : string
    contrast:
	type : boolean
    Order:
	type : Schema.Types.ObjectId
    ImagingCenter:
	type : Schema.Types.ObjectId
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ImagingReport' }]
    Modality:
 	type : String
#
    collection: 'imagingOrders'
}
