package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FeatureSetDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFeatureSet - creates a new db entry
//----------------------------------------------------------------------------
func CreateFeatureSet(obj model.FeatureSet)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FeatureSet with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FeatureSet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFeatureSet", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFeatureSet - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFeatureSet(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FeatureSet

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FeatureSet with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FeatureSet using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FeatureSet using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFeatureSet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFeatureSet - returns all
//----------------------------------------------------------------------------
func GetAllFeatureSet()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FeatureSet

	//----------------------------------------------------------------------------
	// Request the ORM to find all FeatureSet
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FeatureSet" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FeatureSet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFeatureSet", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFeatureSet - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFeatureSet(obj model.FeatureSet)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FeatureSet using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FeatureSet using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFeatureSet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFeatureSet - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFeatureSet(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFeatureSet(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FeatureSet)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FeatureSet using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FeatureSet using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFeatureSet", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a FeatureSet
//----------------------------------------------------------------------------
func AssignWorkspaceToFeatureSet( featureSetId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
			// assign the Workspace	to the FeatureSet
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the FeatureSet
			//----------------------------------------------------------------------------
			return UpdateFeatureSet(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a FeatureSet
//----------------------------------------------------------------------------
func UnassignWorkspaceFromFeatureSet(featureSetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the FeatureSet
		//----------------------------------------------------------------------------
		return UpdateFeatureSet(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more featuresIds as a Features to a FeatureSet
//----------------------------------------------------------------------------
func AddFeaturesToFeatureSet ( featureSetId uint64, featuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( featuresIds, ",")

		for _, featuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Feature

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Feature
			// with a matching featuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Features using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Features").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Features", featuresId )
				return utils.RequestResult{false, msg, "unassignFeatures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featuresIds as a Features from a FeatureSet
//----------------------------------------------------------------------------
func RemoveFeaturesFromFeatureSet( featureSetId uint64, featuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( featuresIds, ",")

		for _, featuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Feature

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Feature
			// with a matching featuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureObj from the Features array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Features").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Features", featuresId )
				return utils.RequestResult{false, msg, "removeFeatures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a FeatureSet
//----------------------------------------------------------------------------
func AddDatasetsToFeatureSet ( featureSetId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a FeatureSet
//----------------------------------------------------------------------------
func RemoveDatasetsFromFeatureSet( featureSetId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a FeatureSet
//----------------------------------------------------------------------------
func AddModelsToFeatureSet ( featureSetId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a FeatureSet
//----------------------------------------------------------------------------
func RemoveModelsFromFeatureSet( featureSetId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelVersionsIds as a ModelVersions to a FeatureSet
//----------------------------------------------------------------------------
func AddModelVersionsToFeatureSet ( featureSetId uint64, modelVersionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelVersionsIds, ",")

		for _, modelVersionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ModelVersion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ModelVersion
			// with a matching modelVersionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelVersionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ModelVersions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ModelVersions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersions", modelVersionsId )
				return utils.RequestResult{false, msg, "unassignModelVersions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelVersionsIds as a ModelVersions from a FeatureSet
//----------------------------------------------------------------------------
func RemoveModelVersionsFromFeatureSet( featureSetId uint64, modelVersionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelVersionsIds, ",")

		for _, modelVersionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ModelVersion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ModelVersion
			// with a matching modelVersionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelVersionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ModelVersionObj from the ModelVersions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ModelVersions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersions", modelVersionsId )
				return utils.RequestResult{false, msg, "removeModelVersions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more tagsIds as a Tags to a FeatureSet
//----------------------------------------------------------------------------
func AddTagsToFeatureSet ( featureSetId uint64, tagsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tagsIds as a Tags from a FeatureSet
//----------------------------------------------------------------------------
func RemoveTagsFromFeatureSet( featureSetId uint64, tagsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FeatureSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeatureSet(featureSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeatureSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeatureSet)

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
		// retrieve the modified FeatureSet from the gorm
		//----------------------------------------------------------------------------
		return GetFeatureSet(featureSetId)

	} else {
		return parentRequestResult
	}
}

