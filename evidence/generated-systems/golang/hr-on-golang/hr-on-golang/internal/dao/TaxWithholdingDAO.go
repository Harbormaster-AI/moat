package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TaxWithholdingDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTaxWithholding - creates a new db entry
//----------------------------------------------------------------------------
func CreateTaxWithholding(obj model.TaxWithholding)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TaxWithholding with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TaxWithholding", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTaxWithholding", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTaxWithholding - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTaxWithholding(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TaxWithholding

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TaxWithholding with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TaxWithholding using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TaxWithholding using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTaxWithholding", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTaxWithholding - returns all
//----------------------------------------------------------------------------
func GetAllTaxWithholding()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TaxWithholding

	//----------------------------------------------------------------------------
	// Request the ORM to find all TaxWithholding
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TaxWithholding" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TaxWithholding", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTaxWithholding", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTaxWithholding - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTaxWithholding(obj model.TaxWithholding)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TaxWithholding using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TaxWithholding using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTaxWithholding", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTaxWithholding - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTaxWithholding(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TaxWithholding with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTaxWithholding(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TaxWithholding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TaxWithholding)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TaxWithholding using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TaxWithholding using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTaxWithholding", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a TaxWithholding
//----------------------------------------------------------------------------
func AssignEmployeeToTaxWithholding( taxWithholdingId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TaxWithholding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTaxWithholding(taxWithholdingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TaxWithholding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TaxWithholding)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the TaxWithholding
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the TaxWithholding
			//----------------------------------------------------------------------------
			return UpdateTaxWithholding(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a TaxWithholding
//----------------------------------------------------------------------------
func UnassignEmployeeFromTaxWithholding(taxWithholdingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TaxWithholding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTaxWithholding(taxWithholdingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TaxWithholding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TaxWithholding)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the TaxWithholding
		//----------------------------------------------------------------------------
		return UpdateTaxWithholding(parentObj)

	} else {
		return parentRequestResult
	}

}


