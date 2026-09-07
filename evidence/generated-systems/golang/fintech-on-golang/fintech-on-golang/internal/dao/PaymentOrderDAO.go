package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PaymentOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePaymentOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreatePaymentOrder(obj model.PaymentOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PaymentOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PaymentOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePaymentOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPaymentOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPaymentOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PaymentOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PaymentOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PaymentOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PaymentOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPaymentOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPaymentOrder - returns all
//----------------------------------------------------------------------------
func GetAllPaymentOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PaymentOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all PaymentOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PaymentOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PaymentOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPaymentOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePaymentOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePaymentOrder(obj model.PaymentOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PaymentOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PaymentOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePaymentOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePaymentOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePaymentOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPaymentOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PaymentOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PaymentOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePaymentOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SourceAccount on a PaymentOrder
//----------------------------------------------------------------------------
func AssignSourceAccountToPaymentOrder( paymentOrderId uint64, sourceAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching sourceAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, sourceAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SourceAccount	to the PaymentOrder
			//----------------------------------------------------------------------------
			parentObj.SourceAccount = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentOrder
			//----------------------------------------------------------------------------
			return UpdatePaymentOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceAccount", sourceAccountId )
			return utils.RequestResult{false, msg, "assignSourceAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceAccount on a PaymentOrder
//----------------------------------------------------------------------------
func UnassignSourceAccountFromPaymentOrder(paymentOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// assign an empty Account to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceAccount
		//----------------------------------------------------------------------------
		parentObj.SourceAccountId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentOrder
		//----------------------------------------------------------------------------
		return UpdatePaymentOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationAccount on a PaymentOrder
//----------------------------------------------------------------------------
func AssignDestinationAccountToPaymentOrder( paymentOrderId uint64, destinationAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

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
			// assign the DestinationAccount	to the PaymentOrder
			//----------------------------------------------------------------------------
			parentObj.DestinationAccount = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentOrder
			//----------------------------------------------------------------------------
			return UpdatePaymentOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationAccount", destinationAccountId )
			return utils.RequestResult{false, msg, "assignDestinationAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationAccount on a PaymentOrder
//----------------------------------------------------------------------------
func UnassignDestinationAccountFromPaymentOrder(paymentOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// assign an empty Account to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationAccount
		//----------------------------------------------------------------------------
		parentObj.DestinationAccountId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentOrder
		//----------------------------------------------------------------------------
		return UpdatePaymentOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Beneficiary on a PaymentOrder
//----------------------------------------------------------------------------
func AssignBeneficiaryToPaymentOrder( paymentOrderId uint64, beneficiaryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Beneficiary

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Beneficiary with a
		// matching beneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, beneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Beneficiary	to the PaymentOrder
			//----------------------------------------------------------------------------
			parentObj.Beneficiary = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentOrder
			//----------------------------------------------------------------------------
			return UpdatePaymentOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiary", beneficiaryId )
			return utils.RequestResult{false, msg, "assignBeneficiary", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Beneficiary on a PaymentOrder
//----------------------------------------------------------------------------
func UnassignBeneficiaryFromPaymentOrder(paymentOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// assign an empty Beneficiary to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.Beneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.BeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentOrder
		//----------------------------------------------------------------------------
		return UpdatePaymentOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FxDeal on a PaymentOrder
//----------------------------------------------------------------------------
func AssignFxDealToPaymentOrder( paymentOrderId uint64, fxDealId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FXDeal

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FXDeal with a
		// matching fxDealId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, fxDealId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FxDeal	to the PaymentOrder
			//----------------------------------------------------------------------------
			parentObj.FxDeal = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentOrder
			//----------------------------------------------------------------------------
			return UpdatePaymentOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FxDeal", fxDealId )
			return utils.RequestResult{false, msg, "assignFxDeal", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FxDeal on a PaymentOrder
//----------------------------------------------------------------------------
func UnassignFxDealFromPaymentOrder(paymentOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		//----------------------------------------------------------------------------
		// assign an empty FXDeal to the FxDeal
		//----------------------------------------------------------------------------
		parentObj.FxDeal = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FxDeal
		//----------------------------------------------------------------------------
		parentObj.FxDealId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentOrder
		//----------------------------------------------------------------------------
		return UpdatePaymentOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a PaymentOrder
//----------------------------------------------------------------------------
func AddTransactionsToPaymentOrder ( paymentOrderId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

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
		// retrieve the modified PaymentOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentOrder(paymentOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a PaymentOrder
//----------------------------------------------------------------------------
func RemoveTransactionsFromPaymentOrder( paymentOrderId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

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
		// retrieve the modified PaymentOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentOrder(paymentOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more feesIds as a Fees to a PaymentOrder
//----------------------------------------------------------------------------
func AddFeesToPaymentOrder ( paymentOrderId uint64, feesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( feesIds, ",")

		for _, feesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AppliedFee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AppliedFee
			// with a matching feesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Fees using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Fees").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Fees", feesId )
				return utils.RequestResult{false, msg, "unassignFees", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentOrder(paymentOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more feesIds as a Fees from a PaymentOrder
//----------------------------------------------------------------------------
func RemoveFeesFromPaymentOrder( paymentOrderId uint64, feesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentOrder(paymentOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( feesIds, ",")

		for _, feesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AppliedFee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AppliedFee
			// with a matching feesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AppliedFeeObj from the Fees array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Fees").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Fees", feesId )
				return utils.RequestResult{false, msg, "removeFees", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentOrder(paymentOrderId)

	} else {
		return parentRequestResult
	}
}

