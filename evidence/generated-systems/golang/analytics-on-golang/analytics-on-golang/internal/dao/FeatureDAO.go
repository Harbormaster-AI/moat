package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FeatureDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFeature - creates a new db entry
//----------------------------------------------------------------------------
func CreateFeature(obj model.Feature)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Feature with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Feature", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFeature", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFeature - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFeature(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Feature

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Feature with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Feature using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Feature using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFeature", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFeature - returns all
//----------------------------------------------------------------------------
func GetAllFeature()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Feature

	//----------------------------------------------------------------------------
	// Request the ORM to find all Feature
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Feature" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Feature", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFeature", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFeature - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFeature(obj model.Feature)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Feature using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Feature using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFeature", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFeature - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFeature(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFeature(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Feature)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Feature using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Feature using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFeature", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a FeatureSet on a Feature
//----------------------------------------------------------------------------
func AssignFeatureSetToFeature( featureId uint64, featureSetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FeatureSet

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FeatureSet with a
		// matching featureSetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, featureSetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FeatureSet	to the Feature
			//----------------------------------------------------------------------------
			parentObj.FeatureSet = &childObj

			//----------------------------------------------------------------------------
			// save the Feature
			//----------------------------------------------------------------------------
			return UpdateFeature(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSet", featureSetId )
			return utils.RequestResult{false, msg, "assignFeatureSet", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FeatureSet on a Feature
//----------------------------------------------------------------------------
func UnassignFeatureSetFromFeature(featureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

		//----------------------------------------------------------------------------
		// assign an empty FeatureSet to the FeatureSet
		//----------------------------------------------------------------------------
		parentObj.FeatureSet = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FeatureSet
		//----------------------------------------------------------------------------
		parentObj.FeatureSetId = nil;

		//----------------------------------------------------------------------------
		// save the Feature
		//----------------------------------------------------------------------------
		return UpdateFeature(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more sourceDatasetsIds as a SourceDatasets to a Feature
//----------------------------------------------------------------------------
func AddSourceDatasetsToFeature ( featureId uint64, sourceDatasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourceDatasetsIds, ",")

		for _, sourceDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching sourceDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourceDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SourceDatasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SourceDatasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceDatasets", sourceDatasetsId )
				return utils.RequestResult{false, msg, "unassignSourceDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more sourceDatasetsIds as a SourceDatasets from a Feature
//----------------------------------------------------------------------------
func RemoveSourceDatasetsFromFeature( featureId uint64, sourceDatasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourceDatasetsIds, ",")

		for _, sourceDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching sourceDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourceDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the SourceDatasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SourceDatasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceDatasets", sourceDatasetsId )
				return utils.RequestResult{false, msg, "removeSourceDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a Feature
//----------------------------------------------------------------------------
func AddModelsToFeature ( featureId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

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
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a Feature
//----------------------------------------------------------------------------
func RemoveModelsFromFeature( featureId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

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
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more trainingRunsIds as a TrainingRuns to a Feature
//----------------------------------------------------------------------------
func AddTrainingRunsToFeature ( featureId uint64, trainingRunsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

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
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more trainingRunsIds as a TrainingRuns from a Feature
//----------------------------------------------------------------------------
func RemoveTrainingRunsFromFeature( featureId uint64, trainingRunsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Feature with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeature(featureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Feature so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Feature)

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
		// retrieve the modified Feature from the gorm
		//----------------------------------------------------------------------------
		return GetFeature(featureId)

	} else {
		return parentRequestResult
	}
}

