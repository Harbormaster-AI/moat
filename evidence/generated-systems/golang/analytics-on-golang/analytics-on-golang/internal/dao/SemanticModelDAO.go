package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SemanticModelDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSemanticModel - creates a new db entry
//----------------------------------------------------------------------------
func CreateSemanticModel(obj model.SemanticModel)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SemanticModel with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SemanticModel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSemanticModel", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSemanticModel - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSemanticModel(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SemanticModel

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SemanticModel with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SemanticModel using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SemanticModel using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSemanticModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSemanticModel - returns all
//----------------------------------------------------------------------------
func GetAllSemanticModel()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SemanticModel

	//----------------------------------------------------------------------------
	// Request the ORM to find all SemanticModel
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SemanticModel" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SemanticModel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSemanticModel", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSemanticModel - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSemanticModel(obj model.SemanticModel)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SemanticModel using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SemanticModel using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSemanticModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSemanticModel - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSemanticModel(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSemanticModel(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SemanticModel)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SemanticModel using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SemanticModel using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSemanticModel", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a SemanticModel
//----------------------------------------------------------------------------
func AddDatasetsToSemanticModel ( semanticModelId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a SemanticModel
//----------------------------------------------------------------------------
func RemoveDatasetsFromSemanticModel( semanticModelId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more metricsIds as a Metrics to a SemanticModel
//----------------------------------------------------------------------------
func AddMetricsToSemanticModel ( semanticModelId uint64, metricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more metricsIds as a Metrics from a SemanticModel
//----------------------------------------------------------------------------
func RemoveMetricsFromSemanticModel( semanticModelId uint64, metricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dimensionsIds as a Dimensions to a SemanticModel
//----------------------------------------------------------------------------
func AddDimensionsToSemanticModel ( semanticModelId uint64, dimensionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dimensionsIds as a Dimensions from a SemanticModel
//----------------------------------------------------------------------------
func RemoveDimensionsFromSemanticModel( semanticModelId uint64, dimensionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more measuresIds as a Measures to a SemanticModel
//----------------------------------------------------------------------------
func AddMeasuresToSemanticModel ( semanticModelId uint64, measuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more measuresIds as a Measures from a SemanticModel
//----------------------------------------------------------------------------
func RemoveMeasuresFromSemanticModel( semanticModelId uint64, measuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more glossaryTermsIds as a GlossaryTerms to a SemanticModel
//----------------------------------------------------------------------------
func AddGlossaryTermsToSemanticModel ( semanticModelId uint64, glossaryTermsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more glossaryTermsIds as a GlossaryTerms from a SemanticModel
//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromSemanticModel( semanticModelId uint64, glossaryTermsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SemanticModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSemanticModel(semanticModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SemanticModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SemanticModel)

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
		// retrieve the modified SemanticModel from the gorm
		//----------------------------------------------------------------------------
		return GetSemanticModel(semanticModelId)

	} else {
		return parentRequestResult
	}
}

