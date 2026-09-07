package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DimensionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDimension - creates a new db entry
//----------------------------------------------------------------------------
func CreateDimension(obj model.Dimension)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Dimension with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Dimension", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDimension", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDimension - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDimension(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Dimension

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Dimension with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Dimension using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Dimension using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDimension", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDimension - returns all
//----------------------------------------------------------------------------
func GetAllDimension()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Dimension

	//----------------------------------------------------------------------------
	// Request the ORM to find all Dimension
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Dimension" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dimension", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDimension", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDimension - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDimension(obj model.Dimension)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Dimension using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Dimension using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDimension", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDimension - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDimension(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDimension(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Dimension)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Dimension using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Dimension using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDimension", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SemanticModel on a Dimension
//----------------------------------------------------------------------------
func AssignSemanticModelToDimension( dimensionId uint64, semanticModelId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

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
			// assign the SemanticModel	to the Dimension
			//----------------------------------------------------------------------------
			parentObj.SemanticModel = &childObj

			//----------------------------------------------------------------------------
			// save the Dimension
			//----------------------------------------------------------------------------
			return UpdateDimension(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModel", semanticModelId )
			return utils.RequestResult{false, msg, "assignSemanticModel", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SemanticModel on a Dimension
//----------------------------------------------------------------------------
func UnassignSemanticModelFromDimension(dimensionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

		//----------------------------------------------------------------------------
		// assign an empty SemanticModel to the SemanticModel
		//----------------------------------------------------------------------------
		parentObj.SemanticModel = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SemanticModel
		//----------------------------------------------------------------------------
		parentObj.SemanticModelId = nil;

		//----------------------------------------------------------------------------
		// save the Dimension
		//----------------------------------------------------------------------------
		return UpdateDimension(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Dimension
//----------------------------------------------------------------------------
func AddDatasetsToDimension ( dimensionId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

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
		// retrieve the modified Dimension from the gorm
		//----------------------------------------------------------------------------
		return GetDimension(dimensionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Dimension
//----------------------------------------------------------------------------
func RemoveDatasetsFromDimension( dimensionId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

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
		// retrieve the modified Dimension from the gorm
		//----------------------------------------------------------------------------
		return GetDimension(dimensionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more glossaryTermsIds as a GlossaryTerms to a Dimension
//----------------------------------------------------------------------------
func AddGlossaryTermsToDimension ( dimensionId uint64, glossaryTermsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

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
		// retrieve the modified Dimension from the gorm
		//----------------------------------------------------------------------------
		return GetDimension(dimensionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more glossaryTermsIds as a GlossaryTerms from a Dimension
//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromDimension( dimensionId uint64, glossaryTermsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dimension with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDimension(dimensionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dimension so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dimension)

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
		// retrieve the modified Dimension from the gorm
		//----------------------------------------------------------------------------
		return GetDimension(dimensionId)

	} else {
		return parentRequestResult
	}
}

