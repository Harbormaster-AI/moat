package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WarrantyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWarranty - creates a new db entry
//----------------------------------------------------------------------------
func CreateWarranty(obj model.Warranty)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Warranty with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Warranty", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWarranty", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWarranty - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWarranty(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Warranty

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Warranty with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Warranty using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Warranty using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWarranty", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWarranty - returns all
//----------------------------------------------------------------------------
func GetAllWarranty()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Warranty

	//----------------------------------------------------------------------------
	// Request the ORM to find all Warranty
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Warranty" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Warranty", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWarranty", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWarranty - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWarranty(obj model.Warranty)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Warranty using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Warranty using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWarranty", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWarranty - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWarranty(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Warranty with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWarranty(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warranty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Warranty)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Warranty using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Warranty using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWarranty", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Aircraft on a Warranty
//----------------------------------------------------------------------------
func AssignAircraftToWarranty( warrantyId uint64, aircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Warranty with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarranty(warrantyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warranty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warranty)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Aircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Aircraft with a
		// matching aircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Aircraft	to the Warranty
			//----------------------------------------------------------------------------
			parentObj.Aircraft = &childObj

			//----------------------------------------------------------------------------
			// save the Warranty
			//----------------------------------------------------------------------------
			return UpdateWarranty(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
			return utils.RequestResult{false, msg, "assignAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Aircraft on a Warranty
//----------------------------------------------------------------------------
func UnassignAircraftFromWarranty(warrantyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warranty with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarranty(warrantyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warranty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warranty)

		//----------------------------------------------------------------------------
		// assign an empty Aircraft to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.Aircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.AircraftId = nil;

		//----------------------------------------------------------------------------
		// save the Warranty
		//----------------------------------------------------------------------------
		return UpdateWarranty(parentObj)

	} else {
		return parentRequestResult
	}

}


