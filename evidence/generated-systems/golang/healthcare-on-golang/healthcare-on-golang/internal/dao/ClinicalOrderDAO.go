package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClinicalOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClinicalOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateClinicalOrder(obj model.ClinicalOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ClinicalOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ClinicalOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClinicalOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClinicalOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClinicalOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ClinicalOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ClinicalOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ClinicalOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ClinicalOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClinicalOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClinicalOrder - returns all
//----------------------------------------------------------------------------
func GetAllClinicalOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ClinicalOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all ClinicalOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ClinicalOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ClinicalOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClinicalOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClinicalOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClinicalOrder(obj model.ClinicalOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ClinicalOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ClinicalOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClinicalOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClinicalOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClinicalOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClinicalOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ClinicalOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ClinicalOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ClinicalOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClinicalOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a ClinicalOrder
//----------------------------------------------------------------------------
func AssignPatientToClinicalOrder( clinicalOrderId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
			// assign the Patient	to the ClinicalOrder
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the ClinicalOrder
			//----------------------------------------------------------------------------
			return UpdateClinicalOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a ClinicalOrder
//----------------------------------------------------------------------------
func UnassignPatientFromClinicalOrder(clinicalOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the ClinicalOrder
		//----------------------------------------------------------------------------
		return UpdateClinicalOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Encounter on a ClinicalOrder
//----------------------------------------------------------------------------
func AssignEncounterToClinicalOrder( clinicalOrderId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
			// assign the Encounter	to the ClinicalOrder
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the ClinicalOrder
			//----------------------------------------------------------------------------
			return UpdateClinicalOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a ClinicalOrder
//----------------------------------------------------------------------------
func UnassignEncounterFromClinicalOrder(clinicalOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the ClinicalOrder
		//----------------------------------------------------------------------------
		return UpdateClinicalOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a OrderingClinician on a ClinicalOrder
//----------------------------------------------------------------------------
func AssignOrderingClinicianToClinicalOrder( clinicalOrderId uint64, orderingClinicianId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Clinician

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Clinician with a
		// matching orderingClinicianId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderingClinicianId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the OrderingClinician	to the ClinicalOrder
			//----------------------------------------------------------------------------
			parentObj.OrderingClinician = &childObj

			//----------------------------------------------------------------------------
			// save the ClinicalOrder
			//----------------------------------------------------------------------------
			return UpdateClinicalOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OrderingClinician", orderingClinicianId )
			return utils.RequestResult{false, msg, "assignOrderingClinician", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a OrderingClinician on a ClinicalOrder
//----------------------------------------------------------------------------
func UnassignOrderingClinicianFromClinicalOrder(clinicalOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the OrderingClinician
		//----------------------------------------------------------------------------
		parentObj.OrderingClinician = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the OrderingClinician
		//----------------------------------------------------------------------------
		parentObj.OrderingClinicianId = nil;

		//----------------------------------------------------------------------------
		// save the ClinicalOrder
		//----------------------------------------------------------------------------
		return UpdateClinicalOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more medicationOrdersIds as a MedicationOrders to a ClinicalOrder
//----------------------------------------------------------------------------
func AddMedicationOrdersToClinicalOrder ( clinicalOrderId uint64, medicationOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more medicationOrdersIds as a MedicationOrders from a ClinicalOrder
//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromClinicalOrder( clinicalOrderId uint64, medicationOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more laboratoryOrdersIds as a LaboratoryOrders to a ClinicalOrder
//----------------------------------------------------------------------------
func AddLaboratoryOrdersToClinicalOrder ( clinicalOrderId uint64, laboratoryOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoryOrdersIds, ",")

		for _, laboratoryOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching laboratoryOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoryOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LaboratoryOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LaboratoryOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LaboratoryOrders", laboratoryOrdersId )
				return utils.RequestResult{false, msg, "unassignLaboratoryOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more laboratoryOrdersIds as a LaboratoryOrders from a ClinicalOrder
//----------------------------------------------------------------------------
func RemoveLaboratoryOrdersFromClinicalOrder( clinicalOrderId uint64, laboratoryOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoryOrdersIds, ",")

		for _, laboratoryOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching laboratoryOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoryOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LaboratoryOrderObj from the LaboratoryOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LaboratoryOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LaboratoryOrders", laboratoryOrdersId )
				return utils.RequestResult{false, msg, "removeLaboratoryOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more imagingOrdersIds as a ImagingOrders to a ClinicalOrder
//----------------------------------------------------------------------------
func AddImagingOrdersToClinicalOrder ( clinicalOrderId uint64, imagingOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingOrdersIds as a ImagingOrders from a ClinicalOrder
//----------------------------------------------------------------------------
func RemoveImagingOrdersFromClinicalOrder( clinicalOrderId uint64, imagingOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

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
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more procedureOrdersIds as a ProcedureOrders to a ClinicalOrder
//----------------------------------------------------------------------------
func AddProcedureOrdersToClinicalOrder ( clinicalOrderId uint64, procedureOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( procedureOrdersIds, ",")

		for _, procedureOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProcedureOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProcedureOrder
			// with a matching procedureOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , procedureOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProcedureOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcedureOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcedureOrders", procedureOrdersId )
				return utils.RequestResult{false, msg, "unassignProcedureOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more procedureOrdersIds as a ProcedureOrders from a ClinicalOrder
//----------------------------------------------------------------------------
func RemoveProcedureOrdersFromClinicalOrder( clinicalOrderId uint64, procedureOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( procedureOrdersIds, ",")

		for _, procedureOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProcedureOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProcedureOrder
			// with a matching procedureOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , procedureOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProcedureOrderObj from the ProcedureOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcedureOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcedureOrders", procedureOrdersId )
				return utils.RequestResult{false, msg, "removeProcedureOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more authorizationsIds as a Authorizations to a ClinicalOrder
//----------------------------------------------------------------------------
func AddAuthorizationsToClinicalOrder ( clinicalOrderId uint64, authorizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizationsIds, ",")

		for _, authorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Authorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Authorization
			// with a matching authorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Authorizations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Authorizations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Authorizations", authorizationsId )
				return utils.RequestResult{false, msg, "unassignAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more authorizationsIds as a Authorizations from a ClinicalOrder
//----------------------------------------------------------------------------
func RemoveAuthorizationsFromClinicalOrder( clinicalOrderId uint64, authorizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ClinicalOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinicalOrder(clinicalOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClinicalOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClinicalOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizationsIds, ",")

		for _, authorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Authorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Authorization
			// with a matching authorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuthorizationObj from the Authorizations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Authorizations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Authorizations", authorizationsId )
				return utils.RequestResult{false, msg, "removeAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ClinicalOrder from the gorm
		//----------------------------------------------------------------------------
		return GetClinicalOrder(clinicalOrderId)

	} else {
		return parentRequestResult
	}
}

