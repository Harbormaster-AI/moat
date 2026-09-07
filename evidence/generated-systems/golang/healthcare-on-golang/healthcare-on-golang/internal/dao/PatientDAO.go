package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PatientDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePatient - creates a new db entry
//----------------------------------------------------------------------------
func CreatePatient(obj model.Patient)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Patient with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Patient", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePatient", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPatient - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPatient(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Patient

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Patient with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Patient using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Patient using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPatient", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPatient - returns all
//----------------------------------------------------------------------------
func GetAllPatient()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Patient

	//----------------------------------------------------------------------------
	// Request the ORM to find all Patient
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Patient" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Patient", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPatient", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePatient - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePatient(obj model.Patient)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Patient using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Patient using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePatient", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePatient - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePatient(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPatient(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Patient)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Patient using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Patient using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePatient", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more appointmentsIds as a Appointments to a Patient
//----------------------------------------------------------------------------
func AddAppointmentsToPatient ( patientId uint64, appointmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Appointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Appointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Appointments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "unassignAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more appointmentsIds as a Appointments from a Patient
//----------------------------------------------------------------------------
func RemoveAppointmentsFromPatient( patientId uint64, appointmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Appointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Appointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AppointmentObj from the Appointments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "removeAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more encountersIds as a Encounters to a Patient
//----------------------------------------------------------------------------
func AddEncountersToPatient ( patientId uint64, encountersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( encountersIds, ",")

		for _, encountersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Encounter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Encounter
			// with a matching encountersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , encountersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Encounters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Encounters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounters", encountersId )
				return utils.RequestResult{false, msg, "unassignEncounters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more encountersIds as a Encounters from a Patient
//----------------------------------------------------------------------------
func RemoveEncountersFromPatient( patientId uint64, encountersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( encountersIds, ",")

		for _, encountersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Encounter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Encounter
			// with a matching encountersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , encountersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EncounterObj from the Encounters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Encounters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounters", encountersId )
				return utils.RequestResult{false, msg, "removeEncounters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more carePlansIds as a CarePlans to a Patient
//----------------------------------------------------------------------------
func AddCarePlansToPatient ( patientId uint64, carePlansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( carePlansIds, ",")

		for _, carePlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CarePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CarePlan
			// with a matching carePlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , carePlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CarePlans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CarePlans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CarePlans", carePlansId )
				return utils.RequestResult{false, msg, "unassignCarePlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more carePlansIds as a CarePlans from a Patient
//----------------------------------------------------------------------------
func RemoveCarePlansFromPatient( patientId uint64, carePlansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( carePlansIds, ",")

		for _, carePlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CarePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CarePlan
			// with a matching carePlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , carePlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CarePlanObj from the CarePlans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CarePlans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CarePlans", carePlansId )
				return utils.RequestResult{false, msg, "removeCarePlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more allergiesIds as a Allergies to a Patient
//----------------------------------------------------------------------------
func AddAllergiesToPatient ( patientId uint64, allergiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( allergiesIds, ",")

		for _, allergiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Allergy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Allergy
			// with a matching allergiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , allergiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Allergies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Allergies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Allergies", allergiesId )
				return utils.RequestResult{false, msg, "unassignAllergies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more allergiesIds as a Allergies from a Patient
//----------------------------------------------------------------------------
func RemoveAllergiesFromPatient( patientId uint64, allergiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( allergiesIds, ",")

		for _, allergiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Allergy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Allergy
			// with a matching allergiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , allergiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AllergyObj from the Allergies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Allergies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Allergies", allergiesId )
				return utils.RequestResult{false, msg, "removeAllergies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more conditionsIds as a Conditions to a Patient
//----------------------------------------------------------------------------
func AddConditionsToPatient ( patientId uint64, conditionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( conditionsIds, ",")

		for _, conditionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Condition

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Condition
			// with a matching conditionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , conditionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Conditions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Conditions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Conditions", conditionsId )
				return utils.RequestResult{false, msg, "unassignConditions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more conditionsIds as a Conditions from a Patient
//----------------------------------------------------------------------------
func RemoveConditionsFromPatient( patientId uint64, conditionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( conditionsIds, ",")

		for _, conditionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Condition

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Condition
			// with a matching conditionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , conditionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConditionObj from the Conditions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Conditions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Conditions", conditionsId )
				return utils.RequestResult{false, msg, "removeConditions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more medicationOrdersIds as a MedicationOrders to a Patient
//----------------------------------------------------------------------------
func AddMedicationOrdersToPatient ( patientId uint64, medicationOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( medicationOrdersIds, ",")

		for _, medicationOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationOrder
			// with a matching medicationOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , medicationOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the MedicationOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MedicationOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MedicationOrders", medicationOrdersId )
				return utils.RequestResult{false, msg, "unassignMedicationOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more medicationOrdersIds as a MedicationOrders from a Patient
//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromPatient( patientId uint64, medicationOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( medicationOrdersIds, ",")

		for _, medicationOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationOrder
			// with a matching medicationOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , medicationOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MedicationOrderObj from the MedicationOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MedicationOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MedicationOrders", medicationOrdersId )
				return utils.RequestResult{false, msg, "removeMedicationOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more labOrdersIds as a LabOrders to a Patient
//----------------------------------------------------------------------------
func AddLabOrdersToPatient ( patientId uint64, labOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( labOrdersIds, ",")

		for _, labOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching labOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , labOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LabOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LabOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LabOrders", labOrdersId )
				return utils.RequestResult{false, msg, "unassignLabOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more labOrdersIds as a LabOrders from a Patient
//----------------------------------------------------------------------------
func RemoveLabOrdersFromPatient( patientId uint64, labOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( labOrdersIds, ",")

		for _, labOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching labOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , labOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LaboratoryOrderObj from the LabOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LabOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LabOrders", labOrdersId )
				return utils.RequestResult{false, msg, "removeLabOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more imagingOrdersIds as a ImagingOrders to a Patient
//----------------------------------------------------------------------------
func AddImagingOrdersToPatient ( patientId uint64, imagingOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingOrdersIds, ",")

		for _, imagingOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingOrder
			// with a matching imagingOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ImagingOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingOrders", imagingOrdersId )
				return utils.RequestResult{false, msg, "unassignImagingOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingOrdersIds as a ImagingOrders from a Patient
//----------------------------------------------------------------------------
func RemoveImagingOrdersFromPatient( patientId uint64, imagingOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingOrdersIds, ",")

		for _, imagingOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingOrder
			// with a matching imagingOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingOrderObj from the ImagingOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingOrders", imagingOrdersId )
				return utils.RequestResult{false, msg, "removeImagingOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more coveragesIds as a Coverages to a Patient
//----------------------------------------------------------------------------
func AddCoveragesToPatient ( patientId uint64, coveragesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Coverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Coverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Coverages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "unassignCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more coveragesIds as a Coverages from a Patient
//----------------------------------------------------------------------------
func RemoveCoveragesFromPatient( patientId uint64, coveragesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Coverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Coverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CoverageObj from the Coverages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "removeCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a Patient
//----------------------------------------------------------------------------
func AddClaimsToPatient ( patientId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Patient
//----------------------------------------------------------------------------
func RemoveClaimsFromPatient( patientId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more devicesIds as a Devices to a Patient
//----------------------------------------------------------------------------
func AddDevicesToPatient ( patientId uint64, devicesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( devicesIds, ",")

		for _, devicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicalDevice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicalDevice
			// with a matching devicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , devicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Devices using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Devices").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Devices", devicesId )
				return utils.RequestResult{false, msg, "unassignDevices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more devicesIds as a Devices from a Patient
//----------------------------------------------------------------------------
func RemoveDevicesFromPatient( patientId uint64, devicesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

		// slice the ids on comma with no spaces
		ids := strings.Split( devicesIds, ",")

		for _, devicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicalDevice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicalDevice
			// with a matching devicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , devicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MedicalDeviceObj from the Devices array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Devices").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Devices", devicesId )
				return utils.RequestResult{false, msg, "removeDevices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more observationsIds as a Observations to a Patient
//----------------------------------------------------------------------------
func AddObservationsToPatient ( patientId uint64, observationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

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
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more observationsIds as a Observations from a Patient
//----------------------------------------------------------------------------
func RemoveObservationsFromPatient( patientId uint64, observationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Patient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPatient(patientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Patient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Patient)

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
		// retrieve the modified Patient from the gorm
		//----------------------------------------------------------------------------
		return GetPatient(patientId)

	} else {
		return parentRequestResult
	}
}

