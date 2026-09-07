package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InspectionResultDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInspectionResult - creates a new db entry
//----------------------------------------------------------------------------
func CreateInspectionResult(obj model.InspectionResult)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InspectionResult with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InspectionResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInspectionResult", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInspectionResult - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInspectionResult(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InspectionResult

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InspectionResult with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InspectionResult using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InspectionResult using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInspectionResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInspectionResult - returns all
//----------------------------------------------------------------------------
func GetAllInspectionResult()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InspectionResult

	//----------------------------------------------------------------------------
	// Request the ORM to find all InspectionResult
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InspectionResult" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InspectionResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInspectionResult", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInspectionResult - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInspectionResult(obj model.InspectionResult)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InspectionResult using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InspectionResult using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInspectionResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInspectionResult - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInspectionResult(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InspectionResult with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInspectionResult(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InspectionResult)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InspectionResult using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InspectionResult using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInspectionResult", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a InspectionLot on a InspectionResult
//----------------------------------------------------------------------------
func AssignInspectionLotToInspectionResult( inspectionResultId uint64, inspectionLotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionResult(inspectionResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionResult)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InspectionLot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InspectionLot with a
		// matching inspectionLotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inspectionLotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InspectionLot	to the InspectionResult
			//----------------------------------------------------------------------------
			parentObj.InspectionLot = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionResult
			//----------------------------------------------------------------------------
			return UpdateInspectionResult(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InspectionLot", inspectionLotId )
			return utils.RequestResult{false, msg, "assignInspectionLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InspectionLot on a InspectionResult
//----------------------------------------------------------------------------
func UnassignInspectionLotFromInspectionResult(inspectionResultId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionResult(inspectionResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionResult)

		//----------------------------------------------------------------------------
		// assign an empty InspectionLot to the InspectionLot
		//----------------------------------------------------------------------------
		parentObj.InspectionLot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InspectionLot
		//----------------------------------------------------------------------------
		parentObj.InspectionLotId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionResult
		//----------------------------------------------------------------------------
		return UpdateInspectionResult(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Characteristic on a InspectionResult
//----------------------------------------------------------------------------
func AssignCharacteristicToInspectionResult( inspectionResultId uint64, characteristicId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionResult(inspectionResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionResult)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InspectionCharacteristic

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InspectionCharacteristic with a
		// matching characteristicId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, characteristicId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Characteristic	to the InspectionResult
			//----------------------------------------------------------------------------
			parentObj.Characteristic = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionResult
			//----------------------------------------------------------------------------
			return UpdateInspectionResult(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Characteristic", characteristicId )
			return utils.RequestResult{false, msg, "assignCharacteristic", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Characteristic on a InspectionResult
//----------------------------------------------------------------------------
func UnassignCharacteristicFromInspectionResult(inspectionResultId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionResult(inspectionResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionResult)

		//----------------------------------------------------------------------------
		// assign an empty InspectionCharacteristic to the Characteristic
		//----------------------------------------------------------------------------
		parentObj.Characteristic = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Characteristic
		//----------------------------------------------------------------------------
		parentObj.CharacteristicId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionResult
		//----------------------------------------------------------------------------
		return UpdateInspectionResult(parentObj)

	} else {
		return parentRequestResult
	}

}


