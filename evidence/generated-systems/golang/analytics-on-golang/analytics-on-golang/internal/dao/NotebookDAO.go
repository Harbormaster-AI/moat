package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing NotebookDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateNotebook - creates a new db entry
//----------------------------------------------------------------------------
func CreateNotebook(obj model.Notebook)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Notebook with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Notebook", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateNotebook", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetNotebook - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetNotebook(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Notebook

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Notebook with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Notebook using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Notebook using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetNotebook", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllNotebook - returns all
//----------------------------------------------------------------------------
func GetAllNotebook()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Notebook

	//----------------------------------------------------------------------------
	// Request the ORM to find all Notebook
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Notebook" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Notebook", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllNotebook", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateNotebook - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateNotebook(obj model.Notebook)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Notebook using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Notebook using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateNotebook", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteNotebook - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteNotebook(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetNotebook(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Notebook)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Notebook using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Notebook using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteNotebook", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a Notebook
//----------------------------------------------------------------------------
func AssignWorkspaceToNotebook( notebookId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

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
			// assign the Workspace	to the Notebook
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the Notebook
			//----------------------------------------------------------------------------
			return UpdateNotebook(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a Notebook
//----------------------------------------------------------------------------
func UnassignWorkspaceFromNotebook(notebookId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the Notebook
		//----------------------------------------------------------------------------
		return UpdateNotebook(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Notebook
//----------------------------------------------------------------------------
func AddDatasetsToNotebook ( notebookId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

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
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Notebook
//----------------------------------------------------------------------------
func RemoveDatasetsFromNotebook( notebookId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

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
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more experimentsIds as a Experiments to a Notebook
//----------------------------------------------------------------------------
func AddExperimentsToNotebook ( notebookId uint64, experimentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

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
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more experimentsIds as a Experiments from a Notebook
//----------------------------------------------------------------------------
func RemoveExperimentsFromNotebook( notebookId uint64, experimentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

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
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more queriesIds as a Queries to a Notebook
//----------------------------------------------------------------------------
func AddQueriesToNotebook ( notebookId uint64, queriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

		// slice the ids on comma with no spaces
		ids := strings.Split( queriesIds, ",")

		for _, queriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BIQuery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BIQuery
			// with a matching queriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , queriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Queries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Queries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Queries", queriesId )
				return utils.RequestResult{false, msg, "unassignQueries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more queriesIds as a Queries from a Notebook
//----------------------------------------------------------------------------
func RemoveQueriesFromNotebook( notebookId uint64, queriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Notebook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNotebook(notebookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Notebook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Notebook)

		// slice the ids on comma with no spaces
		ids := strings.Split( queriesIds, ",")

		for _, queriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BIQuery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BIQuery
			// with a matching queriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , queriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BIQueryObj from the Queries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Queries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Queries", queriesId )
				return utils.RequestResult{false, msg, "removeQueries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Notebook from the gorm
		//----------------------------------------------------------------------------
		return GetNotebook(notebookId)

	} else {
		return parentRequestResult
	}
}

