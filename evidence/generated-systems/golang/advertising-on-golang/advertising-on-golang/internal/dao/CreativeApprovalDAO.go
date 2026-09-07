package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CreativeApprovalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCreativeApproval - creates a new db entry
//----------------------------------------------------------------------------
func CreateCreativeApproval(obj model.CreativeApproval)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CreativeApproval with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CreativeApproval", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCreativeApproval", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCreativeApproval - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCreativeApproval(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CreativeApproval

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CreativeApproval with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CreativeApproval using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CreativeApproval using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCreativeApproval", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCreativeApproval - returns all
//----------------------------------------------------------------------------
func GetAllCreativeApproval()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CreativeApproval

	//----------------------------------------------------------------------------
	// Request the ORM to find all CreativeApproval
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CreativeApproval" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CreativeApproval", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCreativeApproval", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCreativeApproval - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCreativeApproval(obj model.CreativeApproval)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CreativeApproval using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CreativeApproval using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCreativeApproval", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCreativeApproval - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCreativeApproval(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CreativeApproval with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCreativeApproval(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeApproval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CreativeApproval)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CreativeApproval using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CreativeApproval using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCreativeApproval", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CreativeAsset on a CreativeApproval
//----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeApproval( creativeApprovalId uint64, creativeAssetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CreativeApproval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeApproval(creativeApprovalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeApproval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeApproval)

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
			// assign the CreativeAsset	to the CreativeApproval
			//----------------------------------------------------------------------------
			parentObj.CreativeAsset = &childObj

			//----------------------------------------------------------------------------
			// save the CreativeApproval
			//----------------------------------------------------------------------------
			return UpdateCreativeApproval(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeAsset", creativeAssetId )
			return utils.RequestResult{false, msg, "assignCreativeAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CreativeAsset on a CreativeApproval
//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeApproval(creativeApprovalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeApproval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeApproval(creativeApprovalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeApproval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeApproval)

		//----------------------------------------------------------------------------
		// assign an empty CreativeAsset to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAsset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAssetId = nil;

		//----------------------------------------------------------------------------
		// save the CreativeApproval
		//----------------------------------------------------------------------------
		return UpdateCreativeApproval(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Publisher on a CreativeApproval
//----------------------------------------------------------------------------
func AssignPublisherToCreativeApproval( creativeApprovalId uint64, publisherId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CreativeApproval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeApproval(creativeApprovalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeApproval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeApproval)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Publisher

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Publisher with a
		// matching publisherId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, publisherId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Publisher	to the CreativeApproval
			//----------------------------------------------------------------------------
			parentObj.Publisher = &childObj

			//----------------------------------------------------------------------------
			// save the CreativeApproval
			//----------------------------------------------------------------------------
			return UpdateCreativeApproval(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Publisher", publisherId )
			return utils.RequestResult{false, msg, "assignPublisher", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Publisher on a CreativeApproval
//----------------------------------------------------------------------------
func UnassignPublisherFromCreativeApproval(creativeApprovalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeApproval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeApproval(creativeApprovalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeApproval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeApproval)

		//----------------------------------------------------------------------------
		// assign an empty Publisher to the Publisher
		//----------------------------------------------------------------------------
		parentObj.Publisher = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Publisher
		//----------------------------------------------------------------------------
		parentObj.PublisherId = nil;

		//----------------------------------------------------------------------------
		// save the CreativeApproval
		//----------------------------------------------------------------------------
		return UpdateCreativeApproval(parentObj)

	} else {
		return parentRequestResult
	}

}


