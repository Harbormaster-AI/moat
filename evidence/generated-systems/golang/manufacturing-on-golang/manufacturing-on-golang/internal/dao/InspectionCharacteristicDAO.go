package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InspectionCharacteristicDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInspectionCharacteristic - creates a new db entry
//----------------------------------------------------------------------------
func CreateInspectionCharacteristic(obj model.InspectionCharacteristic)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InspectionCharacteristic with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InspectionCharacteristic", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInspectionCharacteristic", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInspectionCharacteristic - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInspectionCharacteristic(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InspectionCharacteristic

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InspectionCharacteristic with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InspectionCharacteristic using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InspectionCharacteristic using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInspectionCharacteristic", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInspectionCharacteristic - returns all
//----------------------------------------------------------------------------
func GetAllInspectionCharacteristic()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InspectionCharacteristic

	//----------------------------------------------------------------------------
	// Request the ORM to find all InspectionCharacteristic
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InspectionCharacteristic" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InspectionCharacteristic", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInspectionCharacteristic", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInspectionCharacteristic - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInspectionCharacteristic(obj model.InspectionCharacteristic)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InspectionCharacteristic using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InspectionCharacteristic using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInspectionCharacteristic", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInspectionCharacteristic - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInspectionCharacteristic(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InspectionCharacteristic with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInspectionCharacteristic(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionCharacteristic so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InspectionCharacteristic)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InspectionCharacteristic using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InspectionCharacteristic using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInspectionCharacteristic", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a InspectionPlan on a InspectionCharacteristic
//----------------------------------------------------------------------------
func AssignInspectionPlanToInspectionCharacteristic( inspectionCharacteristicId uint64, inspectionPlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionCharacteristic with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionCharacteristic(inspectionCharacteristicId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionCharacteristic so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionCharacteristic)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InspectionPlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InspectionPlan with a
		// matching inspectionPlanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inspectionPlanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InspectionPlan	to the InspectionCharacteristic
			//----------------------------------------------------------------------------
			parentObj.InspectionPlan = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionCharacteristic
			//----------------------------------------------------------------------------
			return UpdateInspectionCharacteristic(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InspectionPlan", inspectionPlanId )
			return utils.RequestResult{false, msg, "assignInspectionPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InspectionPlan on a InspectionCharacteristic
//----------------------------------------------------------------------------
func UnassignInspectionPlanFromInspectionCharacteristic(inspectionCharacteristicId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionCharacteristic with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionCharacteristic(inspectionCharacteristicId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionCharacteristic so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionCharacteristic)

		//----------------------------------------------------------------------------
		// assign an empty InspectionPlan to the InspectionPlan
		//----------------------------------------------------------------------------
		parentObj.InspectionPlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InspectionPlan
		//----------------------------------------------------------------------------
		parentObj.InspectionPlanId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionCharacteristic
		//----------------------------------------------------------------------------
		return UpdateInspectionCharacteristic(parentObj)

	} else {
		return parentRequestResult
	}

}


