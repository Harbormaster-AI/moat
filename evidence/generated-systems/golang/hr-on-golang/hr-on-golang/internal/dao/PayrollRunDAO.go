package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PayrollRunDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePayrollRun - creates a new db entry
//----------------------------------------------------------------------------
func CreatePayrollRun(obj model.PayrollRun)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PayrollRun with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PayrollRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePayrollRun", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPayrollRun - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPayrollRun(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PayrollRun

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PayrollRun with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PayrollRun using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PayrollRun using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPayrollRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPayrollRun - returns all
//----------------------------------------------------------------------------
func GetAllPayrollRun()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PayrollRun

	//----------------------------------------------------------------------------
	// Request the ORM to find all PayrollRun
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PayrollRun" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PayrollRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPayrollRun", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePayrollRun - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePayrollRun(obj model.PayrollRun)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PayrollRun using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PayrollRun using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePayrollRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePayrollRun - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePayrollRun(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PayrollRun with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPayrollRun(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PayrollRun)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PayrollRun using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PayrollRun using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePayrollRun", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PayrollCalendar on a PayrollRun
//----------------------------------------------------------------------------
func AssignPayrollCalendarToPayrollRun( payrollRunId uint64, payrollCalendarId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PayrollRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollRun(payrollRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollRun)

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
			// assign the PayrollCalendar	to the PayrollRun
			//----------------------------------------------------------------------------
			parentObj.PayrollCalendar = &childObj

			//----------------------------------------------------------------------------
			// save the PayrollRun
			//----------------------------------------------------------------------------
			return UpdatePayrollRun(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollCalendar", payrollCalendarId )
			return utils.RequestResult{false, msg, "assignPayrollCalendar", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PayrollCalendar on a PayrollRun
//----------------------------------------------------------------------------
func UnassignPayrollCalendarFromPayrollRun(payrollRunId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollRun(payrollRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollRun)

		//----------------------------------------------------------------------------
		// assign an empty PayrollCalendar to the PayrollCalendar
		//----------------------------------------------------------------------------
		parentObj.PayrollCalendar = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PayrollCalendar
		//----------------------------------------------------------------------------
		parentObj.PayrollCalendarId = nil;

		//----------------------------------------------------------------------------
		// save the PayrollRun
		//----------------------------------------------------------------------------
		return UpdatePayrollRun(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more payrollItemsIds as a PayrollItems to a PayrollRun
//----------------------------------------------------------------------------
func AddPayrollItemsToPayrollRun ( payrollRunId uint64, payrollItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PayrollRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollRun(payrollRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollItemsIds, ",")

		for _, payrollItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollItem
			// with a matching payrollItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PayrollItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollItems", payrollItemsId )
				return utils.RequestResult{false, msg, "unassignPayrollItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollRun from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollRun(payrollRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more payrollItemsIds as a PayrollItems from a PayrollRun
//----------------------------------------------------------------------------
func RemovePayrollItemsFromPayrollRun( payrollRunId uint64, payrollItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PayrollRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayrollRun(payrollRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PayrollRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PayrollRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollItemsIds, ",")

		for _, payrollItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollItem
			// with a matching payrollItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PayrollItemObj from the PayrollItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollItems", payrollItemsId )
				return utils.RequestResult{false, msg, "removePayrollItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PayrollRun from the gorm
		//----------------------------------------------------------------------------
		return GetPayrollRun(payrollRunId)

	} else {
		return parentRequestResult
	}
}

