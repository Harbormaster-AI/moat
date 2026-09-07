package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QualityCheckDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQualityCheck - creates a new db entry
//----------------------------------------------------------------------------
func CreateQualityCheck(obj model.QualityCheck)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a QualityCheck with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a QualityCheck", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQualityCheck", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQualityCheck - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQualityCheck(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.QualityCheck

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a QualityCheck with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a QualityCheck using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a QualityCheck using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQualityCheck", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQualityCheck - returns all
//----------------------------------------------------------------------------
func GetAllQualityCheck()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.QualityCheck

	//----------------------------------------------------------------------------
	// Request the ORM to find all QualityCheck
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all QualityCheck" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all QualityCheck", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQualityCheck", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQualityCheck - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQualityCheck(obj model.QualityCheck)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a QualityCheck using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a QualityCheck using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQualityCheck", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQualityCheck - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQualityCheck(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the QualityCheck with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQualityCheck(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.QualityCheck)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a QualityCheck using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a QualityCheck using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQualityCheck", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Rule on a QualityCheck
//----------------------------------------------------------------------------
func AssignRuleToQualityCheck( qualityCheckId uint64, ruleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QualityCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityCheck(qualityCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityCheck)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.QualityRule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a QualityRule with a
		// matching ruleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ruleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Rule	to the QualityCheck
			//----------------------------------------------------------------------------
			parentObj.Rule = &childObj

			//----------------------------------------------------------------------------
			// save the QualityCheck
			//----------------------------------------------------------------------------
			return UpdateQualityCheck(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rule", ruleId )
			return utils.RequestResult{false, msg, "assignRule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Rule on a QualityCheck
//----------------------------------------------------------------------------
func UnassignRuleFromQualityCheck(qualityCheckId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QualityCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityCheck(qualityCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityCheck)

		//----------------------------------------------------------------------------
		// assign an empty QualityRule to the Rule
		//----------------------------------------------------------------------------
		parentObj.Rule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Rule
		//----------------------------------------------------------------------------
		parentObj.RuleId = nil;

		//----------------------------------------------------------------------------
		// save the QualityCheck
		//----------------------------------------------------------------------------
		return UpdateQualityCheck(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a QualityCheck
//----------------------------------------------------------------------------
func AssignDatasetToQualityCheck( qualityCheckId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QualityCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityCheck(qualityCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityCheck)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DataSet

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DataSet with a
		// matching datasetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, datasetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dataset	to the QualityCheck
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the QualityCheck
			//----------------------------------------------------------------------------
			return UpdateQualityCheck(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a QualityCheck
//----------------------------------------------------------------------------
func UnassignDatasetFromQualityCheck(qualityCheckId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QualityCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityCheck(qualityCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityCheck)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the QualityCheck
		//----------------------------------------------------------------------------
		return UpdateQualityCheck(parentObj)

	} else {
		return parentRequestResult
	}

}


