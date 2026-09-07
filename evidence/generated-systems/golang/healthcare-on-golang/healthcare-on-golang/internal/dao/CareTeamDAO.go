package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CareTeamDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCareTeam - creates a new db entry
//----------------------------------------------------------------------------
func CreateCareTeam(obj model.CareTeam)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CareTeam with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CareTeam", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCareTeam", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCareTeam - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCareTeam(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CareTeam

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CareTeam with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CareTeam using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CareTeam using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCareTeam", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCareTeam - returns all
//----------------------------------------------------------------------------
func GetAllCareTeam()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CareTeam

	//----------------------------------------------------------------------------
	// Request the ORM to find all CareTeam
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CareTeam" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CareTeam", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCareTeam", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCareTeam - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCareTeam(obj model.CareTeam)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CareTeam using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CareTeam using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCareTeam", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCareTeam - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCareTeam(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCareTeam(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CareTeam)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CareTeam using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CareTeam using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCareTeam", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Department on a CareTeam
//----------------------------------------------------------------------------
func AssignDepartmentToCareTeam( careTeamId uint64, departmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

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
			// assign the Department	to the CareTeam
			//----------------------------------------------------------------------------
			parentObj.Department = &childObj

			//----------------------------------------------------------------------------
			// save the CareTeam
			//----------------------------------------------------------------------------
			return UpdateCareTeam(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Department", departmentId )
			return utils.RequestResult{false, msg, "assignDepartment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Department on a CareTeam
//----------------------------------------------------------------------------
func UnassignDepartmentFromCareTeam(careTeamId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

		//----------------------------------------------------------------------------
		// assign an empty Department to the Department
		//----------------------------------------------------------------------------
		parentObj.Department = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Department
		//----------------------------------------------------------------------------
		parentObj.DepartmentId = nil;

		//----------------------------------------------------------------------------
		// save the CareTeam
		//----------------------------------------------------------------------------
		return UpdateCareTeam(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more cliniciansIds as a Clinicians to a CareTeam
//----------------------------------------------------------------------------
func AddCliniciansToCareTeam ( careTeamId uint64, cliniciansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

		// slice the ids on comma with no spaces
		ids := strings.Split( cliniciansIds, ",")

		for _, cliniciansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Clinician

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Clinician
			// with a matching cliniciansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cliniciansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Clinicians using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Clinicians").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Clinicians", cliniciansId )
				return utils.RequestResult{false, msg, "unassignClinicians", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CareTeam from the gorm
		//----------------------------------------------------------------------------
		return GetCareTeam(careTeamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more cliniciansIds as a Clinicians from a CareTeam
//----------------------------------------------------------------------------
func RemoveCliniciansFromCareTeam( careTeamId uint64, cliniciansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

		// slice the ids on comma with no spaces
		ids := strings.Split( cliniciansIds, ",")

		for _, cliniciansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Clinician

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Clinician
			// with a matching cliniciansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cliniciansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClinicianObj from the Clinicians array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Clinicians").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Clinicians", cliniciansId )
				return utils.RequestResult{false, msg, "removeClinicians", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CareTeam from the gorm
		//----------------------------------------------------------------------------
		return GetCareTeam(careTeamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more patientsIds as a Patients to a CareTeam
//----------------------------------------------------------------------------
func AddPatientsToCareTeam ( careTeamId uint64, patientsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

		// slice the ids on comma with no spaces
		ids := strings.Split( patientsIds, ",")

		for _, patientsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Patient

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Patient
			// with a matching patientsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , patientsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Patients using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Patients").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patients", patientsId )
				return utils.RequestResult{false, msg, "unassignPatients", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CareTeam from the gorm
		//----------------------------------------------------------------------------
		return GetCareTeam(careTeamId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more patientsIds as a Patients from a CareTeam
//----------------------------------------------------------------------------
func RemovePatientsFromCareTeam( careTeamId uint64, patientsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CareTeam with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTeam(careTeamId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTeam so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTeam)

		// slice the ids on comma with no spaces
		ids := strings.Split( patientsIds, ",")

		for _, patientsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Patient

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Patient
			// with a matching patientsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , patientsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PatientObj from the Patients array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Patients").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patients", patientsId )
				return utils.RequestResult{false, msg, "removePatients", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CareTeam from the gorm
		//----------------------------------------------------------------------------
		return GetCareTeam(careTeamId)

	} else {
		return parentRequestResult
	}
}

