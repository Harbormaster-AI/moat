package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QualityRuleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQualityRule - creates a new db entry
//----------------------------------------------------------------------------
func CreateQualityRule(obj model.QualityRule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a QualityRule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a QualityRule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQualityRule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQualityRule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQualityRule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.QualityRule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a QualityRule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a QualityRule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a QualityRule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQualityRule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQualityRule - returns all
//----------------------------------------------------------------------------
func GetAllQualityRule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.QualityRule

	//----------------------------------------------------------------------------
	// Request the ORM to find all QualityRule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all QualityRule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all QualityRule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQualityRule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQualityRule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQualityRule(obj model.QualityRule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a QualityRule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a QualityRule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQualityRule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQualityRule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQualityRule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the QualityRule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQualityRule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityRule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.QualityRule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a QualityRule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a QualityRule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQualityRule", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Dataset on a QualityRule
//----------------------------------------------------------------------------
func AssignDatasetToQualityRule( qualityRuleId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QualityRule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityRule(qualityRuleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityRule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityRule)

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
			// assign the Dataset	to the QualityRule
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the QualityRule
			//----------------------------------------------------------------------------
			return UpdateQualityRule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a QualityRule
//----------------------------------------------------------------------------
func UnassignDatasetFromQualityRule(qualityRuleId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QualityRule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityRule(qualityRuleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityRule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityRule)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the QualityRule
		//----------------------------------------------------------------------------
		return UpdateQualityRule(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more checksIds as a Checks to a QualityRule
//----------------------------------------------------------------------------
func AddChecksToQualityRule ( qualityRuleId uint64, checksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QualityRule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityRule(qualityRuleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityRule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityRule)

		// slice the ids on comma with no spaces
		ids := strings.Split( checksIds, ",")

		for _, checksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualityCheck

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualityCheck
			// with a matching checksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , checksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Checks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Checks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Checks", checksId )
				return utils.RequestResult{false, msg, "unassignChecks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified QualityRule from the gorm
		//----------------------------------------------------------------------------
		return GetQualityRule(qualityRuleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more checksIds as a Checks from a QualityRule
//----------------------------------------------------------------------------
func RemoveChecksFromQualityRule( qualityRuleId uint64, checksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the QualityRule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQualityRule(qualityRuleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QualityRule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QualityRule)

		// slice the ids on comma with no spaces
		ids := strings.Split( checksIds, ",")

		for _, checksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualityCheck

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualityCheck
			// with a matching checksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , checksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QualityCheckObj from the Checks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Checks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Checks", checksId )
				return utils.RequestResult{false, msg, "removeChecks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified QualityRule from the gorm
		//----------------------------------------------------------------------------
		return GetQualityRule(qualityRuleId)

	} else {
		return parentRequestResult
	}
}

