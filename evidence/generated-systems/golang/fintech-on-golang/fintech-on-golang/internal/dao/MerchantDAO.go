package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MerchantDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMerchant - creates a new db entry
//----------------------------------------------------------------------------
func CreateMerchant(obj model.Merchant)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Merchant with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Merchant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMerchant", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMerchant - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMerchant(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Merchant

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Merchant with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Merchant using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Merchant using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMerchant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMerchant - returns all
//----------------------------------------------------------------------------
func GetAllMerchant()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Merchant

	//----------------------------------------------------------------------------
	// Request the ORM to find all Merchant
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Merchant" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Merchant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMerchant", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMerchant - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMerchant(obj model.Merchant)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Merchant using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Merchant using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMerchant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMerchant - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMerchant(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMerchant(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Merchant)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Merchant using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Merchant using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMerchant", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more terminalsIds as a Terminals to a Merchant
//----------------------------------------------------------------------------
func AddTerminalsToMerchant ( merchantId uint64, terminalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( terminalsIds, ",")

		for _, terminalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Terminal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Terminal
			// with a matching terminalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , terminalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Terminals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Terminals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Terminals", terminalsId )
				return utils.RequestResult{false, msg, "unassignTerminals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more terminalsIds as a Terminals from a Merchant
//----------------------------------------------------------------------------
func RemoveTerminalsFromMerchant( merchantId uint64, terminalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( terminalsIds, ",")

		for _, terminalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Terminal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Terminal
			// with a matching terminalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , terminalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TerminalObj from the Terminals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Terminals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Terminals", terminalsId )
				return utils.RequestResult{false, msg, "removeTerminals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentContractsIds as a PaymentContracts to a Merchant
//----------------------------------------------------------------------------
func AddPaymentContractsToMerchant ( merchantId uint64, paymentContractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentContractsIds, ",")

		for _, paymentContractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentContract
			// with a matching paymentContractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentContractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentContracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentContracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentContracts", paymentContractsId )
				return utils.RequestResult{false, msg, "unassignPaymentContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentContractsIds as a PaymentContracts from a Merchant
//----------------------------------------------------------------------------
func RemovePaymentContractsFromMerchant( merchantId uint64, paymentContractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentContractsIds, ",")

		for _, paymentContractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentContract
			// with a matching paymentContractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentContractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentContractObj from the PaymentContracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentContracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentContracts", paymentContractsId )
				return utils.RequestResult{false, msg, "removePaymentContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more payoutsIds as a Payouts to a Merchant
//----------------------------------------------------------------------------
func AddPayoutsToMerchant ( merchantId uint64, payoutsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

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
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more payoutsIds as a Payouts from a Merchant
//----------------------------------------------------------------------------
func RemovePayoutsFromMerchant( merchantId uint64, payoutsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

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
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more settlementsIds as a Settlements to a Merchant
//----------------------------------------------------------------------------
func AddSettlementsToMerchant ( merchantId uint64, settlementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( settlementsIds, ",")

		for _, settlementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SettlementBatch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SettlementBatch
			// with a matching settlementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , settlementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Settlements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Settlements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Settlements", settlementsId )
				return utils.RequestResult{false, msg, "unassignSettlements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more settlementsIds as a Settlements from a Merchant
//----------------------------------------------------------------------------
func RemoveSettlementsFromMerchant( merchantId uint64, settlementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( settlementsIds, ",")

		for _, settlementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SettlementBatch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SettlementBatch
			// with a matching settlementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , settlementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SettlementBatchObj from the Settlements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Settlements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Settlements", settlementsId )
				return utils.RequestResult{false, msg, "removeSettlements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more disputesIds as a Disputes to a Merchant
//----------------------------------------------------------------------------
func AddDisputesToMerchant ( merchantId uint64, disputesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Disputes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "unassignDisputes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more disputesIds as a Disputes from a Merchant
//----------------------------------------------------------------------------
func RemoveDisputesFromMerchant( merchantId uint64, disputesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DisputeObj from the Disputes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "removeDisputes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more invoicesIds as a Invoices to a Merchant
//----------------------------------------------------------------------------
func AddInvoicesToMerchant ( merchantId uint64, invoicesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( invoicesIds, ",")

		for _, invoicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Invoice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Invoice
			// with a matching invoicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , invoicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Invoices using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Invoices").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Invoices", invoicesId )
				return utils.RequestResult{false, msg, "unassignInvoices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more invoicesIds as a Invoices from a Merchant
//----------------------------------------------------------------------------
func RemoveInvoicesFromMerchant( merchantId uint64, invoicesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Merchant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMerchant(merchantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Merchant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Merchant)

		// slice the ids on comma with no spaces
		ids := strings.Split( invoicesIds, ",")

		for _, invoicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Invoice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Invoice
			// with a matching invoicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , invoicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InvoiceObj from the Invoices array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Invoices").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Invoices", invoicesId )
				return utils.RequestResult{false, msg, "removeInvoices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Merchant from the gorm
		//----------------------------------------------------------------------------
		return GetMerchant(merchantId)

	} else {
		return parentRequestResult
	}
}

