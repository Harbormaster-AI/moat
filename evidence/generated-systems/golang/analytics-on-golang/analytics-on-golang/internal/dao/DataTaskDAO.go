package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataTaskDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataTask - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataTask(obj model.DataTask)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataTask with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataTask", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataTask - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataTask(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataTask

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataTask with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataTask using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataTask using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataTask - returns all
//----------------------------------------------------------------------------
func GetAllDataTask()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataTask

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataTask
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataTask" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataTask", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataTask - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataTask(obj model.DataTask)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataTask using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataTask using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataTask - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataTask(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataTask(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataTask)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataTask using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataTask using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataTask", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Pipeline on a DataTask
//----------------------------------------------------------------------------
func AssignPipelineToDataTask( dataTaskId uint64, pipelineId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DataPipeline

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DataPipeline with a
		// matching pipelineId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, pipelineId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Pipeline	to the DataTask
			//----------------------------------------------------------------------------
			parentObj.Pipeline = &childObj

			//----------------------------------------------------------------------------
			// save the DataTask
			//----------------------------------------------------------------------------
			return UpdateDataTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pipeline", pipelineId )
			return utils.RequestResult{false, msg, "assignPipeline", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Pipeline on a DataTask
//----------------------------------------------------------------------------
func UnassignPipelineFromDataTask(dataTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		//----------------------------------------------------------------------------
		// assign an empty DataPipeline to the Pipeline
		//----------------------------------------------------------------------------
		parentObj.Pipeline = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Pipeline
		//----------------------------------------------------------------------------
		parentObj.PipelineId = nil;

		//----------------------------------------------------------------------------
		// save the DataTask
		//----------------------------------------------------------------------------
		return UpdateDataTask(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inputDatasetsIds as a InputDatasets to a DataTask
//----------------------------------------------------------------------------
func AddInputDatasetsToDataTask ( dataTaskId uint64, inputDatasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputDatasetsIds, ",")

		for _, inputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching inputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InputDatasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InputDatasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InputDatasets", inputDatasetsId )
				return utils.RequestResult{false, msg, "unassignInputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataTask from the gorm
		//----------------------------------------------------------------------------
		return GetDataTask(dataTaskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inputDatasetsIds as a InputDatasets from a DataTask
//----------------------------------------------------------------------------
func RemoveInputDatasetsFromDataTask( dataTaskId uint64, inputDatasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputDatasetsIds, ",")

		for _, inputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching inputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the InputDatasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InputDatasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InputDatasets", inputDatasetsId )
				return utils.RequestResult{false, msg, "removeInputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataTask from the gorm
		//----------------------------------------------------------------------------
		return GetDataTask(dataTaskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more outputDatasetsIds as a OutputDatasets to a DataTask
//----------------------------------------------------------------------------
func AddOutputDatasetsToDataTask ( dataTaskId uint64, outputDatasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputDatasetsIds, ",")

		for _, outputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching outputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OutputDatasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OutputDatasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OutputDatasets", outputDatasetsId )
				return utils.RequestResult{false, msg, "unassignOutputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataTask from the gorm
		//----------------------------------------------------------------------------
		return GetDataTask(dataTaskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more outputDatasetsIds as a OutputDatasets from a DataTask
//----------------------------------------------------------------------------
func RemoveOutputDatasetsFromDataTask( dataTaskId uint64, outputDatasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataTask(dataTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataTask)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputDatasetsIds, ",")

		for _, outputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching outputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the OutputDatasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OutputDatasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OutputDatasets", outputDatasetsId )
				return utils.RequestResult{false, msg, "removeOutputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataTask from the gorm
		//----------------------------------------------------------------------------
		return GetDataTask(dataTaskId)

	} else {
		return parentRequestResult
	}
}

