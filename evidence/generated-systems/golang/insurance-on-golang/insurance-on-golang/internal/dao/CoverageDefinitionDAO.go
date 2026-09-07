package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CoverageDefinitionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCoverageDefinition - creates a new db entry
//----------------------------------------------------------------------------
func CreateCoverageDefinition(obj model.CoverageDefinition)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CoverageDefinition with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CoverageDefinition", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCoverageDefinition", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCoverageDefinition - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCoverageDefinition(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CoverageDefinition

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CoverageDefinition with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CoverageDefinition using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CoverageDefinition using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCoverageDefinition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCoverageDefinition - returns all
//----------------------------------------------------------------------------
func GetAllCoverageDefinition()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CoverageDefinition

	//----------------------------------------------------------------------------
	// Request the ORM to find all CoverageDefinition
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CoverageDefinition" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CoverageDefinition", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCoverageDefinition", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCoverageDefinition - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCoverageDefinition(obj model.CoverageDefinition)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CoverageDefinition using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CoverageDefinition using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCoverageDefinition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCoverageDefinition - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCoverageDefinition(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CoverageDefinition with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCoverageDefinition(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CoverageDefinition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CoverageDefinition)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CoverageDefinition using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CoverageDefinition using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCoverageDefinition", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Product on a CoverageDefinition
//----------------------------------------------------------------------------
func AssignProductToCoverageDefinition( coverageDefinitionId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CoverageDefinition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverageDefinition(coverageDefinitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CoverageDefinition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CoverageDefinition)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsuranceProduct

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsuranceProduct with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the CoverageDefinition
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the CoverageDefinition
			//----------------------------------------------------------------------------
			return UpdateCoverageDefinition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a CoverageDefinition
//----------------------------------------------------------------------------
func UnassignProductFromCoverageDefinition(coverageDefinitionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CoverageDefinition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverageDefinition(coverageDefinitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CoverageDefinition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CoverageDefinition)

		//----------------------------------------------------------------------------
		// assign an empty InsuranceProduct to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the CoverageDefinition
		//----------------------------------------------------------------------------
		return UpdateCoverageDefinition(parentObj)

	} else {
		return parentRequestResult
	}

}


