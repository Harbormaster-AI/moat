package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CreativeFileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCreativeFile - creates a new db entry
//----------------------------------------------------------------------------
func CreateCreativeFile(obj model.CreativeFile)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a CreativeFile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CreativeFile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCreativeFile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCreativeFile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCreativeFile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CreativeFile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CreativeFile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CreativeFile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CreativeFile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCreativeFile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCreativeFile - returns all
//----------------------------------------------------------------------------
func GetAllCreativeFile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CreativeFile

	//----------------------------------------------------------------------------
	// Request the ORM to find all CreativeFile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CreativeFile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CreativeFile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCreativeFile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCreativeFile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCreativeFile(obj model.CreativeFile)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a CreativeFile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CreativeFile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCreativeFile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCreativeFile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCreativeFile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CreativeFile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCreativeFile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeFile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CreativeFile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CreativeFile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CreativeFile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCreativeFile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CreativeAsset on a CreativeFile
//----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeFile( creativeFileId uint64, creativeAssetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CreativeFile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeFile(creativeFileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeFile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeFile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CreativeAsset

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CreativeAsset with a
		// matching creativeAssetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, creativeAssetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CreativeAsset	to the CreativeFile
			//----------------------------------------------------------------------------
			parentObj.CreativeAsset = &childObj

			//----------------------------------------------------------------------------
			// save the CreativeFile
			//----------------------------------------------------------------------------
			return UpdateCreativeFile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeAsset", creativeAssetId )
			return utils.RequestResult{false, msg, "assignCreativeAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CreativeAsset on a CreativeFile
//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeFile(creativeFileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeFile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeFile(creativeFileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeFile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeFile)

		//----------------------------------------------------------------------------
		// assign an empty CreativeAsset to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAsset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAssetId = nil;

		//----------------------------------------------------------------------------
		// save the CreativeFile
		//----------------------------------------------------------------------------
		return UpdateCreativeFile(parentObj)

	} else {
		return parentRequestResult
	}

}


