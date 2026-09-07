package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SettlementBatchDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSettlementBatch - creates a new db entry
//----------------------------------------------------------------------------
func CreateSettlementBatch(obj model.SettlementBatch)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SettlementBatch with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SettlementBatch", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSettlementBatch", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSettlementBatch - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSettlementBatch(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SettlementBatch

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SettlementBatch with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SettlementBatch using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SettlementBatch using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSettlementBatch", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSettlementBatch - returns all
//----------------------------------------------------------------------------
func GetAllSettlementBatch()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SettlementBatch

	//----------------------------------------------------------------------------
	// Request the ORM to find all SettlementBatch
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SettlementBatch" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SettlementBatch", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSettlementBatch", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSettlementBatch - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSettlementBatch(obj model.SettlementBatch)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SettlementBatch using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SettlementBatch using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSettlementBatch", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSettlementBatch - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSettlementBatch(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSettlementBatch(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SettlementBatch)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SettlementBatch using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SettlementBatch using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSettlementBatch", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Processor on a SettlementBatch
//----------------------------------------------------------------------------
func AssignProcessorToSettlementBatch( settlementBatchId uint64, processorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PaymentProcessor

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentProcessor with a
		// matching processorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, processorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Processor	to the SettlementBatch
			//----------------------------------------------------------------------------
			parentObj.Processor = &childObj

			//----------------------------------------------------------------------------
			// save the SettlementBatch
			//----------------------------------------------------------------------------
			return UpdateSettlementBatch(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Processor", processorId )
			return utils.RequestResult{false, msg, "assignProcessor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Processor on a SettlementBatch
//----------------------------------------------------------------------------
func UnassignProcessorFromSettlementBatch(settlementBatchId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		//----------------------------------------------------------------------------
		// assign an empty PaymentProcessor to the Processor
		//----------------------------------------------------------------------------
		parentObj.Processor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Processor
		//----------------------------------------------------------------------------
		parentObj.ProcessorId = nil;

		//----------------------------------------------------------------------------
		// save the SettlementBatch
		//----------------------------------------------------------------------------
		return UpdateSettlementBatch(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Merchant on a SettlementBatch
//----------------------------------------------------------------------------
func AssignMerchantToSettlementBatch( settlementBatchId uint64, merchantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

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
			// assign the Merchant	to the SettlementBatch
			//----------------------------------------------------------------------------
			parentObj.Merchant = &childObj

			//----------------------------------------------------------------------------
			// save the SettlementBatch
			//----------------------------------------------------------------------------
			return UpdateSettlementBatch(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Merchant", merchantId )
			return utils.RequestResult{false, msg, "assignMerchant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Merchant on a SettlementBatch
//----------------------------------------------------------------------------
func UnassignMerchantFromSettlementBatch(settlementBatchId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		//----------------------------------------------------------------------------
		// assign an empty Merchant to the Merchant
		//----------------------------------------------------------------------------
		parentObj.Merchant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Merchant
		//----------------------------------------------------------------------------
		parentObj.MerchantId = nil;

		//----------------------------------------------------------------------------
		// save the SettlementBatch
		//----------------------------------------------------------------------------
		return UpdateSettlementBatch(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more payoutsIds as a Payouts to a SettlementBatch
//----------------------------------------------------------------------------
func AddPayoutsToSettlementBatch ( settlementBatchId uint64, payoutsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		// slice the ids on comma with no spaces
		ids := strings.Split( payoutsIds, ",")

		for _, payoutsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Payout

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Payout
			// with a matching payoutsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payoutsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payouts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payouts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payouts", payoutsId )
				return utils.RequestResult{false, msg, "unassignPayouts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SettlementBatch from the gorm
		//----------------------------------------------------------------------------
		return GetSettlementBatch(settlementBatchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more payoutsIds as a Payouts from a SettlementBatch
//----------------------------------------------------------------------------
func RemovePayoutsFromSettlementBatch( settlementBatchId uint64, payoutsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		// slice the ids on comma with no spaces
		ids := strings.Split( payoutsIds, ",")

		for _, payoutsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Payout

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Payout
			// with a matching payoutsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payoutsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PayoutObj from the Payouts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payouts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payouts", payoutsId )
				return utils.RequestResult{false, msg, "removePayouts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SettlementBatch from the gorm
		//----------------------------------------------------------------------------
		return GetSettlementBatch(settlementBatchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a SettlementBatch
//----------------------------------------------------------------------------
func AddTransactionsToSettlementBatch ( settlementBatchId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "unassignTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SettlementBatch from the gorm
		//----------------------------------------------------------------------------
		return GetSettlementBatch(settlementBatchId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a SettlementBatch
//----------------------------------------------------------------------------
func RemoveTransactionsFromSettlementBatch( settlementBatchId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SettlementBatch with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSettlementBatch(settlementBatchId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SettlementBatch so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SettlementBatch)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "removeTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SettlementBatch from the gorm
		//----------------------------------------------------------------------------
		return GetSettlementBatch(settlementBatchId)

	} else {
		return parentRequestResult
	}
}

