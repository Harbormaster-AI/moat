package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ObservationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateObservation - creates a new db entry
//----------------------------------------------------------------------------
func CreateObservation(obj model.Observation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Observation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Observation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateObservation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetObservation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetObservation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Observation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Observation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Observation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Observation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetObservation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllObservation - returns all
//----------------------------------------------------------------------------
func GetAllObservation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Observation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Observation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Observation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Observation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllObservation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateObservation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateObservation(obj model.Observation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Observation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Observation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateObservation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteObservation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteObservation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetObservation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Observation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Observation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteObservation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Encounter on a Observation
//----------------------------------------------------------------------------
func AssignEncounterToObservation( observationId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Encounter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Encounter with a
		// matching encounterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, encounterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Encounter	to the Observation
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Observation
			//----------------------------------------------------------------------------
			return UpdateObservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Observation
//----------------------------------------------------------------------------
func UnassignEncounterFromObservation(observationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Observation
		//----------------------------------------------------------------------------
		return UpdateObservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Patient on a Observation
//----------------------------------------------------------------------------
func AssignPatientToObservation( observationId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Patient

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Patient with a
		// matching patientId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, patientId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Patient	to the Observation
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Observation
			//----------------------------------------------------------------------------
			return UpdateObservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Observation
//----------------------------------------------------------------------------
func UnassignPatientFromObservation(observationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Observation
		//----------------------------------------------------------------------------
		return UpdateObservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Device on a Observation
//----------------------------------------------------------------------------
func AssignDeviceToObservation( observationId uint64, deviceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MedicalDevice

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MedicalDevice with a
		// matching deviceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, deviceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Device	to the Observation
			//----------------------------------------------------------------------------
			parentObj.Device = &childObj

			//----------------------------------------------------------------------------
			// save the Observation
			//----------------------------------------------------------------------------
			return UpdateObservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Device", deviceId )
			return utils.RequestResult{false, msg, "assignDevice", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Device on a Observation
//----------------------------------------------------------------------------
func UnassignDeviceFromObservation(observationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// assign an empty MedicalDevice to the Device
		//----------------------------------------------------------------------------
		parentObj.Device = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Device
		//----------------------------------------------------------------------------
		parentObj.DeviceId = nil;

		//----------------------------------------------------------------------------
		// save the Observation
		//----------------------------------------------------------------------------
		return UpdateObservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LabResult on a Observation
//----------------------------------------------------------------------------
func AssignLabResultToObservation( observationId uint64, labResultId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LabResult

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LabResult with a
		// matching labResultId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, labResultId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LabResult	to the Observation
			//----------------------------------------------------------------------------
			parentObj.LabResult = &childObj

			//----------------------------------------------------------------------------
			// save the Observation
			//----------------------------------------------------------------------------
			return UpdateObservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LabResult", labResultId )
			return utils.RequestResult{false, msg, "assignLabResult", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LabResult on a Observation
//----------------------------------------------------------------------------
func UnassignLabResultFromObservation(observationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Observation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObservation(observationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Observation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Observation)

		//----------------------------------------------------------------------------
		// assign an empty LabResult to the LabResult
		//----------------------------------------------------------------------------
		parentObj.LabResult = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LabResult
		//----------------------------------------------------------------------------
		parentObj.LabResultId = nil;

		//----------------------------------------------------------------------------
		// save the Observation
		//----------------------------------------------------------------------------
		return UpdateObservation(parentObj)

	} else {
		return parentRequestResult
	}

}


