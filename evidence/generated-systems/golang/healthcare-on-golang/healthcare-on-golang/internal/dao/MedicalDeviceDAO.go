package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MedicalDeviceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMedicalDevice - creates a new db entry
//----------------------------------------------------------------------------
func CreateMedicalDevice(obj model.MedicalDevice)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MedicalDevice with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MedicalDevice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMedicalDevice", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMedicalDevice - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMedicalDevice(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MedicalDevice

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MedicalDevice with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MedicalDevice using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MedicalDevice using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMedicalDevice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMedicalDevice - returns all
//----------------------------------------------------------------------------
func GetAllMedicalDevice()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MedicalDevice

	//----------------------------------------------------------------------------
	// Request the ORM to find all MedicalDevice
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MedicalDevice" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MedicalDevice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMedicalDevice", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMedicalDevice - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMedicalDevice(obj model.MedicalDevice)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MedicalDevice using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MedicalDevice using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMedicalDevice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMedicalDevice - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMedicalDevice(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMedicalDevice(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MedicalDevice)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MedicalDevice using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MedicalDevice using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMedicalDevice", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a MedicalDevice
//----------------------------------------------------------------------------
func AssignPatientToMedicalDevice( medicalDeviceId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

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
			// assign the Patient	to the MedicalDevice
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the MedicalDevice
			//----------------------------------------------------------------------------
			return UpdateMedicalDevice(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a MedicalDevice
//----------------------------------------------------------------------------
func UnassignPatientFromMedicalDevice(medicalDeviceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the MedicalDevice
		//----------------------------------------------------------------------------
		return UpdateMedicalDevice(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more observationsIds as a Observations to a MedicalDevice
//----------------------------------------------------------------------------
func AddObservationsToMedicalDevice ( medicalDeviceId uint64, observationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Observations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "unassignObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalDevice from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalDevice(medicalDeviceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more observationsIds as a Observations from a MedicalDevice
//----------------------------------------------------------------------------
func RemoveObservationsFromMedicalDevice( medicalDeviceId uint64, observationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ObservationObj from the Observations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "removeObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalDevice from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalDevice(medicalDeviceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more softwareUpdatesIds as a SoftwareUpdates to a MedicalDevice
//----------------------------------------------------------------------------
func AddSoftwareUpdatesToMedicalDevice ( medicalDeviceId uint64, softwareUpdatesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareUpdatesIds, ",")

		for _, softwareUpdatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareUpdate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareUpdate
			// with a matching softwareUpdatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareUpdatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SoftwareUpdates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareUpdates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareUpdates", softwareUpdatesId )
				return utils.RequestResult{false, msg, "unassignSoftwareUpdates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalDevice from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalDevice(medicalDeviceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more softwareUpdatesIds as a SoftwareUpdates from a MedicalDevice
//----------------------------------------------------------------------------
func RemoveSoftwareUpdatesFromMedicalDevice( medicalDeviceId uint64, softwareUpdatesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MedicalDevice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalDevice(medicalDeviceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalDevice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalDevice)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareUpdatesIds, ",")

		for _, softwareUpdatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareUpdate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareUpdate
			// with a matching softwareUpdatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareUpdatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SoftwareUpdateObj from the SoftwareUpdates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareUpdates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareUpdates", softwareUpdatesId )
				return utils.RequestResult{false, msg, "removeSoftwareUpdates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalDevice from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalDevice(medicalDeviceId)

	} else {
		return parentRequestResult
	}
}

