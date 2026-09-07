package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QualitySpecificationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQualitySpecification - creates a new db entry
//----------------------------------------------------------------------------
func CreateQualitySpecification(obj model.QualitySpecification)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a QualitySpecification with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a QualitySpecification", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQualitySpecification", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQualitySpecification - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQualitySpecification(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.QualitySpecification

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a QualitySpecification with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a QualitySpecification using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a QualitySpecification using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQualitySpecification", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQualitySpecification - returns all
//----------------------------------------------------------------------------
func GetAllQualitySpecification()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.QualitySpecification

	//----------------------------------------------------------------------------
	// Request the ORM to find all QualitySpecification
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all QualitySpecification" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all QualitySpecification", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQualitySpecification", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQualitySpecification - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQualitySpecification(obj model.QualitySpecification)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a QualitySpecification using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a QualitySpecification using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQualitySpecification", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQualitySpecification - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQualitySpecification(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the QualitySpecification with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQualitySpecification(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualitySpecification so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.QualitySpecification)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a QualitySpecification using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a QualitySpecification using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQualitySpecification", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a QualitySpecification
//----------------------------------------------------------------------------
func AssignItemToQualitySpecification( qualitySpecificationId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QualitySpecification with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualitySpecification(qualitySpecificationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualitySpecification so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualitySpecification)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the QualitySpecification
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the QualitySpecification
			//----------------------------------------------------------------------------
			return UpdateQualitySpecification(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a QualitySpecification
//----------------------------------------------------------------------------
func UnassignItemFromQualitySpecification(qualitySpecificationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QualitySpecification with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualitySpecification(qualitySpecificationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualitySpecification so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualitySpecification)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the QualitySpecification
		//----------------------------------------------------------------------------
		return UpdateQualitySpecification(parentObj)

	} else {
		return parentRequestResult
	}

}


