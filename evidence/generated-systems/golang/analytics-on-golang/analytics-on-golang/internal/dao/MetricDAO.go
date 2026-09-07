package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MetricDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMetric - creates a new db entry
//----------------------------------------------------------------------------
func CreateMetric(obj model.Metric)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Metric with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Metric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMetric", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMetric - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMetric(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Metric

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Metric with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Metric using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Metric using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMetric - returns all
//----------------------------------------------------------------------------
func GetAllMetric()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Metric

	//----------------------------------------------------------------------------
	// Request the ORM to find all Metric
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Metric" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Metric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMetric", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMetric - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMetric(obj model.Metric)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Metric using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Metric using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMetric - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMetric(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMetric(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Metric)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Metric using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Metric using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMetric", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SemanticModel on a Metric
//----------------------------------------------------------------------------
func AssignSemanticModelToMetric( metricId uint64, semanticModelId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SemanticModel

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SemanticModel with a
		// matching semanticModelId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, semanticModelId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SemanticModel	to the Metric
			//----------------------------------------------------------------------------
			parentObj.SemanticModel = &childObj

			//----------------------------------------------------------------------------
			// save the Metric
			//----------------------------------------------------------------------------
			return UpdateMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModel", semanticModelId )
			return utils.RequestResult{false, msg, "assignSemanticModel", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SemanticModel on a Metric
//----------------------------------------------------------------------------
func UnassignSemanticModelFromMetric(metricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		//----------------------------------------------------------------------------
		// assign an empty SemanticModel to the SemanticModel
		//----------------------------------------------------------------------------
		parentObj.SemanticModel = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SemanticModel
		//----------------------------------------------------------------------------
		parentObj.SemanticModelId = nil;

		//----------------------------------------------------------------------------
		// save the Metric
		//----------------------------------------------------------------------------
		return UpdateMetric(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Metric
//----------------------------------------------------------------------------
func AddDatasetsToMetric ( metricId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

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
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Metric
//----------------------------------------------------------------------------
func RemoveDatasetsFromMetric( metricId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

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
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more glossaryTermsIds as a GlossaryTerms to a Metric
//----------------------------------------------------------------------------
func AddGlossaryTermsToMetric ( metricId uint64, glossaryTermsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		// slice the ids on comma with no spaces
		ids := strings.Split( glossaryTermsIds, ",")

		for _, glossaryTermsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessGlossaryTerm

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessGlossaryTerm
			// with a matching glossaryTermsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , glossaryTermsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the GlossaryTerms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GlossaryTerms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GlossaryTerms", glossaryTermsId )
				return utils.RequestResult{false, msg, "unassignGlossaryTerms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more glossaryTermsIds as a GlossaryTerms from a Metric
//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromMetric( metricId uint64, glossaryTermsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		// slice the ids on comma with no spaces
		ids := strings.Split( glossaryTermsIds, ",")

		for _, glossaryTermsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessGlossaryTerm

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessGlossaryTerm
			// with a matching glossaryTermsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , glossaryTermsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BusinessGlossaryTermObj from the GlossaryTerms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GlossaryTerms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GlossaryTerms", glossaryTermsId )
				return utils.RequestResult{false, msg, "removeGlossaryTerms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a Metric
//----------------------------------------------------------------------------
func AddAlertsToMetric ( metricId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

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
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a Metric
//----------------------------------------------------------------------------
func RemoveAlertsFromMetric( metricId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

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
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more visualizationsIds as a Visualizations to a Metric
//----------------------------------------------------------------------------
func AddVisualizationsToMetric ( metricId uint64, visualizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		// slice the ids on comma with no spaces
		ids := strings.Split( visualizationsIds, ",")

		for _, visualizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Visualization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Visualization
			// with a matching visualizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , visualizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Visualizations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Visualizations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Visualizations", visualizationsId )
				return utils.RequestResult{false, msg, "unassignVisualizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more visualizationsIds as a Visualizations from a Metric
//----------------------------------------------------------------------------
func RemoveVisualizationsFromMetric( metricId uint64, visualizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Metric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMetric(metricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Metric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Metric)

		// slice the ids on comma with no spaces
		ids := strings.Split( visualizationsIds, ",")

		for _, visualizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Visualization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Visualization
			// with a matching visualizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , visualizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove VisualizationObj from the Visualizations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Visualizations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Visualizations", visualizationsId )
				return utils.RequestResult{false, msg, "removeVisualizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Metric from the gorm
		//----------------------------------------------------------------------------
		return GetMetric(metricId)

	} else {
		return parentRequestResult
	}
}

