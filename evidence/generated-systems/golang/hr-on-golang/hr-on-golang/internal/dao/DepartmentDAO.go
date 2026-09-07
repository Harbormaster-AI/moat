package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DepartmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDepartment - creates a new db entry
//----------------------------------------------------------------------------
func CreateDepartment(obj model.Department)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Department with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Department", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDepartment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDepartment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDepartment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Department

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Department with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Department using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Department using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDepartment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDepartment - returns all
//----------------------------------------------------------------------------
func GetAllDepartment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Department

	//----------------------------------------------------------------------------
	// Request the ORM to find all Department
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Department" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Department", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDepartment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDepartment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDepartment(obj model.Department)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Department using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Department using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDepartment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDepartment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDepartment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDepartment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Department)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Department using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Department using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDepartment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Department
//----------------------------------------------------------------------------
func AssignOrganizationToDepartment( departmentId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

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
			// assign the Organization	to the Department
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Department
			//----------------------------------------------------------------------------
			return UpdateDepartment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Department
//----------------------------------------------------------------------------
func UnassignOrganizationFromDepartment(departmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Department
		//----------------------------------------------------------------------------
		return UpdateDepartment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Manager on a Department
//----------------------------------------------------------------------------
func AssignManagerToDepartment( departmentId uint64, managerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

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
			// assign the Manager	to the Department
			//----------------------------------------------------------------------------
			parentObj.Manager = &childObj

			//----------------------------------------------------------------------------
			// save the Department
			//----------------------------------------------------------------------------
			return UpdateDepartment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manager", managerId )
			return utils.RequestResult{false, msg, "assignManager", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Manager on a Department
//----------------------------------------------------------------------------
func UnassignManagerFromDepartment(departmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Manager
		//----------------------------------------------------------------------------
		parentObj.Manager = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Manager
		//----------------------------------------------------------------------------
		parentObj.ManagerId = nil;

		//----------------------------------------------------------------------------
		// save the Department
		//----------------------------------------------------------------------------
		return UpdateDepartment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CostCenter on a Department
//----------------------------------------------------------------------------
func AssignCostCenterToDepartment( departmentId uint64, costCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

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
			// assign the CostCenter	to the Department
			//----------------------------------------------------------------------------
			parentObj.CostCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Department
			//----------------------------------------------------------------------------
			return UpdateDepartment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenter", costCenterId )
			return utils.RequestResult{false, msg, "assignCostCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CostCenter on a Department
//----------------------------------------------------------------------------
func UnassignCostCenterFromDepartment(departmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

		//----------------------------------------------------------------------------
		// assign an empty CostCenter to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Department
		//----------------------------------------------------------------------------
		return UpdateDepartment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more positionsIds as a Positions to a Department
//----------------------------------------------------------------------------
func AddPositionsToDepartment ( departmentId uint64, positionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

		// slice the ids on comma with no spaces
		ids := strings.Split( positionsIds, ",")

		for _, positionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching positionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , positionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Positions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Positions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Positions", positionsId )
				return utils.RequestResult{false, msg, "unassignPositions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Department from the gorm
		//----------------------------------------------------------------------------
		return GetDepartment(departmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more positionsIds as a Positions from a Department
//----------------------------------------------------------------------------
func RemovePositionsFromDepartment( departmentId uint64, positionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

		// slice the ids on comma with no spaces
		ids := strings.Split( positionsIds, ",")

		for _, positionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching positionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , positionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PositionObj from the Positions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Positions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Positions", positionsId )
				return utils.RequestResult{false, msg, "removePositions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Department from the gorm
		//----------------------------------------------------------------------------
		return GetDepartment(departmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more employeesIds as a Employees to a Department
//----------------------------------------------------------------------------
func AddEmployeesToDepartment ( departmentId uint64, employeesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

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
		// retrieve the modified Department from the gorm
		//----------------------------------------------------------------------------
		return GetDepartment(departmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more employeesIds as a Employees from a Department
//----------------------------------------------------------------------------
func RemoveEmployeesFromDepartment( departmentId uint64, employeesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Department with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDepartment(departmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Department so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Department)

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
		// retrieve the modified Department from the gorm
		//----------------------------------------------------------------------------
		return GetDepartment(departmentId)

	} else {
		return parentRequestResult
	}
}

