package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EmployeeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEmployee - creates a new db entry
//----------------------------------------------------------------------------
func CreateEmployee(obj model.Employee)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Employee with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Employee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEmployee", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEmployee - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEmployee(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Employee

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Employee with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Employee using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Employee using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEmployee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEmployee - returns all
//----------------------------------------------------------------------------
func GetAllEmployee()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Employee

	//----------------------------------------------------------------------------
	// Request the ORM to find all Employee
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Employee" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Employee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEmployee", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEmployee - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEmployee(obj model.Employee)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Employee using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Employee using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEmployee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEmployee - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEmployee(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEmployee(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Employee using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Employee using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEmployee", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Manager on a Employee
//----------------------------------------------------------------------------
func AssignManagerToEmployee( employeeId uint64, managerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching managerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, managerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Manager	to the Employee
			//----------------------------------------------------------------------------
			parentObj.Manager = &childObj

			//----------------------------------------------------------------------------
			// save the Employee
			//----------------------------------------------------------------------------
			return UpdateEmployee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manager", managerId )
			return utils.RequestResult{false, msg, "assignManager", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Manager on a Employee
//----------------------------------------------------------------------------
func UnassignManagerFromEmployee(employeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Manager
		//----------------------------------------------------------------------------
		parentObj.Manager = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Manager
		//----------------------------------------------------------------------------
		parentObj.ManagerId = nil;

		//----------------------------------------------------------------------------
		// save the Employee
		//----------------------------------------------------------------------------
		return UpdateEmployee(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Department on a Employee
//----------------------------------------------------------------------------
func AssignDepartmentToEmployee( employeeId uint64, departmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Department

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Department with a
		// matching departmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, departmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Department	to the Employee
			//----------------------------------------------------------------------------
			parentObj.Department = &childObj

			//----------------------------------------------------------------------------
			// save the Employee
			//----------------------------------------------------------------------------
			return UpdateEmployee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Department", departmentId )
			return utils.RequestResult{false, msg, "assignDepartment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Department on a Employee
//----------------------------------------------------------------------------
func UnassignDepartmentFromEmployee(employeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// assign an empty Department to the Department
		//----------------------------------------------------------------------------
		parentObj.Department = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Department
		//----------------------------------------------------------------------------
		parentObj.DepartmentId = nil;

		//----------------------------------------------------------------------------
		// save the Employee
		//----------------------------------------------------------------------------
		return UpdateEmployee(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PrimaryLocation on a Employee
//----------------------------------------------------------------------------
func AssignPrimaryLocationToEmployee( employeeId uint64, primaryLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Location

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Location with a
		// matching primaryLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, primaryLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PrimaryLocation	to the Employee
			//----------------------------------------------------------------------------
			parentObj.PrimaryLocation = &childObj

			//----------------------------------------------------------------------------
			// save the Employee
			//----------------------------------------------------------------------------
			return UpdateEmployee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PrimaryLocation", primaryLocationId )
			return utils.RequestResult{false, msg, "assignPrimaryLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PrimaryLocation on a Employee
//----------------------------------------------------------------------------
func UnassignPrimaryLocationFromEmployee(employeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// assign an empty Location to the PrimaryLocation
		//----------------------------------------------------------------------------
		parentObj.PrimaryLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PrimaryLocation
		//----------------------------------------------------------------------------
		parentObj.PrimaryLocationId = nil;

		//----------------------------------------------------------------------------
		// save the Employee
		//----------------------------------------------------------------------------
		return UpdateEmployee(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CostCenter on a Employee
//----------------------------------------------------------------------------
func AssignCostCenterToEmployee( employeeId uint64, costCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CostCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CostCenter with a
		// matching costCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, costCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CostCenter	to the Employee
			//----------------------------------------------------------------------------
			parentObj.CostCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Employee
			//----------------------------------------------------------------------------
			return UpdateEmployee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenter", costCenterId )
			return utils.RequestResult{false, msg, "assignCostCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CostCenter on a Employee
//----------------------------------------------------------------------------
func UnassignCostCenterFromEmployee(employeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// assign an empty CostCenter to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Employee
		//----------------------------------------------------------------------------
		return UpdateEmployee(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more directReportsIds as a DirectReports to a Employee
//----------------------------------------------------------------------------
func AddDirectReportsToEmployee ( employeeId uint64, directReportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( directReportsIds, ",")

		for _, directReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching directReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , directReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DirectReports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DirectReports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DirectReports", directReportsId )
				return utils.RequestResult{false, msg, "unassignDirectReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more directReportsIds as a DirectReports from a Employee
//----------------------------------------------------------------------------
func RemoveDirectReportsFromEmployee( employeeId uint64, directReportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( directReportsIds, ",")

		for _, directReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching directReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , directReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmployeeObj from the DirectReports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DirectReports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DirectReports", directReportsId )
				return utils.RequestResult{false, msg, "removeDirectReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more employmentAssignmentsIds as a EmploymentAssignments to a Employee
//----------------------------------------------------------------------------
func AddEmploymentAssignmentsToEmployee ( employeeId uint64, employmentAssignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( employmentAssignmentsIds, ",")

		for _, employmentAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment
			// with a matching employmentAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , employmentAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EmploymentAssignments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmploymentAssignments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmploymentAssignments", employmentAssignmentsId )
				return utils.RequestResult{false, msg, "unassignEmploymentAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more employmentAssignmentsIds as a EmploymentAssignments from a Employee
//----------------------------------------------------------------------------
func RemoveEmploymentAssignmentsFromEmployee( employeeId uint64, employmentAssignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( employmentAssignmentsIds, ",")

		for _, employmentAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment
			// with a matching employmentAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , employmentAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmploymentAssignmentObj from the EmploymentAssignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmploymentAssignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmploymentAssignments", employmentAssignmentsId )
				return utils.RequestResult{false, msg, "removeEmploymentAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a Employee
//----------------------------------------------------------------------------
func AddContractsToEmployee ( employeeId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentContract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "unassignContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a Employee
//----------------------------------------------------------------------------
func RemoveContractsFromEmployee( employeeId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentContract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmploymentContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more benefitEnrollmentsIds as a BenefitEnrollments to a Employee
//----------------------------------------------------------------------------
func AddBenefitEnrollmentsToEmployee ( employeeId uint64, benefitEnrollmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( benefitEnrollmentsIds, ",")

		for _, benefitEnrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment
			// with a matching benefitEnrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , benefitEnrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BenefitEnrollments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BenefitEnrollments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitEnrollments", benefitEnrollmentsId )
				return utils.RequestResult{false, msg, "unassignBenefitEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more benefitEnrollmentsIds as a BenefitEnrollments from a Employee
//----------------------------------------------------------------------------
func RemoveBenefitEnrollmentsFromEmployee( employeeId uint64, benefitEnrollmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( benefitEnrollmentsIds, ",")

		for _, benefitEnrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment
			// with a matching benefitEnrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , benefitEnrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BenefitEnrollmentObj from the BenefitEnrollments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BenefitEnrollments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitEnrollments", benefitEnrollmentsId )
				return utils.RequestResult{false, msg, "removeBenefitEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more timesheetsIds as a Timesheets to a Employee
//----------------------------------------------------------------------------
func AddTimesheetsToEmployee ( employeeId uint64, timesheetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( timesheetsIds, ",")

		for _, timesheetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Timesheet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Timesheet
			// with a matching timesheetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , timesheetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Timesheets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Timesheets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Timesheets", timesheetsId )
				return utils.RequestResult{false, msg, "unassignTimesheets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more timesheetsIds as a Timesheets from a Employee
//----------------------------------------------------------------------------
func RemoveTimesheetsFromEmployee( employeeId uint64, timesheetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( timesheetsIds, ",")

		for _, timesheetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Timesheet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Timesheet
			// with a matching timesheetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , timesheetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TimesheetObj from the Timesheets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Timesheets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Timesheets", timesheetsId )
				return utils.RequestResult{false, msg, "removeTimesheets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more leaveRequestsIds as a LeaveRequests to a Employee
//----------------------------------------------------------------------------
func AddLeaveRequestsToEmployee ( employeeId uint64, leaveRequestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( leaveRequestsIds, ",")

		for _, leaveRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LeaveRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LeaveRequest
			// with a matching leaveRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leaveRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LeaveRequests using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LeaveRequests").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeaveRequests", leaveRequestsId )
				return utils.RequestResult{false, msg, "unassignLeaveRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more leaveRequestsIds as a LeaveRequests from a Employee
//----------------------------------------------------------------------------
func RemoveLeaveRequestsFromEmployee( employeeId uint64, leaveRequestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( leaveRequestsIds, ",")

		for _, leaveRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LeaveRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LeaveRequest
			// with a matching leaveRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leaveRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LeaveRequestObj from the LeaveRequests array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LeaveRequests").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeaveRequests", leaveRequestsId )
				return utils.RequestResult{false, msg, "removeLeaveRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more performanceReviewsIds as a PerformanceReviews to a Employee
//----------------------------------------------------------------------------
func AddPerformanceReviewsToEmployee ( employeeId uint64, performanceReviewsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceReviewsIds, ",")

		for _, performanceReviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceReview
			// with a matching performanceReviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceReviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PerformanceReviews using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceReviews").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceReviews", performanceReviewsId )
				return utils.RequestResult{false, msg, "unassignPerformanceReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more performanceReviewsIds as a PerformanceReviews from a Employee
//----------------------------------------------------------------------------
func RemovePerformanceReviewsFromEmployee( employeeId uint64, performanceReviewsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceReviewsIds, ",")

		for _, performanceReviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceReview
			// with a matching performanceReviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceReviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PerformanceReviewObj from the PerformanceReviews array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceReviews").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceReviews", performanceReviewsId )
				return utils.RequestResult{false, msg, "removePerformanceReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more trainingEnrollmentsIds as a TrainingEnrollments to a Employee
//----------------------------------------------------------------------------
func AddTrainingEnrollmentsToEmployee ( employeeId uint64, trainingEnrollmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( trainingEnrollmentsIds, ",")

		for _, trainingEnrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingEnrollment
			// with a matching trainingEnrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingEnrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TrainingEnrollments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingEnrollments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingEnrollments", trainingEnrollmentsId )
				return utils.RequestResult{false, msg, "unassignTrainingEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more trainingEnrollmentsIds as a TrainingEnrollments from a Employee
//----------------------------------------------------------------------------
func RemoveTrainingEnrollmentsFromEmployee( employeeId uint64, trainingEnrollmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( trainingEnrollmentsIds, ",")

		for _, trainingEnrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingEnrollment
			// with a matching trainingEnrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingEnrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrainingEnrollmentObj from the TrainingEnrollments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingEnrollments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingEnrollments", trainingEnrollmentsId )
				return utils.RequestResult{false, msg, "removeTrainingEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more workAuthorizationsIds as a WorkAuthorizations to a Employee
//----------------------------------------------------------------------------
func AddWorkAuthorizationsToEmployee ( employeeId uint64, workAuthorizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( workAuthorizationsIds, ",")

		for _, workAuthorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkAuthorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkAuthorization
			// with a matching workAuthorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workAuthorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkAuthorizations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkAuthorizations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkAuthorizations", workAuthorizationsId )
				return utils.RequestResult{false, msg, "unassignWorkAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workAuthorizationsIds as a WorkAuthorizations from a Employee
//----------------------------------------------------------------------------
func RemoveWorkAuthorizationsFromEmployee( employeeId uint64, workAuthorizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( workAuthorizationsIds, ",")

		for _, workAuthorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkAuthorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkAuthorization
			// with a matching workAuthorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workAuthorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkAuthorizationObj from the WorkAuthorizations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkAuthorizations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkAuthorizations", workAuthorizationsId )
				return utils.RequestResult{false, msg, "removeWorkAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

