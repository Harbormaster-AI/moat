package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PayrollItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePayrollItem - creates a new db entry
//----------------------------------------------------------------------------
func CreatePayrollItem(obj model.PayrollItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PayrollItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PayrollItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePayrollItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPayrollItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPayrollItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PayrollItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PayrollItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PayrollItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PayrollItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPayrollItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPayrollItem - returns all
//----------------------------------------------------------------------------
func GetAllPayrollItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PayrollItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all PayrollItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PayrollItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PayrollItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPayrollItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePayrollItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePayrollItem(obj model.PayrollItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PayrollItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PayrollItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePayrollItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePayrollItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePayrollItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PayrollItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPayrollItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PayrollItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PayrollItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PayrollItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePayrollItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PayrollRun on a PayrollItem
//----------------------------------------------------------------------------
func AssignPayrollRunToPayrollItem( payrollItemId uint64, payrollRunId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PayrollItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollItem(payrollItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PayrollRun

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PayrollRun with a
		// matching payrollRunId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, payrollRunId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PayrollRun	to the PayrollItem
			//----------------------------------------------------------------------------
			parentObj.PayrollRun = &childObj

			//----------------------------------------------------------------------------
			// save the PayrollItem
			//----------------------------------------------------------------------------
			return UpdatePayrollItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollRun", payrollRunId )
			return utils.RequestResult{false, msg, "assignPayrollRun", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PayrollRun on a PayrollItem
//----------------------------------------------------------------------------
func UnassignPayrollRunFromPayrollItem(payrollItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollItem(payrollItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollItem)

		//----------------------------------------------------------------------------
		// assign an empty PayrollRun to the PayrollRun
		//----------------------------------------------------------------------------
		parentObj.PayrollRun = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PayrollRun
		//----------------------------------------------------------------------------
		parentObj.PayrollRunId = nil;

		//----------------------------------------------------------------------------
		// save the PayrollItem
		//----------------------------------------------------------------------------
		return UpdatePayrollItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a PayrollItem
//----------------------------------------------------------------------------
func AssignEmployeeToPayrollItem( payrollItemId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PayrollItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollItem(payrollItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollItem)

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
			// assign the Employee	to the PayrollItem
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the PayrollItem
			//----------------------------------------------------------------------------
			return UpdatePayrollItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a PayrollItem
//----------------------------------------------------------------------------
func UnassignEmployeeFromPayrollItem(payrollItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollItem(payrollItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollItem)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the PayrollItem
		//----------------------------------------------------------------------------
		return UpdatePayrollItem(parentObj)

	} else {
		return parentRequestResult
	}

}


