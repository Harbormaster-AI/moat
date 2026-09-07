package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ChargebackDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateChargeback - creates a new db entry
//----------------------------------------------------------------------------
func CreateChargeback(obj model.Chargeback)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Chargeback with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Chargeback", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateChargeback", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetChargeback - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetChargeback(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Chargeback

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Chargeback with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Chargeback using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Chargeback using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetChargeback", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllChargeback - returns all
//----------------------------------------------------------------------------
func GetAllChargeback()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Chargeback

	//----------------------------------------------------------------------------
	// Request the ORM to find all Chargeback
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Chargeback" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Chargeback", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllChargeback", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateChargeback - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateChargeback(obj model.Chargeback)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Chargeback using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Chargeback using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateChargeback", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteChargeback - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteChargeback(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Chargeback with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetChargeback(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Chargeback so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Chargeback)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Chargeback using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Chargeback using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteChargeback", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Dispute on a Chargeback
//----------------------------------------------------------------------------
func AssignDisputeToChargeback( chargebackId uint64, disputeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Chargeback with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetChargeback(chargebackId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Chargeback so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Chargeback)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Dispute

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Dispute with a
		// matching disputeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, disputeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dispute	to the Chargeback
			//----------------------------------------------------------------------------
			parentObj.Dispute = &childObj

			//----------------------------------------------------------------------------
			// save the Chargeback
			//----------------------------------------------------------------------------
			return UpdateChargeback(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dispute", disputeId )
			return utils.RequestResult{false, msg, "assignDispute", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dispute on a Chargeback
//----------------------------------------------------------------------------
func UnassignDisputeFromChargeback(chargebackId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Chargeback with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetChargeback(chargebackId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Chargeback so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Chargeback)

		//----------------------------------------------------------------------------
		// assign an empty Dispute to the Dispute
		//----------------------------------------------------------------------------
		parentObj.Dispute = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dispute
		//----------------------------------------------------------------------------
		parentObj.DisputeId = nil;

		//----------------------------------------------------------------------------
		// save the Chargeback
		//----------------------------------------------------------------------------
		return UpdateChargeback(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a Chargeback
//----------------------------------------------------------------------------
func AssignTransactionToChargeback( chargebackId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Chargeback with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetChargeback(chargebackId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Chargeback so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Chargeback)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Transaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Transaction with a
		// matching transactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, transactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Transaction	to the Chargeback
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the Chargeback
			//----------------------------------------------------------------------------
			return UpdateChargeback(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a Chargeback
//----------------------------------------------------------------------------
func UnassignTransactionFromChargeback(chargebackId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Chargeback with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetChargeback(chargebackId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Chargeback so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Chargeback)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the Chargeback
		//----------------------------------------------------------------------------
		return UpdateChargeback(parentObj)

	} else {
		return parentRequestResult
	}

}


