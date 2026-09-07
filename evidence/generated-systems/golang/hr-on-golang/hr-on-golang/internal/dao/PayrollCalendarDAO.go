package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PayrollCalendarDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePayrollCalendar - creates a new db entry
//----------------------------------------------------------------------------
func CreatePayrollCalendar(obj model.PayrollCalendar)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PayrollCalendar with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PayrollCalendar", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePayrollCalendar", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPayrollCalendar - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPayrollCalendar(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PayrollCalendar

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PayrollCalendar with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PayrollCalendar using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PayrollCalendar using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPayrollCalendar", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPayrollCalendar - returns all
//----------------------------------------------------------------------------
func GetAllPayrollCalendar()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PayrollCalendar

	//----------------------------------------------------------------------------
	// Request the ORM to find all PayrollCalendar
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PayrollCalendar" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PayrollCalendar", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPayrollCalendar", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePayrollCalendar - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePayrollCalendar(obj model.PayrollCalendar)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PayrollCalendar using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PayrollCalendar using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePayrollCalendar", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePayrollCalendar - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePayrollCalendar(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPayrollCalendar(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PayrollCalendar)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PayrollCalendar using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PayrollCalendar using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePayrollCalendar", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a PayrollCalendar
//----------------------------------------------------------------------------
func AssignOrganizationToPayrollCalendar( payrollCalendarId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the PayrollCalendar
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the PayrollCalendar
			//----------------------------------------------------------------------------
			return UpdatePayrollCalendar(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a PayrollCalendar
//----------------------------------------------------------------------------
func UnassignOrganizationFromPayrollCalendar(payrollCalendarId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the PayrollCalendar
		//----------------------------------------------------------------------------
		return UpdatePayrollCalendar(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more payrollRunsIds as a PayrollRuns to a PayrollCalendar
//----------------------------------------------------------------------------
func AddPayrollRunsToPayrollCalendar ( payrollCalendarId uint64, payrollRunsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollRunsIds, ",")

		for _, payrollRunsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollRun

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollRun
			// with a matching payrollRunsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollRunsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PayrollRuns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollRuns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollRuns", payrollRunsId )
				return utils.RequestResult{false, msg, "unassignPayrollRuns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollCalendar from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollCalendar(payrollCalendarId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more payrollRunsIds as a PayrollRuns from a PayrollCalendar
//----------------------------------------------------------------------------
func RemovePayrollRunsFromPayrollCalendar( payrollCalendarId uint64, payrollRunsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollRunsIds, ",")

		for _, payrollRunsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollRun

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollRun
			// with a matching payrollRunsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollRunsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PayrollRunObj from the PayrollRuns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollRuns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollRuns", payrollRunsId )
				return utils.RequestResult{false, msg, "removePayrollRuns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollCalendar from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollCalendar(payrollCalendarId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more employeesIds as a Employees to a PayrollCalendar
//----------------------------------------------------------------------------
func AddEmployeesToPayrollCalendar ( payrollCalendarId uint64, employeesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		// slice the ids on comma with no spaces
		ids := strings.Split( employeesIds, ",")

		for _, employeesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching employeesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , employeesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Employees using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Employees").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employees", employeesId )
				return utils.RequestResult{false, msg, "unassignEmployees", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollCalendar from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollCalendar(payrollCalendarId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more employeesIds as a Employees from a PayrollCalendar
//----------------------------------------------------------------------------
func RemoveEmployeesFromPayrollCalendar( payrollCalendarId uint64, employeesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PayrollCalendar with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollCalendar(payrollCalendarId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollCalendar so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollCalendar)

		// slice the ids on comma with no spaces
		ids := strings.Split( employeesIds, ",")

		for _, employeesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching employeesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , employeesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmployeeObj from the Employees array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Employees").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employees", employeesId )
				return utils.RequestResult{false, msg, "removeEmployees", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollCalendar from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollCalendar(payrollCalendarId)

	} else {
		return parentRequestResult
	}
}

