package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RecommendationScenarioDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRecommendationScenario - creates a new db entry
//----------------------------------------------------------------------------
func CreateRecommendationScenario(obj model.RecommendationScenario)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RecommendationScenario with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RecommendationScenario", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRecommendationScenario", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRecommendationScenario - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRecommendationScenario(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RecommendationScenario

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RecommendationScenario with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RecommendationScenario using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RecommendationScenario using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRecommendationScenario", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRecommendationScenario - returns all
//----------------------------------------------------------------------------
func GetAllRecommendationScenario()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RecommendationScenario

	//----------------------------------------------------------------------------
	// Request the ORM to find all RecommendationScenario
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RecommendationScenario" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RecommendationScenario", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRecommendationScenario", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRecommendationScenario - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRecommendationScenario(obj model.RecommendationScenario)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RecommendationScenario using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RecommendationScenario using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRecommendationScenario", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRecommendationScenario - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRecommendationScenario(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRecommendationScenario(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RecommendationScenario)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RecommendationScenario using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RecommendationScenario using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRecommendationScenario", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a RecommendationScenario
//----------------------------------------------------------------------------
func AddModelsToRecommendationScenario ( recommendationScenarioId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Models using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "unassignModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a RecommendationScenario
//----------------------------------------------------------------------------
func RemoveModelsFromRecommendationScenario( recommendationScenarioId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Model_Obj from the Models array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "removeModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a RecommendationScenario
//----------------------------------------------------------------------------
func AddDatasetsToRecommendationScenario ( recommendationScenarioId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Datasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "unassignDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a RecommendationScenario
//----------------------------------------------------------------------------
func RemoveDatasetsFromRecommendationScenario( recommendationScenarioId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the Datasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "removeDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more experimentsIds as a Experiments to a RecommendationScenario
//----------------------------------------------------------------------------
func AddExperimentsToRecommendationScenario ( recommendationScenarioId uint64, experimentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( experimentsIds, ",")

		for _, experimentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Experiment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Experiment
			// with a matching experimentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , experimentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Experiments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Experiments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Experiments", experimentsId )
				return utils.RequestResult{false, msg, "unassignExperiments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more experimentsIds as a Experiments from a RecommendationScenario
//----------------------------------------------------------------------------
func RemoveExperimentsFromRecommendationScenario( recommendationScenarioId uint64, experimentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( experimentsIds, ",")

		for _, experimentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Experiment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Experiment
			// with a matching experimentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , experimentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExperimentObj from the Experiments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Experiments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Experiments", experimentsId )
				return utils.RequestResult{false, msg, "removeExperiments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a RecommendationScenario
//----------------------------------------------------------------------------
func AddAlertsToRecommendationScenario ( recommendationScenarioId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Alerts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "unassignAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a RecommendationScenario
//----------------------------------------------------------------------------
func RemoveAlertsFromRecommendationScenario( recommendationScenarioId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecommendationScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecommendationScenario(recommendationScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecommendationScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecommendationScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AlertObj from the Alerts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "removeAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecommendationScenario from the gorm
		//----------------------------------------------------------------------------
		return GetRecommendationScenario(recommendationScenarioId)

	} else {
		return parentRequestResult
	}
}

