package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CostCenterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCostCenter - creates a new db entry
//----------------------------------------------------------------------------
func CreateCostCenter(obj model.CostCenter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CostCenter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CostCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCostCenter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCostCenter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCostCenter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CostCenter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CostCenter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CostCenter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CostCenter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCostCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCostCenter - returns all
//----------------------------------------------------------------------------
func GetAllCostCenter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CostCenter

	//----------------------------------------------------------------------------
	// Request the ORM to find all CostCenter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CostCenter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CostCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCostCenter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCostCenter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCostCenter(obj model.CostCenter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CostCenter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CostCenter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCostCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCostCenter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCostCenter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCostCenter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CostCenter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CostCenter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CostCenter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCostCenter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a CostCenter
//----------------------------------------------------------------------------
func AssignOrganizationToCostCenter( costCenterId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

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
			// assign the Organization	to the CostCenter
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the CostCenter
			//----------------------------------------------------------------------------
			return UpdateCostCenter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a CostCenter
//----------------------------------------------------------------------------
func UnassignOrganizationFromCostCenter(costCenterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the CostCenter
		//----------------------------------------------------------------------------
		return UpdateCostCenter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more departmentsIds as a Departments to a CostCenter
//----------------------------------------------------------------------------
func AddDepartmentsToCostCenter ( costCenterId uint64, departmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Departments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "unassignDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more departmentsIds as a Departments from a CostCenter
//----------------------------------------------------------------------------
func RemoveDepartmentsFromCostCenter( costCenterId uint64, departmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DepartmentObj from the Departments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "removeDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more positionsIds as a Positions to a CostCenter
//----------------------------------------------------------------------------
func AddPositionsToCostCenter ( costCenterId uint64, positionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

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
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more positionsIds as a Positions from a CostCenter
//----------------------------------------------------------------------------
func RemovePositionsFromCostCenter( costCenterId uint64, positionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

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
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more employeesIds as a Employees to a CostCenter
//----------------------------------------------------------------------------
func AddEmployeesToCostCenter ( costCenterId uint64, employeesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

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
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more employeesIds as a Employees from a CostCenter
//----------------------------------------------------------------------------
func RemoveEmployeesFromCostCenter( costCenterId uint64, employeesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CostCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCostCenter(costCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CostCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CostCenter)

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
		// retrieve the modified CostCenter from the gorm
		//----------------------------------------------------------------------------
		return GetCostCenter(costCenterId)

	} else {
		return parentRequestResult
	}
}

