package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CreditorDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCreditor - creates a new db entry
//----------------------------------------------------------------------------
func CreateCreditor(obj model.Creditor)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Creditor with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Creditor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCreditor", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCreditor - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCreditor(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Creditor

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Creditor with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Creditor using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Creditor using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCreditor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCreditor - returns all
//----------------------------------------------------------------------------
func GetAllCreditor()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Creditor

	//----------------------------------------------------------------------------
	// Request the ORM to find all Creditor
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Creditor" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Creditor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCreditor", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCreditor - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCreditor(obj model.Creditor)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Creditor using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Creditor using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCreditor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCreditor - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCreditor(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Creditor with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCreditor(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Creditor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Creditor)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Creditor using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Creditor using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCreditor", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more mandatesIds as a Mandates to a Creditor
//----------------------------------------------------------------------------
func AddMandatesToCreditor ( creditorId uint64, mandatesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Creditor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreditor(creditorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Creditor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Creditor)

		// slice the ids on comma with no spaces
		ids := strings.Split( mandatesIds, ",")

		for _, mandatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DirectDebitMandate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DirectDebitMandate
			// with a matching mandatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , mandatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Mandates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Mandates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Mandates", mandatesId )
				return utils.RequestResult{false, msg, "unassignMandates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Creditor from the gorm
		//----------------------------------------------------------------------------
		return GetCreditor(creditorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more mandatesIds as a Mandates from a Creditor
//----------------------------------------------------------------------------
func RemoveMandatesFromCreditor( creditorId uint64, mandatesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Creditor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreditor(creditorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Creditor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Creditor)

		// slice the ids on comma with no spaces
		ids := strings.Split( mandatesIds, ",")

		for _, mandatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DirectDebitMandate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DirectDebitMandate
			// with a matching mandatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , mandatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DirectDebitMandateObj from the Mandates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Mandates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Mandates", mandatesId )
				return utils.RequestResult{false, msg, "removeMandates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Creditor from the gorm
		//----------------------------------------------------------------------------
		return GetCreditor(creditorId)

	} else {
		return parentRequestResult
	}
}

