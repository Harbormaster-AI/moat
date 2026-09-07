package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExperimentVariantDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExperimentVariant - creates a new db entry
//----------------------------------------------------------------------------
func CreateExperimentVariant(obj model.ExperimentVariant)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ExperimentVariant with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ExperimentVariant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExperimentVariant", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExperimentVariant - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExperimentVariant(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ExperimentVariant

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ExperimentVariant with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ExperimentVariant using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ExperimentVariant using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExperimentVariant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExperimentVariant - returns all
//----------------------------------------------------------------------------
func GetAllExperimentVariant()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ExperimentVariant

	//----------------------------------------------------------------------------
	// Request the ORM to find all ExperimentVariant
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ExperimentVariant" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExperimentVariant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExperimentVariant", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExperimentVariant - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExperimentVariant(obj model.ExperimentVariant)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ExperimentVariant using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ExperimentVariant using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExperimentVariant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExperimentVariant - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExperimentVariant(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExperimentVariant(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ExperimentVariant using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ExperimentVariant using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExperimentVariant", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Experiment on a ExperimentVariant
//----------------------------------------------------------------------------
func AssignExperimentToExperimentVariant( experimentVariantId uint64, experimentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Experiment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Experiment with a
		// matching experimentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, experimentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Experiment	to the ExperimentVariant
			//----------------------------------------------------------------------------
			parentObj.Experiment = &childObj

			//----------------------------------------------------------------------------
			// save the ExperimentVariant
			//----------------------------------------------------------------------------
			return UpdateExperimentVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Experiment", experimentId )
			return utils.RequestResult{false, msg, "assignExperiment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Experiment on a ExperimentVariant
//----------------------------------------------------------------------------
func UnassignExperimentFromExperimentVariant(experimentVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// assign an empty Experiment to the Experiment
		//----------------------------------------------------------------------------
		parentObj.Experiment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Experiment
		//----------------------------------------------------------------------------
		parentObj.ExperimentId = nil;

		//----------------------------------------------------------------------------
		// save the ExperimentVariant
		//----------------------------------------------------------------------------
		return UpdateExperimentVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CreativeVariation on a ExperimentVariant
//----------------------------------------------------------------------------
func AssignCreativeVariationToExperimentVariant( experimentVariantId uint64, creativeVariationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CreativeVariation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CreativeVariation with a
		// matching creativeVariationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, creativeVariationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CreativeVariation	to the ExperimentVariant
			//----------------------------------------------------------------------------
			parentObj.CreativeVariation = &childObj

			//----------------------------------------------------------------------------
			// save the ExperimentVariant
			//----------------------------------------------------------------------------
			return UpdateExperimentVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeVariation", creativeVariationId )
			return utils.RequestResult{false, msg, "assignCreativeVariation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CreativeVariation on a ExperimentVariant
//----------------------------------------------------------------------------
func UnassignCreativeVariationFromExperimentVariant(experimentVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// assign an empty CreativeVariation to the CreativeVariation
		//----------------------------------------------------------------------------
		parentObj.CreativeVariation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CreativeVariation
		//----------------------------------------------------------------------------
		parentObj.CreativeVariationId = nil;

		//----------------------------------------------------------------------------
		// save the ExperimentVariant
		//----------------------------------------------------------------------------
		return UpdateExperimentVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineItem on a ExperimentVariant
//----------------------------------------------------------------------------
func AssignLineItemToExperimentVariant( experimentVariantId uint64, lineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LineItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LineItem with a
		// matching lineItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lineItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LineItem	to the ExperimentVariant
			//----------------------------------------------------------------------------
			parentObj.LineItem = &childObj

			//----------------------------------------------------------------------------
			// save the ExperimentVariant
			//----------------------------------------------------------------------------
			return UpdateExperimentVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItem", lineItemId )
			return utils.RequestResult{false, msg, "assignLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineItem on a ExperimentVariant
//----------------------------------------------------------------------------
func UnassignLineItemFromExperimentVariant(experimentVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExperimentVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperimentVariant(experimentVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExperimentVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExperimentVariant)

		//----------------------------------------------------------------------------
		// assign an empty LineItem to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItemId = nil;

		//----------------------------------------------------------------------------
		// save the ExperimentVariant
		//----------------------------------------------------------------------------
		return UpdateExperimentVariant(parentObj)

	} else {
		return parentRequestResult
	}

}


