package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PurchaseAgreementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePurchaseAgreement - creates a new db entry
//----------------------------------------------------------------------------
func CreatePurchaseAgreement(obj model.PurchaseAgreement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PurchaseAgreement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PurchaseAgreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePurchaseAgreement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPurchaseAgreement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPurchaseAgreement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PurchaseAgreement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PurchaseAgreement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PurchaseAgreement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PurchaseAgreement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPurchaseAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPurchaseAgreement - returns all
//----------------------------------------------------------------------------
func GetAllPurchaseAgreement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PurchaseAgreement

	//----------------------------------------------------------------------------
	// Request the ORM to find all PurchaseAgreement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PurchaseAgreement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PurchaseAgreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPurchaseAgreement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePurchaseAgreement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePurchaseAgreement(obj model.PurchaseAgreement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PurchaseAgreement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PurchaseAgreement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePurchaseAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePurchaseAgreement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePurchaseAgreement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PurchaseAgreement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPurchaseAgreement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PurchaseAgreement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PurchaseAgreement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PurchaseAgreement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePurchaseAgreement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AircraftOrder on a PurchaseAgreement
//----------------------------------------------------------------------------
func AssignAircraftOrderToPurchaseAgreement( purchaseAgreementId uint64, aircraftOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PurchaseAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseAgreement(purchaseAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseAgreement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftOrder with a
		// matching aircraftOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AircraftOrder	to the PurchaseAgreement
			//----------------------------------------------------------------------------
			parentObj.AircraftOrder = &childObj

			//----------------------------------------------------------------------------
			// save the PurchaseAgreement
			//----------------------------------------------------------------------------
			return UpdatePurchaseAgreement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftOrder", aircraftOrderId )
			return utils.RequestResult{false, msg, "assignAircraftOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AircraftOrder on a PurchaseAgreement
//----------------------------------------------------------------------------
func UnassignAircraftOrderFromPurchaseAgreement(purchaseAgreementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseAgreement(purchaseAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseAgreement)

		//----------------------------------------------------------------------------
		// assign an empty AircraftOrder to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrderId = nil;

		//----------------------------------------------------------------------------
		// save the PurchaseAgreement
		//----------------------------------------------------------------------------
		return UpdatePurchaseAgreement(parentObj)

	} else {
		return parentRequestResult
	}

}


