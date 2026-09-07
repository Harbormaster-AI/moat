package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PayoutDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePayout - creates a new db entry
//----------------------------------------------------------------------------
func CreatePayout(obj model.Payout)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Payout with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Payout", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePayout", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPayout - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPayout(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Payout

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Payout with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Payout using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Payout using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPayout", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPayout - returns all
//----------------------------------------------------------------------------
func GetAllPayout()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Payout

	//----------------------------------------------------------------------------
	// Request the ORM to find all Payout
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Payout" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Payout", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPayout", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePayout - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePayout(obj model.Payout)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Payout using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Payout using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePayout", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePayout - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePayout(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPayout(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Payout using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Payout using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePayout", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Merchant on a Payout
//----------------------------------------------------------------------------
func AssignMerchantToPayout( payoutId uint64, merchantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Merchant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Merchant with a
		// matching merchantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, merchantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Merchant	to the Payout
			//----------------------------------------------------------------------------
			parentObj.Merchant = &childObj

			//----------------------------------------------------------------------------
			// save the Payout
			//----------------------------------------------------------------------------
			return UpdatePayout(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Merchant", merchantId )
			return utils.RequestResult{false, msg, "assignMerchant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Merchant on a Payout
//----------------------------------------------------------------------------
func UnassignMerchantFromPayout(payoutId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// assign an empty Merchant to the Merchant
		//----------------------------------------------------------------------------
		parentObj.Merchant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Merchant
		//----------------------------------------------------------------------------
		parentObj.MerchantId = nil;

		//----------------------------------------------------------------------------
		// save the Payout
		//----------------------------------------------------------------------------
		return UpdatePayout(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SettlementBatch on a Payout
//----------------------------------------------------------------------------
func AssignSettlementBatchToPayout( payoutId uint64, settlementBatchId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SettlementBatch

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SettlementBatch with a
		// matching settlementBatchId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, settlementBatchId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SettlementBatch	to the Payout
			//----------------------------------------------------------------------------
			parentObj.SettlementBatch = &childObj

			//----------------------------------------------------------------------------
			// save the Payout
			//----------------------------------------------------------------------------
			return UpdatePayout(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SettlementBatch", settlementBatchId )
			return utils.RequestResult{false, msg, "assignSettlementBatch", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SettlementBatch on a Payout
//----------------------------------------------------------------------------
func UnassignSettlementBatchFromPayout(payoutId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// assign an empty SettlementBatch to the SettlementBatch
		//----------------------------------------------------------------------------
		parentObj.SettlementBatch = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SettlementBatch
		//----------------------------------------------------------------------------
		parentObj.SettlementBatchId = nil;

		//----------------------------------------------------------------------------
		// save the Payout
		//----------------------------------------------------------------------------
		return UpdatePayout(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a Payout
//----------------------------------------------------------------------------
func AssignDestinationAccountToPayout( payoutId uint64, destinationAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching destinationAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, destinationAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DestinationAccount	to the Payout
			//----------------------------------------------------------------------------
			parentObj.DestinationAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Payout
			//----------------------------------------------------------------------------
			return UpdatePayout(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )
			return utils.RequestResult{false, msg, "assignDestinationAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a Payout
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromPayout(payoutId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayout(payoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payout)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Payout
		//----------------------------------------------------------------------------
		return UpdatePayout(parentObj)

	} else {
		return parentRequestResult
	}

}


