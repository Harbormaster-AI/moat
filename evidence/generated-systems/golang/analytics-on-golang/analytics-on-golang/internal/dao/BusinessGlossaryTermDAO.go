package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BusinessGlossaryTermDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBusinessGlossaryTerm - creates a new db entry
//----------------------------------------------------------------------------
func CreateBusinessGlossaryTerm(obj model.BusinessGlossaryTerm)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BusinessGlossaryTerm with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BusinessGlossaryTerm", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBusinessGlossaryTerm", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBusinessGlossaryTerm - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBusinessGlossaryTerm(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BusinessGlossaryTerm

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BusinessGlossaryTerm with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BusinessGlossaryTerm using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BusinessGlossaryTerm using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBusinessGlossaryTerm", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBusinessGlossaryTerm - returns all
//----------------------------------------------------------------------------
func GetAllBusinessGlossaryTerm()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BusinessGlossaryTerm

	//----------------------------------------------------------------------------
	// Request the ORM to find all BusinessGlossaryTerm
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BusinessGlossaryTerm" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BusinessGlossaryTerm", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBusinessGlossaryTerm", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBusinessGlossaryTerm - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBusinessGlossaryTerm(obj model.BusinessGlossaryTerm)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BusinessGlossaryTerm using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BusinessGlossaryTerm using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBusinessGlossaryTerm", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBusinessGlossaryTerm - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBusinessGlossaryTerm(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBusinessGlossaryTerm(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BusinessGlossaryTerm)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BusinessGlossaryTerm using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BusinessGlossaryTerm using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBusinessGlossaryTerm", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more relatedTermsIds as a RelatedTerms to a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func AddRelatedTermsToBusinessGlossaryTerm ( businessGlossaryTermId uint64, relatedTermsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedTermsIds, ",")

		for _, relatedTermsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessGlossaryTerm

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessGlossaryTerm
			// with a matching relatedTermsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedTermsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedTerms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedTerms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedTerms", relatedTermsId )
				return utils.RequestResult{false, msg, "unassignRelatedTerms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedTermsIds as a RelatedTerms from a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func RemoveRelatedTermsFromBusinessGlossaryTerm( businessGlossaryTermId uint64, relatedTermsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedTermsIds, ",")

		for _, relatedTermsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessGlossaryTerm

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessGlossaryTerm
			// with a matching relatedTermsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedTermsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BusinessGlossaryTermObj from the RelatedTerms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedTerms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedTerms", relatedTermsId )
				return utils.RequestResult{false, msg, "removeRelatedTerms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more metricsIds as a Metrics to a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func AddMetricsToBusinessGlossaryTerm ( businessGlossaryTermId uint64, metricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more metricsIds as a Metrics from a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func RemoveMetricsFromBusinessGlossaryTerm( businessGlossaryTermId uint64, metricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func AddDatasetsToBusinessGlossaryTerm ( businessGlossaryTermId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func RemoveDatasetsFromBusinessGlossaryTerm( businessGlossaryTermId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dimensionsIds as a Dimensions to a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func AddDimensionsToBusinessGlossaryTerm ( businessGlossaryTermId uint64, dimensionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dimensionsIds as a Dimensions from a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func RemoveDimensionsFromBusinessGlossaryTerm( businessGlossaryTermId uint64, dimensionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more measuresIds as a Measures to a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func AddMeasuresToBusinessGlossaryTerm ( businessGlossaryTermId uint64, measuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more measuresIds as a Measures from a BusinessGlossaryTerm
//----------------------------------------------------------------------------
func RemoveMeasuresFromBusinessGlossaryTerm( businessGlossaryTermId uint64, measuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessGlossaryTerm with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessGlossaryTerm(businessGlossaryTermId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessGlossaryTerm so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessGlossaryTerm)

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
		// retrieve the modified BusinessGlossaryTerm from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessGlossaryTerm(businessGlossaryTermId)

	} else {
		return parentRequestResult
	}
}

