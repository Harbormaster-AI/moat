package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TagDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTag - creates a new db entry
//----------------------------------------------------------------------------
func CreateTag(obj model.Tag)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Tag with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Tag", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTag", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTag - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTag(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Tag

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Tag with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Tag using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Tag using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTag", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTag - returns all
//----------------------------------------------------------------------------
func GetAllTag()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Tag

	//----------------------------------------------------------------------------
	// Request the ORM to find all Tag
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Tag" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Tag", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTag", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTag - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTag(obj model.Tag)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Tag using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Tag using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTag", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTag - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTag(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTag(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Tag)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Tag using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Tag using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTag", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Tag
//----------------------------------------------------------------------------
func AddDatasetsToTag ( tagId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Tag
//----------------------------------------------------------------------------
func RemoveDatasetsFromTag( tagId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a Tag
//----------------------------------------------------------------------------
func AddModelsToTag ( tagId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a Tag
//----------------------------------------------------------------------------
func RemoveModelsFromTag( tagId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelVersionsIds as a ModelVersions to a Tag
//----------------------------------------------------------------------------
func AddModelVersionsToTag ( tagId uint64, modelVersionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelVersionsIds as a ModelVersions from a Tag
//----------------------------------------------------------------------------
func RemoveModelVersionsFromTag( tagId uint64, modelVersionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dashboardsIds as a Dashboards to a Tag
//----------------------------------------------------------------------------
func AddDashboardsToTag ( tagId uint64, dashboardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dashboards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "unassignDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dashboardsIds as a Dashboards from a Tag
//----------------------------------------------------------------------------
func RemoveDashboardsFromTag( tagId uint64, dashboardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DashboardObj from the Dashboards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "removeDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a Tag
//----------------------------------------------------------------------------
func AddReportsToTag ( tagId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "unassignReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a Tag
//----------------------------------------------------------------------------
func RemoveReportsFromTag( tagId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReportObj from the Reports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "removeReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featureSetsIds as a FeatureSets to a Tag
//----------------------------------------------------------------------------
func AddFeatureSetsToTag ( tagId uint64, featureSetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featureSetsIds as a FeatureSets from a Tag
//----------------------------------------------------------------------------
func RemoveFeatureSetsFromTag( tagId uint64, featureSetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

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
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more metricsIds as a Metrics to a Tag
//----------------------------------------------------------------------------
func AddMetricsToTag ( tagId uint64, metricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( metricsIds, ",")

		for _, metricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Metric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Metric
			// with a matching metricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , metricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Metrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Metrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metrics", metricsId )
				return utils.RequestResult{false, msg, "unassignMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more metricsIds as a Metrics from a Tag
//----------------------------------------------------------------------------
func RemoveMetricsFromTag( tagId uint64, metricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tag with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTag(tagId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tag so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Tag)

		// slice the ids on comma with no spaces
		ids := strings.Split( metricsIds, ",")

		for _, metricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Metric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Metric
			// with a matching metricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , metricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MetricObj from the Metrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Metrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metrics", metricsId )
				return utils.RequestResult{false, msg, "removeMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Tag from the gorm
		//----------------------------------------------------------------------------
		return GetTag(tagId)

	} else {
		return parentRequestResult
	}
}

