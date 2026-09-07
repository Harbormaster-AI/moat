package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PaymentProcessorDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePaymentProcessor - creates a new db entry
//----------------------------------------------------------------------------
func CreatePaymentProcessor(obj model.PaymentProcessor)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PaymentProcessor with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PaymentProcessor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePaymentProcessor", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPaymentProcessor - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPaymentProcessor(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PaymentProcessor

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PaymentProcessor with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PaymentProcessor using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PaymentProcessor using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPaymentProcessor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPaymentProcessor - returns all
//----------------------------------------------------------------------------
func GetAllPaymentProcessor()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PaymentProcessor

	//----------------------------------------------------------------------------
	// Request the ORM to find all PaymentProcessor
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PaymentProcessor" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PaymentProcessor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPaymentProcessor", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePaymentProcessor - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePaymentProcessor(obj model.PaymentProcessor)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PaymentProcessor using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PaymentProcessor using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePaymentProcessor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePaymentProcessor - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePaymentProcessor(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPaymentProcessor(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PaymentProcessor)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PaymentProcessor using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PaymentProcessor using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePaymentProcessor", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more institutionsIds as a Institutions to a PaymentProcessor
//----------------------------------------------------------------------------
func AddInstitutionsToPaymentProcessor ( paymentProcessorId uint64, institutionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

		// slice the ids on comma with no spaces
		ids := strings.Split( institutionsIds, ",")

		for _, institutionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FinancialInstitution

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FinancialInstitution
			// with a matching institutionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , institutionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Institutions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Institutions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Institutions", institutionsId )
				return utils.RequestResult{false, msg, "unassignInstitutions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more institutionsIds as a Institutions from a PaymentProcessor
//----------------------------------------------------------------------------
func RemoveInstitutionsFromPaymentProcessor( paymentProcessorId uint64, institutionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

		// slice the ids on comma with no spaces
		ids := strings.Split( institutionsIds, ",")

		for _, institutionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FinancialInstitution

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FinancialInstitution
			// with a matching institutionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , institutionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FinancialInstitutionObj from the Institutions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Institutions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Institutions", institutionsId )
				return utils.RequestResult{false, msg, "removeInstitutions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a PaymentProcessor
//----------------------------------------------------------------------------
func AddContractsToPaymentProcessor ( paymentProcessorId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentContract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "unassignContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a PaymentProcessor
//----------------------------------------------------------------------------
func RemoveContractsFromPaymentProcessor( paymentProcessorId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentContract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more settlementsIds as a Settlements to a PaymentProcessor
//----------------------------------------------------------------------------
func AddSettlementsToPaymentProcessor ( paymentProcessorId uint64, settlementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

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
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more settlementsIds as a Settlements from a PaymentProcessor
//----------------------------------------------------------------------------
func RemoveSettlementsFromPaymentProcessor( paymentProcessorId uint64, settlementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PaymentProcessor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentProcessor(paymentProcessorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentProcessor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentProcessor)

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
		// retrieve the modified PaymentProcessor from the gorm
		//----------------------------------------------------------------------------
		return GetPaymentProcessor(paymentProcessorId)

	} else {
		return parentRequestResult
	}
}

