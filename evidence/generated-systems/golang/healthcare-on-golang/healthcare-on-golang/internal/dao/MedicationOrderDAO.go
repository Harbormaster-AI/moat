package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MedicationOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMedicationOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateMedicationOrder(obj model.MedicationOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MedicationOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MedicationOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMedicationOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMedicationOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMedicationOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MedicationOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MedicationOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MedicationOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MedicationOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMedicationOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMedicationOrder - returns all
//----------------------------------------------------------------------------
func GetAllMedicationOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MedicationOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all MedicationOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MedicationOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MedicationOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMedicationOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMedicationOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMedicationOrder(obj model.MedicationOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MedicationOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MedicationOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMedicationOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMedicationOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMedicationOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMedicationOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MedicationOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MedicationOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MedicationOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMedicationOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Order on a MedicationOrder
//----------------------------------------------------------------------------
func AssignOrderToMedicationOrder( medicationOrderId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ClinicalOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ClinicalOrder with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the MedicationOrder
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the MedicationOrder
			//----------------------------------------------------------------------------
			return UpdateMedicationOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a MedicationOrder
//----------------------------------------------------------------------------
func UnassignOrderFromMedicationOrder(medicationOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		//----------------------------------------------------------------------------
		// assign an empty ClinicalOrder to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the MedicationOrder
		//----------------------------------------------------------------------------
		return UpdateMedicationOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Pharmacy on a MedicationOrder
//----------------------------------------------------------------------------
func AssignPharmacyToMedicationOrder( medicationOrderId uint64, pharmacyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Pharmacy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Pharmacy with a
		// matching pharmacyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, pharmacyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Pharmacy	to the MedicationOrder
			//----------------------------------------------------------------------------
			parentObj.Pharmacy = &childObj

			//----------------------------------------------------------------------------
			// save the MedicationOrder
			//----------------------------------------------------------------------------
			return UpdateMedicationOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pharmacy", pharmacyId )
			return utils.RequestResult{false, msg, "assignPharmacy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Pharmacy on a MedicationOrder
//----------------------------------------------------------------------------
func UnassignPharmacyFromMedicationOrder(medicationOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		//----------------------------------------------------------------------------
		// assign an empty Pharmacy to the Pharmacy
		//----------------------------------------------------------------------------
		parentObj.Pharmacy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Pharmacy
		//----------------------------------------------------------------------------
		parentObj.PharmacyId = nil;

		//----------------------------------------------------------------------------
		// save the MedicationOrder
		//----------------------------------------------------------------------------
		return UpdateMedicationOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more dispensesIds as a Dispenses to a MedicationOrder
//----------------------------------------------------------------------------
func AddDispensesToMedicationOrder ( medicationOrderId uint64, dispensesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( dispensesIds, ",")

		for _, dispensesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationDispense

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationDispense
			// with a matching dispensesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dispensesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dispenses using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dispenses").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dispenses", dispensesId )
				return utils.RequestResult{false, msg, "unassignDispenses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicationOrder from the gorm
		//----------------------------------------------------------------------------
		return GetMedicationOrder(medicationOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dispensesIds as a Dispenses from a MedicationOrder
//----------------------------------------------------------------------------
func RemoveDispensesFromMedicationOrder( medicationOrderId uint64, dispensesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MedicationOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationOrder(medicationOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( dispensesIds, ",")

		for _, dispensesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicationDispense

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicationDispense
			// with a matching dispensesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dispensesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MedicationDispenseObj from the Dispenses array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dispenses").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dispenses", dispensesId )
				return utils.RequestResult{false, msg, "removeDispenses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicationOrder from the gorm
		//----------------------------------------------------------------------------
		return GetMedicationOrder(medicationOrderId)

	} else {
		return parentRequestResult
	}
}

