package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SalaryComponentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSalaryComponent - creates a new db entry
//----------------------------------------------------------------------------
func CreateSalaryComponent(obj model.SalaryComponent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SalaryComponent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SalaryComponent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSalaryComponent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSalaryComponent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSalaryComponent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SalaryComponent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SalaryComponent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SalaryComponent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SalaryComponent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSalaryComponent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSalaryComponent - returns all
//----------------------------------------------------------------------------
func GetAllSalaryComponent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SalaryComponent

	//----------------------------------------------------------------------------
	// Request the ORM to find all SalaryComponent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SalaryComponent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SalaryComponent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSalaryComponent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSalaryComponent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSalaryComponent(obj model.SalaryComponent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SalaryComponent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SalaryComponent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSalaryComponent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSalaryComponent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSalaryComponent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SalaryComponent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSalaryComponent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalaryComponent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SalaryComponent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SalaryComponent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SalaryComponent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSalaryComponent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CompensationPackage on a SalaryComponent
//----------------------------------------------------------------------------
func AssignCompensationPackageToSalaryComponent( salaryComponentId uint64, compensationPackageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalaryComponent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalaryComponent(salaryComponentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalaryComponent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalaryComponent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CompensationPackage

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CompensationPackage with a
		// matching compensationPackageId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, compensationPackageId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CompensationPackage	to the SalaryComponent
			//----------------------------------------------------------------------------
			parentObj.CompensationPackage = &childObj

			//----------------------------------------------------------------------------
			// save the SalaryComponent
			//----------------------------------------------------------------------------
			return UpdateSalaryComponent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompensationPackage", compensationPackageId )
			return utils.RequestResult{false, msg, "assignCompensationPackage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CompensationPackage on a SalaryComponent
//----------------------------------------------------------------------------
func UnassignCompensationPackageFromSalaryComponent(salaryComponentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalaryComponent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalaryComponent(salaryComponentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalaryComponent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalaryComponent)

		//----------------------------------------------------------------------------
		// assign an empty CompensationPackage to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackageId = nil;

		//----------------------------------------------------------------------------
		// save the SalaryComponent
		//----------------------------------------------------------------------------
		return UpdateSalaryComponent(parentObj)

	} else {
		return parentRequestResult
	}

}


