package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CreativeVariationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCreativeVariation - creates a new db entry
//----------------------------------------------------------------------------
func CreateCreativeVariation(obj model.CreativeVariation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CreativeVariation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CreativeVariation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCreativeVariation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCreativeVariation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCreativeVariation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CreativeVariation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CreativeVariation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CreativeVariation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CreativeVariation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCreativeVariation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCreativeVariation - returns all
//----------------------------------------------------------------------------
func GetAllCreativeVariation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CreativeVariation

	//----------------------------------------------------------------------------
	// Request the ORM to find all CreativeVariation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CreativeVariation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CreativeVariation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCreativeVariation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCreativeVariation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCreativeVariation(obj model.CreativeVariation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CreativeVariation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CreativeVariation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCreativeVariation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCreativeVariation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCreativeVariation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CreativeVariation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCreativeVariation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeVariation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CreativeVariation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CreativeVariation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CreativeVariation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCreativeVariation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CreativeAsset on a CreativeVariation
//----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeVariation( creativeVariationId uint64, creativeAssetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CreativeVariation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeVariation(creativeVariationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeVariation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeVariation)

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
			// assign the CreativeAsset	to the CreativeVariation
			//----------------------------------------------------------------------------
			parentObj.CreativeAsset = &childObj

			//----------------------------------------------------------------------------
			// save the CreativeVariation
			//----------------------------------------------------------------------------
			return UpdateCreativeVariation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeAsset", creativeAssetId )
			return utils.RequestResult{false, msg, "assignCreativeAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CreativeAsset on a CreativeVariation
//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeVariation(creativeVariationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeVariation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeVariation(creativeVariationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeVariation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeVariation)

		//----------------------------------------------------------------------------
		// assign an empty CreativeAsset to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAsset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAssetId = nil;

		//----------------------------------------------------------------------------
		// save the CreativeVariation
		//----------------------------------------------------------------------------
		return UpdateCreativeVariation(parentObj)

	} else {
		return parentRequestResult
	}

}


