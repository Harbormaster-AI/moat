package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EquityGrantDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEquityGrant - creates a new db entry
//----------------------------------------------------------------------------
func CreateEquityGrant(obj model.EquityGrant)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EquityGrant with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EquityGrant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEquityGrant", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEquityGrant - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEquityGrant(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EquityGrant

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EquityGrant with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EquityGrant using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EquityGrant using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEquityGrant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEquityGrant - returns all
//----------------------------------------------------------------------------
func GetAllEquityGrant()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EquityGrant

	//----------------------------------------------------------------------------
	// Request the ORM to find all EquityGrant
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EquityGrant" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EquityGrant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEquityGrant", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEquityGrant - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEquityGrant(obj model.EquityGrant)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EquityGrant using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EquityGrant using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEquityGrant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEquityGrant - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEquityGrant(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EquityGrant with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEquityGrant(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EquityGrant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EquityGrant)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EquityGrant using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EquityGrant using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEquityGrant", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CompensationPackage on a EquityGrant
//----------------------------------------------------------------------------
func AssignCompensationPackageToEquityGrant( equityGrantId uint64, compensationPackageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EquityGrant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEquityGrant(equityGrantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EquityGrant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EquityGrant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CompensationPackage

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CompensationPackage with a
		// matching compensationPackageId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, compensationPackageId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CompensationPackage	to the EquityGrant
			//----------------------------------------------------------------------------
			parentObj.CompensationPackage = &childObj

			//----------------------------------------------------------------------------
			// save the EquityGrant
			//----------------------------------------------------------------------------
			return UpdateEquityGrant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompensationPackage", compensationPackageId )
			return utils.RequestResult{false, msg, "assignCompensationPackage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CompensationPackage on a EquityGrant
//----------------------------------------------------------------------------
func UnassignCompensationPackageFromEquityGrant(equityGrantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EquityGrant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEquityGrant(equityGrantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EquityGrant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EquityGrant)

		//----------------------------------------------------------------------------
		// assign an empty CompensationPackage to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CompensationPackage
		//----------------------------------------------------------------------------
		parentObj.CompensationPackageId = nil;

		//----------------------------------------------------------------------------
		// save the EquityGrant
		//----------------------------------------------------------------------------
		return UpdateEquityGrant(parentObj)

	} else {
		return parentRequestResult
	}

}


