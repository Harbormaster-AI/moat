package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CardTokenizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCardTokenization - creates a new db entry
//----------------------------------------------------------------------------
func CreateCardTokenization(obj model.CardTokenization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CardTokenization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CardTokenization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCardTokenization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCardTokenization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCardTokenization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CardTokenization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CardTokenization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CardTokenization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CardTokenization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCardTokenization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCardTokenization - returns all
//----------------------------------------------------------------------------
func GetAllCardTokenization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CardTokenization

	//----------------------------------------------------------------------------
	// Request the ORM to find all CardTokenization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CardTokenization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CardTokenization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCardTokenization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCardTokenization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCardTokenization(obj model.CardTokenization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CardTokenization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CardTokenization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCardTokenization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCardTokenization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCardTokenization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CardTokenization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCardTokenization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CardTokenization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CardTokenization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CardTokenization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CardTokenization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCardTokenization", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Card on a CardTokenization
//----------------------------------------------------------------------------
func AssignCardToCardTokenization( cardTokenizationId uint64, cardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CardTokenization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCardTokenization(cardTokenizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CardTokenization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CardTokenization)

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
			// assign the Card	to the CardTokenization
			//----------------------------------------------------------------------------
			parentObj.Card = &childObj

			//----------------------------------------------------------------------------
			// save the CardTokenization
			//----------------------------------------------------------------------------
			return UpdateCardTokenization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Card", cardId )
			return utils.RequestResult{false, msg, "assignCard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Card on a CardTokenization
//----------------------------------------------------------------------------
func UnassignCardFromCardTokenization(cardTokenizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CardTokenization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCardTokenization(cardTokenizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CardTokenization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CardTokenization)

		//----------------------------------------------------------------------------
		// assign an empty PaymentCard to the Card
		//----------------------------------------------------------------------------
		parentObj.Card = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Card
		//----------------------------------------------------------------------------
		parentObj.CardId = nil;

		//----------------------------------------------------------------------------
		// save the CardTokenization
		//----------------------------------------------------------------------------
		return UpdateCardTokenization(parentObj)

	} else {
		return parentRequestResult
	}

}


