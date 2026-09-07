
// Define collection and schema for CreativeFile
export interface CreativeFile {
    uri:
	type : URL
    fileSizeKB:
	type : number
    mimeType:
	type : string
    checksum:
	type : string
    CreativeAsset:
	type : Schema.Types.ObjectId
#
    collection: 'creativeFiles'
}
