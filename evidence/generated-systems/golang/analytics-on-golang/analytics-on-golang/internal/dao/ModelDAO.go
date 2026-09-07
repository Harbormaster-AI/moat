package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ModelDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateModel - creates a new db entry
//----------------------------------------------------------------------------
func CreateModel(obj model.Model)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Model with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Model", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateModel", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetModel - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetModel(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Model

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Model with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Model using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Model using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllModel - returns all
//----------------------------------------------------------------------------
func GetAllModel()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Model

	//----------------------------------------------------------------------------
	// Request the ORM to find all Model
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Model" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Model", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllModel", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateModel - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateModel(obj model.Model)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Model using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Model using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteModel - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteModel(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetModel(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Model)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Model using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Model using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteModel", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a Model
//----------------------------------------------------------------------------
func AssignWorkspaceToModel( modelId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

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
			// assign the Workspace	to the Model
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the Model
			//----------------------------------------------------------------------------
			return UpdateModel(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a Model
//----------------------------------------------------------------------------
func UnassignWorkspaceFromModel(modelId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the Model
		//----------------------------------------------------------------------------
		return UpdateModel(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more versionsIds as a Versions to a Model
//----------------------------------------------------------------------------
func AddVersionsToModel ( modelId uint64, versionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( versionsIds, ",")

		for _, versionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ModelVersion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ModelVersion
			// with a matching versionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , versionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Versions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Versions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Versions", versionsId )
				return utils.RequestResult{false, msg, "unassignVersions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more versionsIds as a Versions from a Model
//----------------------------------------------------------------------------
func RemoveVersionsFromModel( modelId uint64, versionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( versionsIds, ",")

		for _, versionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ModelVersion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ModelVersion
			// with a matching versionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , versionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ModelVersionObj from the Versions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Versions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Versions", versionsId )
				return utils.RequestResult{false, msg, "removeVersions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featureSetsIds as a FeatureSets to a Model
//----------------------------------------------------------------------------
func AddFeatureSetsToModel ( modelId uint64, featureSetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeatureSets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "unassignFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featureSetsIds as a FeatureSets from a Model
//----------------------------------------------------------------------------
func RemoveFeatureSetsFromModel( modelId uint64, featureSetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureSetObj from the FeatureSets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "removeFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more experimentsIds as a Experiments to a Model
//----------------------------------------------------------------------------
func AddExperimentsToModel ( modelId uint64, experimentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

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
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more experimentsIds as a Experiments from a Model
//----------------------------------------------------------------------------
func RemoveExperimentsFromModel( modelId uint64, experimentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

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
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more tagsIds as a Tags to a Model
//----------------------------------------------------------------------------
func AddTagsToModel ( modelId uint64, tagsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( tagsIds, ",")

		for _, tagsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Tag

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Tag
			// with a matching tagsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tagsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Tags using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tags").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tags", tagsId )
				return utils.RequestResult{false, msg, "unassignTags", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tagsIds as a Tags from a Model
//----------------------------------------------------------------------------
func RemoveTagsFromModel( modelId uint64, tagsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Model with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModel(modelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Model so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Model)

		// slice the ids on comma with no spaces
		ids := strings.Split( tagsIds, ",")

		for _, tagsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Tag

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Tag
			// with a matching tagsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tagsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TagObj from the Tags array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tags").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tags", tagsId )
				return utils.RequestResult{false, msg, "removeTags", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Model from the gorm
		//----------------------------------------------------------------------------
		return GetModel(modelId)

	} else {
		return parentRequestResult
	}
}

