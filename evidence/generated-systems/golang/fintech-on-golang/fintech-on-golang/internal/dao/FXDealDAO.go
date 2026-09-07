package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FXDealDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFXDeal - creates a new db entry
//----------------------------------------------------------------------------
func CreateFXDeal(obj model.FXDeal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FXDeal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FXDeal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFXDeal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFXDeal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFXDeal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FXDeal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FXDeal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FXDeal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FXDeal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFXDeal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFXDeal - returns all
//----------------------------------------------------------------------------
func GetAllFXDeal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FXDeal

	//----------------------------------------------------------------------------
	// Request the ORM to find all FXDeal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FXDeal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FXDeal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFXDeal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFXDeal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFXDeal(obj model.FXDeal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FXDeal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FXDeal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFXDeal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFXDeal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFXDeal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FXDeal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFXDeal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXDeal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FXDeal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FXDeal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FXDeal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFXDeal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Quote on a FXDeal
//----------------------------------------------------------------------------
func AssignQuoteToFXDeal( fXDealId uint64, quoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXDeal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXDeal(fXDealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXDeal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXDeal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FXQuote

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FXQuote with a
		// matching quoteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, quoteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Quote	to the FXDeal
			//----------------------------------------------------------------------------
			parentObj.Quote = &childObj

			//----------------------------------------------------------------------------
			// save the FXDeal
			//----------------------------------------------------------------------------
			return UpdateFXDeal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quote", quoteId )
			return utils.RequestResult{false, msg, "assignQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Quote on a FXDeal
//----------------------------------------------------------------------------
func UnassignQuoteFromFXDeal(fXDealId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXDeal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXDeal(fXDealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXDeal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXDeal)

		//----------------------------------------------------------------------------
		// assign an empty FXQuote to the Quote
		//----------------------------------------------------------------------------
		parentObj.Quote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Quote
		//----------------------------------------------------------------------------
		parentObj.QuoteId = nil;

		//----------------------------------------------------------------------------
		// save the FXDeal
		//----------------------------------------------------------------------------
		return UpdateFXDeal(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more paymentOrdersIds as a PaymentOrders to a FXDeal
//----------------------------------------------------------------------------
func AddPaymentOrdersToFXDeal ( fXDealId uint64, paymentOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXDeal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXDeal(fXDealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXDeal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXDeal)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentOrdersIds, ",")

		for _, paymentOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentOrder
			// with a matching paymentOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentOrders", paymentOrdersId )
				return utils.RequestResult{false, msg, "unassignPaymentOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FXDeal from the gorm
		//----------------------------------------------------------------------------
		return GetFXDeal(fXDealId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentOrdersIds as a PaymentOrders from a FXDeal
//----------------------------------------------------------------------------
func RemovePaymentOrdersFromFXDeal( fXDealId uint64, paymentOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FXDeal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXDeal(fXDealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXDeal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXDeal)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentOrdersIds, ",")

		for _, paymentOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentOrder
			// with a matching paymentOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentOrderObj from the PaymentOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentOrders", paymentOrdersId )
				return utils.RequestResult{false, msg, "removePaymentOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FXDeal from the gorm
		//----------------------------------------------------------------------------
		return GetFXDeal(fXDealId)

	} else {
		return parentRequestResult
	}
}

