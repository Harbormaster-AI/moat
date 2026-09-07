package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExperimentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExperiment - creates a new db entry
//----------------------------------------------------------------------------
func CreateExperiment(obj model.Experiment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Experiment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Experiment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExperiment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExperiment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExperiment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Experiment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Experiment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Experiment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Experiment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExperiment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExperiment - returns all
//----------------------------------------------------------------------------
func GetAllExperiment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Experiment

	//----------------------------------------------------------------------------
	// Request the ORM to find all Experiment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Experiment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Experiment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExperiment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExperiment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExperiment(obj model.Experiment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Experiment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Experiment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExperiment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExperiment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExperiment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Experiment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExperiment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Experiment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Experiment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Experiment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Experiment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExperiment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a Experiment
//----------------------------------------------------------------------------
func AssignCampaignToExperiment( experimentId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Experiment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperiment(experimentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Experiment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Experiment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching campaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, campaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Campaign	to the Experiment
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the Experiment
			//----------------------------------------------------------------------------
			return UpdateExperiment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a Experiment
//----------------------------------------------------------------------------
func UnassignCampaignFromExperiment(experimentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Experiment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperiment(experimentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Experiment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Experiment)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the Experiment
		//----------------------------------------------------------------------------
		return UpdateExperiment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a Experiment
//----------------------------------------------------------------------------
func AddVariantsToExperiment ( experimentId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Experiment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperiment(experimentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Experiment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Experiment)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExperimentVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExperimentVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Variants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "unassignVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Experiment from the gorm
		//----------------------------------------------------------------------------
		return GetExperiment(experimentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a Experiment
//----------------------------------------------------------------------------
func RemoveVariantsFromExperiment( experimentId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Experiment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExperiment(experimentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Experiment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Experiment)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ExperimentVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ExperimentVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExperimentVariantObj from the Variants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "removeVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Experiment from the gorm
		//----------------------------------------------------------------------------
		return GetExperiment(experimentId)

	} else {
		return parentRequestResult
	}
}

