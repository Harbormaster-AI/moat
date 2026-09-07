package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftFamilyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftFamily - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftFamily(obj model.AircraftFamily)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftFamily with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftFamily", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftFamily", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftFamily - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftFamily(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftFamily

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftFamily with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftFamily using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftFamily using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftFamily", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftFamily - returns all
//----------------------------------------------------------------------------
func GetAllAircraftFamily()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftFamily

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftFamily
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftFamily" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftFamily", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftFamily", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftFamily - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftFamily(obj model.AircraftFamily)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftFamily using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftFamily using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftFamily", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftFamily - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftFamily(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftFamily with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftFamily(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftFamily)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftFamily using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftFamily using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftFamily", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Program on a AircraftFamily
//----------------------------------------------------------------------------
func AssignProgramToAircraftFamily( aircraftFamilyId uint64, programId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftFamily(aircraftFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftFamily)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftProgram

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftProgram with a
		// matching programId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, programId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Program	to the AircraftFamily
			//----------------------------------------------------------------------------
			parentObj.Program = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftFamily
			//----------------------------------------------------------------------------
			return UpdateAircraftFamily(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Program", programId )
			return utils.RequestResult{false, msg, "assignProgram", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Program on a AircraftFamily
//----------------------------------------------------------------------------
func UnassignProgramFromAircraftFamily(aircraftFamilyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftFamily(aircraftFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftFamily)

		//----------------------------------------------------------------------------
		// assign an empty AircraftProgram to the Program
		//----------------------------------------------------------------------------
		parentObj.Program = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Program
		//----------------------------------------------------------------------------
		parentObj.ProgramId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftFamily
		//----------------------------------------------------------------------------
		return UpdateAircraftFamily(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more aircraftModelsIds as a AircraftModels to a AircraftFamily
//----------------------------------------------------------------------------
func AddAircraftModelsToAircraftFamily ( aircraftFamilyId uint64, aircraftModelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftFamily(aircraftFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftFamily)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftModelsIds, ",")

		for _, aircraftModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftModel
			// with a matching aircraftModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AircraftModels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftModels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftModels", aircraftModelsId )
				return utils.RequestResult{false, msg, "unassignAircraftModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftFamily from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftFamily(aircraftFamilyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more aircraftModelsIds as a AircraftModels from a AircraftFamily
//----------------------------------------------------------------------------
func RemoveAircraftModelsFromAircraftFamily( aircraftFamilyId uint64, aircraftModelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftFamily(aircraftFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftFamily)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftModelsIds, ",")

		for _, aircraftModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftModel
			// with a matching aircraftModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftModelObj from the AircraftModels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftModels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftModels", aircraftModelsId )
				return utils.RequestResult{false, msg, "removeAircraftModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftFamily from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftFamily(aircraftFamilyId)

	} else {
		return parentRequestResult
	}
}

