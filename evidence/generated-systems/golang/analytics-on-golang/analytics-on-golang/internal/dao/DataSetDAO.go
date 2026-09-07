package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataSetDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataSet - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataSet(obj model.DataSet)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataSet with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataSet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataSet", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataSet - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataSet(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataSet

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataSet with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataSet using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataSet using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataSet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataSet - returns all
//----------------------------------------------------------------------------
func GetAllDataSet()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataSet

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataSet
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataSet" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataSet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataSet", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataSet - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataSet(obj model.DataSet)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataSet using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataSet using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataSet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataSet - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataSet(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataSet(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataSet)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataSet using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataSet using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataSet", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a DataSet
//----------------------------------------------------------------------------
func AssignWorkspaceToDataSet( dataSetId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

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
			// assign the Workspace	to the DataSet
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the DataSet
			//----------------------------------------------------------------------------
			return UpdateDataSet(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a DataSet
//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataSet(dataSetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the DataSet
		//----------------------------------------------------------------------------
		return UpdateDataSet(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineageNode on a DataSet
//----------------------------------------------------------------------------
func AssignLineageNodeToDataSet( dataSetId uint64, lineageNodeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LineageNode

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LineageNode with a
		// matching lineageNodeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lineageNodeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LineageNode	to the DataSet
			//----------------------------------------------------------------------------
			parentObj.LineageNode = &childObj

			//----------------------------------------------------------------------------
			// save the DataSet
			//----------------------------------------------------------------------------
			return UpdateDataSet(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineageNode", lineageNodeId )
			return utils.RequestResult{false, msg, "assignLineageNode", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineageNode on a DataSet
//----------------------------------------------------------------------------
func UnassignLineageNodeFromDataSet(dataSetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		//----------------------------------------------------------------------------
		// assign an empty LineageNode to the LineageNode
		//----------------------------------------------------------------------------
		parentObj.LineageNode = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineageNode
		//----------------------------------------------------------------------------
		parentObj.LineageNodeId = nil;

		//----------------------------------------------------------------------------
		// save the DataSet
		//----------------------------------------------------------------------------
		return UpdateDataSet(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more sourcesIds as a Sources to a DataSet
//----------------------------------------------------------------------------
func AddSourcesToDataSet ( dataSetId uint64, sourcesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourcesIds, ",")

		for _, sourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching sourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Sources using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Sources").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sources", sourcesId )
				return utils.RequestResult{false, msg, "unassignSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more sourcesIds as a Sources from a DataSet
//----------------------------------------------------------------------------
func RemoveSourcesFromDataSet( dataSetId uint64, sourcesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourcesIds, ",")

		for _, sourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching sourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSourceObj from the Sources array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Sources").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sources", sourcesId )
				return utils.RequestResult{false, msg, "removeSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more pipelinesIds as a Pipelines to a DataSet
//----------------------------------------------------------------------------
func AddPipelinesToDataSet ( dataSetId uint64, pipelinesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( pipelinesIds, ",")

		for _, pipelinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataPipeline

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataPipeline
			// with a matching pipelinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pipelinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Pipelines using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pipelines").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pipelines", pipelinesId )
				return utils.RequestResult{false, msg, "unassignPipelines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pipelinesIds as a Pipelines from a DataSet
//----------------------------------------------------------------------------
func RemovePipelinesFromDataSet( dataSetId uint64, pipelinesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( pipelinesIds, ",")

		for _, pipelinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataPipeline

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataPipeline
			// with a matching pipelinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pipelinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataPipelineObj from the Pipelines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pipelines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pipelines", pipelinesId )
				return utils.RequestResult{false, msg, "removePipelines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more semanticModelsIds as a SemanticModels to a DataSet
//----------------------------------------------------------------------------
func AddSemanticModelsToDataSet ( dataSetId uint64, semanticModelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( semanticModelsIds, ",")

		for _, semanticModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SemanticModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SemanticModel
			// with a matching semanticModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , semanticModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SemanticModels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SemanticModels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModels", semanticModelsId )
				return utils.RequestResult{false, msg, "unassignSemanticModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more semanticModelsIds as a SemanticModels from a DataSet
//----------------------------------------------------------------------------
func RemoveSemanticModelsFromDataSet( dataSetId uint64, semanticModelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( semanticModelsIds, ",")

		for _, semanticModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SemanticModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SemanticModel
			// with a matching semanticModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , semanticModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SemanticModelObj from the SemanticModels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SemanticModels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModels", semanticModelsId )
				return utils.RequestResult{false, msg, "removeSemanticModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dimensionsIds as a Dimensions to a DataSet
//----------------------------------------------------------------------------
func AddDimensionsToDataSet ( dataSetId uint64, dimensionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( dimensionsIds, ",")

		for _, dimensionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dimension

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dimension
			// with a matching dimensionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dimensionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dimensions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dimensions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dimensions", dimensionsId )
				return utils.RequestResult{false, msg, "unassignDimensions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dimensionsIds as a Dimensions from a DataSet
//----------------------------------------------------------------------------
func RemoveDimensionsFromDataSet( dataSetId uint64, dimensionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( dimensionsIds, ",")

		for _, dimensionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dimension

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dimension
			// with a matching dimensionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dimensionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DimensionObj from the Dimensions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dimensions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dimensions", dimensionsId )
				return utils.RequestResult{false, msg, "removeDimensions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more measuresIds as a Measures to a DataSet
//----------------------------------------------------------------------------
func AddMeasuresToDataSet ( dataSetId uint64, measuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( measuresIds, ",")

		for _, measuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Measure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Measure
			// with a matching measuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , measuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Measures using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Measures").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Measures", measuresId )
				return utils.RequestResult{false, msg, "unassignMeasures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more measuresIds as a Measures from a DataSet
//----------------------------------------------------------------------------
func RemoveMeasuresFromDataSet( dataSetId uint64, measuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( measuresIds, ",")

		for _, measuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Measure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Measure
			// with a matching measuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , measuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MeasureObj from the Measures array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Measures").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Measures", measuresId )
				return utils.RequestResult{false, msg, "removeMeasures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more metricsIds as a Metrics to a DataSet
//----------------------------------------------------------------------------
func AddMetricsToDataSet ( dataSetId uint64, metricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

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
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more metricsIds as a Metrics from a DataSet
//----------------------------------------------------------------------------
func RemoveMetricsFromDataSet( dataSetId uint64, metricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

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
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more qualityRulesIds as a QualityRules to a DataSet
//----------------------------------------------------------------------------
func AddQualityRulesToDataSet ( dataSetId uint64, qualityRulesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( qualityRulesIds, ",")

		for _, qualityRulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualityRule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualityRule
			// with a matching qualityRulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , qualityRulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the QualityRules using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QualityRules").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QualityRules", qualityRulesId )
				return utils.RequestResult{false, msg, "unassignQualityRules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more qualityRulesIds as a QualityRules from a DataSet
//----------------------------------------------------------------------------
func RemoveQualityRulesFromDataSet( dataSetId uint64, qualityRulesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

		// slice the ids on comma with no spaces
		ids := strings.Split( qualityRulesIds, ",")

		for _, qualityRulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualityRule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualityRule
			// with a matching qualityRulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , qualityRulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QualityRuleObj from the QualityRules array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QualityRules").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QualityRules", qualityRulesId )
				return utils.RequestResult{false, msg, "removeQualityRules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more tagsIds as a Tags to a DataSet
//----------------------------------------------------------------------------
func AddTagsToDataSet ( dataSetId uint64, tagsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

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
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tagsIds as a Tags from a DataSet
//----------------------------------------------------------------------------
func RemoveTagsFromDataSet( dataSetId uint64, tagsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSet(dataSetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSet)

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
		// retrieve the modified DataSet from the gorm
		//----------------------------------------------------------------------------
		return GetDataSet(dataSetId)

	} else {
		return parentRequestResult
	}
}

