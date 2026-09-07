
// Define collection and schema for TrackingPixel
export interface TrackingPixel {
    name:
	type : string
    url:
	type : URL
    Campaign:
	type : Schema.Types.ObjectId
    Advertiser:
	type : Schema.Types.ObjectId
    ConversionEvents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ConversionEvent' }]
    EventType:
 	type : String
    PixelType:
 	type : String
#
    collection: 'trackingPixels'
}
