package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TransactionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTransaction - creates a new db entry
//----------------------------------------------------------------------------
func CreateTransaction(obj model.Transaction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Transaction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Transaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTransaction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTransaction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTransaction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Transaction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Transaction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Transaction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Transaction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTransaction - returns all
//----------------------------------------------------------------------------
func GetAllTransaction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Transaction

	//----------------------------------------------------------------------------
	// Request the ORM to find all Transaction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Transaction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Transaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTransaction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTransaction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTransaction(obj model.Transaction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Transaction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Transaction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTransaction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTransaction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTransaction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Transaction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Transaction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTransaction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a Transaction
//----------------------------------------------------------------------------
func AssignAccountToTransaction( transactionId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Transaction
//----------------------------------------------------------------------------
func UnassignAccountFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Wallet on a Transaction
//----------------------------------------------------------------------------
func AssignWalletToTransaction( transactionId uint64, walletId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Wallet

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Wallet with a
		// matching walletId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, walletId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Wallet	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Wallet = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Wallet", walletId )
			return utils.RequestResult{false, msg, "assignWallet", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Wallet on a Transaction
//----------------------------------------------------------------------------
func UnassignWalletFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Wallet to the Wallet
		//----------------------------------------------------------------------------
		parentObj.Wallet = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Wallet
		//----------------------------------------------------------------------------
		parentObj.WalletId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PaymentOrder on a Transaction
//----------------------------------------------------------------------------
func AssignPaymentOrderToTransaction( transactionId uint64, paymentOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PaymentOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentOrder with a
		// matching paymentOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, paymentOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PaymentOrder	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.PaymentOrder = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentOrder", paymentOrderId )
			return utils.RequestResult{false, msg, "assignPaymentOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PaymentOrder on a Transaction
//----------------------------------------------------------------------------
func UnassignPaymentOrderFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty PaymentOrder to the PaymentOrder
		//----------------------------------------------------------------------------
		parentObj.PaymentOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentOrder
		//----------------------------------------------------------------------------
		parentObj.PaymentOrderId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Merchant on a Transaction
//----------------------------------------------------------------------------
func AssignMerchantToTransaction( transactionId uint64, merchantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the Merchant	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Merchant = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Merchant", merchantId )
			return utils.RequestResult{false, msg, "assignMerchant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Merchant on a Transaction
//----------------------------------------------------------------------------
func UnassignMerchantFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty Merchant to the Merchant
		//----------------------------------------------------------------------------
		parentObj.Merchant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Merchant
		//----------------------------------------------------------------------------
		parentObj.MerchantId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Card on a Transaction
//----------------------------------------------------------------------------
func AssignCardToTransaction( transactionId uint64, cardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

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
			// assign the Card	to the Transaction
			//----------------------------------------------------------------------------
			parentObj.Card = &childObj

			//----------------------------------------------------------------------------
			// save the Transaction
			//----------------------------------------------------------------------------
			return UpdateTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Card", cardId )
			return utils.RequestResult{false, msg, "assignCard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Card on a Transaction
//----------------------------------------------------------------------------
func UnassignCardFromTransaction(transactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the Card
		//----------------------------------------------------------------------------
		parentObj.Card = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Card
		//----------------------------------------------------------------------------
		parentObj.CardId = nil;

		//----------------------------------------------------------------------------
		// save the Transaction
		//----------------------------------------------------------------------------
		return UpdateTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more relatedTransactionsIds as a RelatedTransactions to a Transaction
//----------------------------------------------------------------------------
func AddRelatedTransactionsToTransaction ( transactionId uint64, relatedTransactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedTransactionsIds, ",")

		for _, relatedTransactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching relatedTransactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedTransactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedTransactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedTransactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedTransactions", relatedTransactionsId )
				return utils.RequestResult{false, msg, "unassignRelatedTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Transaction from the gorm
		//----------------------------------------------------------------------------
		return GetTransaction(transactionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedTransactionsIds as a RelatedTransactions from a Transaction
//----------------------------------------------------------------------------
func RemoveRelatedTransactionsFromTransaction( transactionId uint64, relatedTransactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedTransactionsIds, ",")

		for _, relatedTransactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching relatedTransactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedTransactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the RelatedTransactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedTransactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedTransactions", relatedTransactionsId )
				return utils.RequestResult{false, msg, "removeRelatedTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Transaction from the gorm
		//----------------------------------------------------------------------------
		return GetTransaction(transactionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a Transaction
//----------------------------------------------------------------------------
func AddAlertsToTransaction ( transactionId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceAlert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceAlert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Alerts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "unassignAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Transaction from the gorm
		//----------------------------------------------------------------------------
		return GetTransaction(transactionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a Transaction
//----------------------------------------------------------------------------
func RemoveAlertsFromTransaction( transactionId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Transaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransaction(transactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Transaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Transaction)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceAlert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceAlert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ComplianceAlertObj from the Alerts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "removeAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Transaction from the gorm
		//----------------------------------------------------------------------------
		return GetTransaction(transactionId)

	} else {
		return parentRequestResult
	}
}

