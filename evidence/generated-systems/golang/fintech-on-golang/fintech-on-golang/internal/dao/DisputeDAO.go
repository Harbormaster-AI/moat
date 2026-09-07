package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DisputeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDispute - creates a new db entry
//----------------------------------------------------------------------------
func CreateDispute(obj model.Dispute)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Dispute with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Dispute", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDispute", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDispute - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDispute(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Dispute

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Dispute with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Dispute using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Dispute using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDispute", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDispute - returns all
//----------------------------------------------------------------------------
func GetAllDispute()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Dispute

	//----------------------------------------------------------------------------
	// Request the ORM to find all Dispute
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Dispute" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dispute", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDispute", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDispute - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDispute(obj model.Dispute)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Dispute using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Dispute using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDispute", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDispute - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDispute(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDispute(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Dispute using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Dispute using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDispute", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func AssignTransactionToDispute( disputeId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

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
			// assign the Transaction	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a Dispute
//----------------------------------------------------------------------------
func UnassignTransactionFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Card on a Dispute
//----------------------------------------------------------------------------
func AssignCardToDispute( disputeId uint64, cardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PaymentCard

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentCard with a
		// matching cardId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, cardId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Card	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Card = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Card", cardId )
			return utils.RequestResult{false, msg, "assignCard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Card on a Dispute
//----------------------------------------------------------------------------
func UnassignCardFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the Card
		//----------------------------------------------------------------------------
		parentObj.Card = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Card
		//----------------------------------------------------------------------------
		parentObj.CardId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Merchant on a Dispute
//----------------------------------------------------------------------------
func AssignMerchantToDispute( disputeId uint64, merchantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

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
			// assign the Merchant	to the Dispute
			//----------------------------------------------------------------------------
			parentObj.Merchant = &childObj

			//----------------------------------------------------------------------------
			// save the Dispute
			//----------------------------------------------------------------------------
			return UpdateDispute(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Merchant", merchantId )
			return utils.RequestResult{false, msg, "assignMerchant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Merchant on a Dispute
//----------------------------------------------------------------------------
func UnassignMerchantFromDispute(disputeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		//----------------------------------------------------------------------------
		// assign an empty Merchant to the Merchant
		//----------------------------------------------------------------------------
		parentObj.Merchant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Merchant
		//----------------------------------------------------------------------------
		parentObj.MerchantId = nil;

		//----------------------------------------------------------------------------
		// save the Dispute
		//----------------------------------------------------------------------------
		return UpdateDispute(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more chargebacksIds as a Chargebacks to a Dispute
//----------------------------------------------------------------------------
func AddChargebacksToDispute ( disputeId uint64, chargebacksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		// slice the ids on comma with no spaces
		ids := strings.Split( chargebacksIds, ",")

		for _, chargebacksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Chargeback

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Chargeback
			// with a matching chargebacksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , chargebacksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Chargebacks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Chargebacks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Chargebacks", chargebacksId )
				return utils.RequestResult{false, msg, "unassignChargebacks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dispute from the gorm
		//----------------------------------------------------------------------------
		return GetDispute(disputeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more chargebacksIds as a Chargebacks from a Dispute
//----------------------------------------------------------------------------
func RemoveChargebacksFromDispute( disputeId uint64, chargebacksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dispute with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispute(disputeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dispute so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dispute)

		// slice the ids on comma with no spaces
		ids := strings.Split( chargebacksIds, ",")

		for _, chargebacksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Chargeback

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Chargeback
			// with a matching chargebacksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , chargebacksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ChargebackObj from the Chargebacks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Chargebacks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Chargebacks", chargebacksId )
				return utils.RequestResult{false, msg, "removeChargebacks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dispute from the gorm
		//----------------------------------------------------------------------------
		return GetDispute(disputeId)

	} else {
		return parentRequestResult
	}
}

