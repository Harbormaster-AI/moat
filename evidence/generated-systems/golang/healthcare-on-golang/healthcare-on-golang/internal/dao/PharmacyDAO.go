package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PharmacyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePharmacy - creates a new db entry
//----------------------------------------------------------------------------
func CreatePharmacy(obj model.Pharmacy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Pharmacy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Pharmacy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePharmacy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPharmacy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPharmacy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Pharmacy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Pharmacy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Pharmacy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Pharmacy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPharmacy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPharmacy - returns all
//----------------------------------------------------------------------------
func GetAllPharmacy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Pharmacy

	//----------------------------------------------------------------------------
	// Request the ORM to find all Pharmacy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Pharmacy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Pharmacy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPharmacy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePharmacy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePharmacy(obj model.Pharmacy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Pharmacy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Pharmacy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePharmacy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePharmacy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePharmacy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPharmacy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Pharmacy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Pharmacy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Pharmacy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePharmacy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Facility on a Pharmacy
//----------------------------------------------------------------------------
func AssignFacilityToPharmacy( pharmacyId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

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
			// assign the Facility	to the Pharmacy
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Pharmacy
			//----------------------------------------------------------------------------
			return UpdatePharmacy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Pharmacy
//----------------------------------------------------------------------------
func UnassignFacilityFromPharmacy(pharmacyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Pharmacy
		//----------------------------------------------------------------------------
		return UpdatePharmacy(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more medicationDispensesIds as a MedicationDispenses to a Pharmacy
//----------------------------------------------------------------------------
func AddMedicationDispensesToPharmacy ( pharmacyId uint64, medicationDispensesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

		// slice the ids on comma with no spaces
		ids := strings.Split( medicationDispensesIds, ",")

		for _, medicationDispensesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationDispense

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationDispense
			// with a matching medicationDispensesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , medicationDispensesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the MedicationDispenses using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MedicationDispenses").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MedicationDispenses", medicationDispensesId )
				return utils.RequestResult{false, msg, "unassignMedicationDispenses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Pharmacy from the gorm
		//----------------------------------------------------------------------------
		return GetPharmacy(pharmacyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more medicationDispensesIds as a MedicationDispenses from a Pharmacy
//----------------------------------------------------------------------------
func RemoveMedicationDispensesFromPharmacy( pharmacyId uint64, medicationDispensesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

		// slice the ids on comma with no spaces
		ids := strings.Split( medicationDispensesIds, ",")

		for _, medicationDispensesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationDispense

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationDispense
			// with a matching medicationDispensesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , medicationDispensesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MedicationDispenseObj from the MedicationDispenses array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MedicationDispenses").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MedicationDispenses", medicationDispensesId )
				return utils.RequestResult{false, msg, "removeMedicationDispenses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Pharmacy from the gorm
		//----------------------------------------------------------------------------
		return GetPharmacy(pharmacyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more medicationOrdersIds as a MedicationOrders to a Pharmacy
//----------------------------------------------------------------------------
func AddMedicationOrdersToPharmacy ( pharmacyId uint64, medicationOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

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
		// retrieve the modified Pharmacy from the gorm
		//----------------------------------------------------------------------------
		return GetPharmacy(pharmacyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more medicationOrdersIds as a MedicationOrders from a Pharmacy
//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromPharmacy( pharmacyId uint64, medicationOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Pharmacy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPharmacy(pharmacyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Pharmacy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Pharmacy)

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
		// retrieve the modified Pharmacy from the gorm
		//----------------------------------------------------------------------------
		return GetPharmacy(pharmacyId)

	} else {
		return parentRequestResult
	}
}

