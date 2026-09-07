package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftOrder(obj model.AircraftOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftOrder - returns all
//----------------------------------------------------------------------------
func GetAllAircraftOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftOrder(obj model.AircraftOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Operator on a AircraftOrder
//----------------------------------------------------------------------------
func AssignOperatorToAircraftOrder( aircraftOrderId uint64, operatorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Operator

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Operator with a
		// matching operatorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, operatorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Operator	to the AircraftOrder
			//----------------------------------------------------------------------------
			parentObj.Operator = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftOrder
			//----------------------------------------------------------------------------
			return UpdateAircraftOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operator", operatorId )
			return utils.RequestResult{false, msg, "assignOperator", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Operator on a AircraftOrder
//----------------------------------------------------------------------------
func UnassignOperatorFromAircraftOrder(aircraftOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// assign an empty Operator to the Operator
		//----------------------------------------------------------------------------
		parentObj.Operator = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Operator
		//----------------------------------------------------------------------------
		parentObj.OperatorId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftOrder
		//----------------------------------------------------------------------------
		return UpdateAircraftOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Variant on a AircraftOrder
//----------------------------------------------------------------------------
func AssignVariantToAircraftOrder( aircraftOrderId uint64, variantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftVariant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftVariant with a
		// matching variantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, variantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Variant	to the AircraftOrder
			//----------------------------------------------------------------------------
			parentObj.Variant = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftOrder
			//----------------------------------------------------------------------------
			return UpdateAircraftOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variant", variantId )
			return utils.RequestResult{false, msg, "assignVariant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Variant on a AircraftOrder
//----------------------------------------------------------------------------
func UnassignVariantFromAircraftOrder(aircraftOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// assign an empty AircraftVariant to the Variant
		//----------------------------------------------------------------------------
		parentObj.Variant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Variant
		//----------------------------------------------------------------------------
		parentObj.VariantId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftOrder
		//----------------------------------------------------------------------------
		return UpdateAircraftOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Quote on a AircraftOrder
//----------------------------------------------------------------------------
func AssignQuoteToAircraftOrder( aircraftOrderId uint64, quoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Quote

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Quote with a
		// matching quoteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, quoteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Quote	to the AircraftOrder
			//----------------------------------------------------------------------------
			parentObj.Quote = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftOrder
			//----------------------------------------------------------------------------
			return UpdateAircraftOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quote", quoteId )
			return utils.RequestResult{false, msg, "assignQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Quote on a AircraftOrder
//----------------------------------------------------------------------------
func UnassignQuoteFromAircraftOrder(aircraftOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// assign an empty Quote to the Quote
		//----------------------------------------------------------------------------
		parentObj.Quote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Quote
		//----------------------------------------------------------------------------
		parentObj.QuoteId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftOrder
		//----------------------------------------------------------------------------
		return UpdateAircraftOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PurchaseAgreement on a AircraftOrder
//----------------------------------------------------------------------------
func AssignPurchaseAgreementToAircraftOrder( aircraftOrderId uint64, purchaseAgreementId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PurchaseAgreement

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PurchaseAgreement with a
		// matching purchaseAgreementId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, purchaseAgreementId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PurchaseAgreement	to the AircraftOrder
			//----------------------------------------------------------------------------
			parentObj.PurchaseAgreement = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftOrder
			//----------------------------------------------------------------------------
			return UpdateAircraftOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseAgreement", purchaseAgreementId )
			return utils.RequestResult{false, msg, "assignPurchaseAgreement", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PurchaseAgreement on a AircraftOrder
//----------------------------------------------------------------------------
func UnassignPurchaseAgreementFromAircraftOrder(aircraftOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOrder(aircraftOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOrder)

		//----------------------------------------------------------------------------
		// assign an empty PurchaseAgreement to the PurchaseAgreement
		//----------------------------------------------------------------------------
		parentObj.PurchaseAgreement = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PurchaseAgreement
		//----------------------------------------------------------------------------
		parentObj.PurchaseAgreementId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftOrder
		//----------------------------------------------------------------------------
		return UpdateAircraftOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


