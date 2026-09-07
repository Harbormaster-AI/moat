package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ForecastLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateForecastLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateForecastLine(obj model.ForecastLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ForecastLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ForecastLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateForecastLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetForecastLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetForecastLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ForecastLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ForecastLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ForecastLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ForecastLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetForecastLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllForecastLine - returns all
//----------------------------------------------------------------------------
func GetAllForecastLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ForecastLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all ForecastLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ForecastLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ForecastLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllForecastLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateForecastLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateForecastLine(obj model.ForecastLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ForecastLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ForecastLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateForecastLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteForecastLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteForecastLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ForecastLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetForecastLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ForecastLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ForecastLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ForecastLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ForecastLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteForecastLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Forecast on a ForecastLine
//----------------------------------------------------------------------------
func AssignForecastToForecastLine( forecastLineId uint64, forecastId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ForecastLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecastLine(forecastLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ForecastLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ForecastLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Forecast

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Forecast with a
		// matching forecastId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, forecastId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Forecast	to the ForecastLine
			//----------------------------------------------------------------------------
			parentObj.Forecast = &childObj

			//----------------------------------------------------------------------------
			// save the ForecastLine
			//----------------------------------------------------------------------------
			return UpdateForecastLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Forecast", forecastId )
			return utils.RequestResult{false, msg, "assignForecast", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Forecast on a ForecastLine
//----------------------------------------------------------------------------
func UnassignForecastFromForecastLine(forecastLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ForecastLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecastLine(forecastLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ForecastLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ForecastLine)

		//----------------------------------------------------------------------------
		// assign an empty Forecast to the Forecast
		//----------------------------------------------------------------------------
		parentObj.Forecast = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Forecast
		//----------------------------------------------------------------------------
		parentObj.ForecastId = nil;

		//----------------------------------------------------------------------------
		// save the ForecastLine
		//----------------------------------------------------------------------------
		return UpdateForecastLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Item on a ForecastLine
//----------------------------------------------------------------------------
func AssignItemToForecastLine( forecastLineId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ForecastLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecastLine(forecastLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ForecastLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ForecastLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the ForecastLine
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the ForecastLine
			//----------------------------------------------------------------------------
			return UpdateForecastLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a ForecastLine
//----------------------------------------------------------------------------
func UnassignItemFromForecastLine(forecastLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ForecastLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecastLine(forecastLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ForecastLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ForecastLine)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the ForecastLine
		//----------------------------------------------------------------------------
		return UpdateForecastLine(parentObj)

	} else {
		return parentRequestResult
	}

}


