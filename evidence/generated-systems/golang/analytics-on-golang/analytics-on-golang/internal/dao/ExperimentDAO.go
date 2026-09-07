package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
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
// assigns a Workspace on a Experiment
//----------------------------------------------------------------------------
func AssignWorkspaceToExperiment( experimentId uint64, workspaceId uint64 )(utils.RequestResult){

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
		var childObj model.AnalyticsWorkspace

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AnalyticsWorkspace with a
		// matching workspaceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workspaceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Workspace	to the Experiment
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the Experiment
			//----------------------------------------------------------------------------
			return UpdateExperiment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a Experiment
//----------------------------------------------------------------------------
func UnassignWorkspaceFromExperiment(experimentId uint64)(utils.RequestResult) {

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
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the Experiment
		//----------------------------------------------------------------------------
		return UpdateExperiment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more trainingRunsIds as a TrainingRuns to a Experiment
//----------------------------------------------------------------------------
func AddTrainingRunsToExperiment ( experimentId uint64, trainingRunsIds string )(utils.RequestResult) {

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
		ids := strings.Split( trainingRunsIds, ",")

		for _, trainingRunsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingRun

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingRun
			// with a matching trainingRunsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingRunsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TrainingRuns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingRuns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRuns", trainingRunsId )
				return utils.RequestResult{false, msg, "unassignTrainingRuns", childObj}
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
// removes one or more trainingRunsIds as a TrainingRuns from a Experiment
//----------------------------------------------------------------------------
func RemoveTrainingRunsFromExperiment( experimentId uint64, trainingRunsIds string )(utils.RequestResult) {
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
		ids := strings.Split( trainingRunsIds, ",")

		for _, trainingRunsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingRun

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingRun
			// with a matching trainingRunsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingRunsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrainingRunObj from the TrainingRuns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingRuns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRuns", trainingRunsId )
				return utils.RequestResult{false, msg, "removeTrainingRuns", childObj}
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
// adds one or more modelsIds as a Models to a Experiment
//----------------------------------------------------------------------------
func AddModelsToExperiment ( experimentId uint64, modelsIds string )(utils.RequestResult) {

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
		// retrieve the modified Experiment from the gorm
		//----------------------------------------------------------------------------
		return GetExperiment(experimentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a Experiment
//----------------------------------------------------------------------------
func RemoveModelsFromExperiment( experimentId uint64, modelsIds string )(utils.RequestResult) {
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
		// retrieve the modified Experiment from the gorm
		//----------------------------------------------------------------------------
		return GetExperiment(experimentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notebooksIds as a Notebooks to a Experiment
//----------------------------------------------------------------------------
func AddNotebooksToExperiment ( experimentId uint64, notebooksIds string )(utils.RequestResult) {

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
		ids := strings.Split( notebooksIds, ",")

		for _, notebooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Notebook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Notebook
			// with a matching notebooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notebooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Notebooks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notebooks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notebooks", notebooksId )
				return utils.RequestResult{false, msg, "unassignNotebooks", childObj}
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
// removes one or more notebooksIds as a Notebooks from a Experiment
//----------------------------------------------------------------------------
func RemoveNotebooksFromExperiment( experimentId uint64, notebooksIds string )(utils.RequestResult) {
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
		ids := strings.Split( notebooksIds, ",")

		for _, notebooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Notebook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Notebook
			// with a matching notebooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notebooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove NotebookObj from the Notebooks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notebooks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notebooks", notebooksId )
				return utils.RequestResult{false, msg, "removeNotebooks", childObj}
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

