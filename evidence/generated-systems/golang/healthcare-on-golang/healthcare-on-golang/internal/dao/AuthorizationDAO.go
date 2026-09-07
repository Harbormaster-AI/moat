package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AuthorizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAuthorization - creates a new db entry
//----------------------------------------------------------------------------
func CreateAuthorization(obj model.Authorization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Authorization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Authorization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAuthorization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAuthorization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAuthorization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Authorization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Authorization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Authorization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Authorization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAuthorization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAuthorization - returns all
//----------------------------------------------------------------------------
func GetAllAuthorization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Authorization

	//----------------------------------------------------------------------------
	// Request the ORM to find all Authorization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Authorization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Authorization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAuthorization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAuthorization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAuthorization(obj model.Authorization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Authorization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Authorization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAuthorization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAuthorization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAuthorization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Authorization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAuthorization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Authorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Authorization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Authorization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Authorization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAuthorization", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Coverage on a Authorization
//----------------------------------------------------------------------------
func AssignCoverageToAuthorization( authorizationId uint64, coverageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Authorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuthorization(authorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Authorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Authorization)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Coverage

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Coverage with a
		// matching coverageId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, coverageId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Coverage	to the Authorization
			//----------------------------------------------------------------------------
			parentObj.Coverage = &childObj

			//----------------------------------------------------------------------------
			// save the Authorization
			//----------------------------------------------------------------------------
			return UpdateAuthorization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverage", coverageId )
			return utils.RequestResult{false, msg, "assignCoverage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Coverage on a Authorization
//----------------------------------------------------------------------------
func UnassignCoverageFromAuthorization(authorizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Authorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuthorization(authorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Authorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Authorization)

		//----------------------------------------------------------------------------
		// assign an empty Coverage to the Coverage
		//----------------------------------------------------------------------------
		parentObj.Coverage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Coverage
		//----------------------------------------------------------------------------
		parentObj.CoverageId = nil;

		//----------------------------------------------------------------------------
		// save the Authorization
		//----------------------------------------------------------------------------
		return UpdateAuthorization(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Order on a Authorization
//----------------------------------------------------------------------------
func AssignOrderToAuthorization( authorizationId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Authorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuthorization(authorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Authorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Authorization)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ClinicalOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ClinicalOrder with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the Authorization
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the Authorization
			//----------------------------------------------------------------------------
			return UpdateAuthorization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a Authorization
//----------------------------------------------------------------------------
func UnassignOrderFromAuthorization(authorizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Authorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuthorization(authorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Authorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Authorization)

		//----------------------------------------------------------------------------
		// assign an empty ClinicalOrder to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the Authorization
		//----------------------------------------------------------------------------
		return UpdateAuthorization(parentObj)

	} else {
		return parentRequestResult
	}

}


