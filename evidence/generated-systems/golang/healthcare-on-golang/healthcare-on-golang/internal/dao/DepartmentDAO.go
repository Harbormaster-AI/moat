package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
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
// assigns a Facility on a Department
//----------------------------------------------------------------------------
func AssignFacilityToDepartment( departmentId uint64, facilityId uint64 )(utils.RequestResult){

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
		var childObj model.Facility

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Facility with a
		// matching facilityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, facilityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Facility	to the Department
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Department
			//----------------------------------------------------------------------------
			return UpdateDepartment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Department
//----------------------------------------------------------------------------
func UnassignFacilityFromDepartment(departmentId uint64)(utils.RequestResult) {

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
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Department
		//----------------------------------------------------------------------------
		return UpdateDepartment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more careTeamsIds as a CareTeams to a Department
//----------------------------------------------------------------------------
func AddCareTeamsToDepartment ( departmentId uint64, careTeamsIds string )(utils.RequestResult) {

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
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CareTeams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "unassignCareTeams", childObj}
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
// removes one or more careTeamsIds as a CareTeams from a Department
//----------------------------------------------------------------------------
func RemoveCareTeamsFromDepartment( departmentId uint64, careTeamsIds string )(utils.RequestResult) {
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
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CareTeamObj from the CareTeams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "removeCareTeams", childObj}
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

