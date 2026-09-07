package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EmploymentContractDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEmploymentContract - creates a new db entry
//----------------------------------------------------------------------------
func CreateEmploymentContract(obj model.EmploymentContract)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EmploymentContract with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EmploymentContract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEmploymentContract", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEmploymentContract - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEmploymentContract(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EmploymentContract

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EmploymentContract with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EmploymentContract using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EmploymentContract using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEmploymentContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEmploymentContract - returns all
//----------------------------------------------------------------------------
func GetAllEmploymentContract()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EmploymentContract

	//----------------------------------------------------------------------------
	// Request the ORM to find all EmploymentContract
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EmploymentContract" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EmploymentContract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEmploymentContract", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEmploymentContract - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEmploymentContract(obj model.EmploymentContract)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EmploymentContract using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EmploymentContract using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEmploymentContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEmploymentContract - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEmploymentContract(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEmploymentContract(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EmploymentContract using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EmploymentContract using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEmploymentContract", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a EmploymentContract
//----------------------------------------------------------------------------
func AssignEmployeeToEmploymentContract( employmentContractId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

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
			// assign the Employee	to the EmploymentContract
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentContract
			//----------------------------------------------------------------------------
			return UpdateEmploymentContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a EmploymentContract
//----------------------------------------------------------------------------
func UnassignEmployeeFromEmploymentContract(employmentContractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentContract
		//----------------------------------------------------------------------------
		return UpdateEmploymentContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CompensationPackage on a EmploymentContract
//----------------------------------------------------------------------------
func AssignCompensationPackageToEmploymentContract( employmentContractId uint64, compensationPackageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

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
			// assign the CompensationPackage	to the EmploymentContract
			//----------------------------------------------------------------------------
			parentObj.CompensationPackage = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentContract
			//----------------------------------------------------------------------------
			return UpdateEmploymentContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompensationPackage", compensationPackageId )
			return utils.RequestResult{false, msg, "assignCompensationPackage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CompensationPackage on a EmploymentContract
//----------------------------------------------------------------------------
func UnassignCompensationPackageFromEmploymentContract(employmentContractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// assign an empty CompensationPackage to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackageId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentContract
		//----------------------------------------------------------------------------
		return UpdateEmploymentContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkSchedule on a EmploymentContract
//----------------------------------------------------------------------------
func AssignWorkScheduleToEmploymentContract( employmentContractId uint64, workScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkSchedule with a
		// matching workScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkSchedule	to the EmploymentContract
			//----------------------------------------------------------------------------
			parentObj.WorkSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentContract
			//----------------------------------------------------------------------------
			return UpdateEmploymentContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkSchedule", workScheduleId )
			return utils.RequestResult{false, msg, "assignWorkSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkSchedule on a EmploymentContract
//----------------------------------------------------------------------------
func UnassignWorkScheduleFromEmploymentContract(employmentContractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// assign an empty WorkSchedule to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentContract
		//----------------------------------------------------------------------------
		return UpdateEmploymentContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a EmploymentContract
//----------------------------------------------------------------------------
func AssignLocationToEmploymentContract( employmentContractId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Location

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Location with a
		// matching locationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, locationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Location	to the EmploymentContract
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentContract
			//----------------------------------------------------------------------------
			return UpdateEmploymentContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a EmploymentContract
//----------------------------------------------------------------------------
func UnassignLocationFromEmploymentContract(employmentContractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// assign an empty Location to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentContract
		//----------------------------------------------------------------------------
		return UpdateEmploymentContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PayrollCalendar on a EmploymentContract
//----------------------------------------------------------------------------
func AssignPayrollCalendarToEmploymentContract( employmentContractId uint64, payrollCalendarId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PayrollCalendar

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PayrollCalendar with a
		// matching payrollCalendarId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, payrollCalendarId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PayrollCalendar	to the EmploymentContract
			//----------------------------------------------------------------------------
			parentObj.PayrollCalendar = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentContract
			//----------------------------------------------------------------------------
			return UpdateEmploymentContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollCalendar", payrollCalendarId )
			return utils.RequestResult{false, msg, "assignPayrollCalendar", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PayrollCalendar on a EmploymentContract
//----------------------------------------------------------------------------
func UnassignPayrollCalendarFromEmploymentContract(employmentContractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentContract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentContract(employmentContractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentContract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentContract)

		//----------------------------------------------------------------------------
		// assign an empty PayrollCalendar to the PayrollCalendar
		//----------------------------------------------------------------------------
		parentObj.PayrollCalendar = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PayrollCalendar
		//----------------------------------------------------------------------------
		parentObj.PayrollCalendarId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentContract
		//----------------------------------------------------------------------------
		return UpdateEmploymentContract(parentObj)

	} else {
		return parentRequestResult
	}

}


