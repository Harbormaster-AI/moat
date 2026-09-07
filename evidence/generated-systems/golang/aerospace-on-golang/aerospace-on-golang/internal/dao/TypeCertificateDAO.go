package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TypeCertificateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTypeCertificate - creates a new db entry
//----------------------------------------------------------------------------
func CreateTypeCertificate(obj model.TypeCertificate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TypeCertificate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TypeCertificate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTypeCertificate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTypeCertificate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTypeCertificate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TypeCertificate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TypeCertificate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TypeCertificate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TypeCertificate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTypeCertificate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTypeCertificate - returns all
//----------------------------------------------------------------------------
func GetAllTypeCertificate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TypeCertificate

	//----------------------------------------------------------------------------
	// Request the ORM to find all TypeCertificate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TypeCertificate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TypeCertificate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTypeCertificate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTypeCertificate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTypeCertificate(obj model.TypeCertificate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TypeCertificate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TypeCertificate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTypeCertificate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTypeCertificate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTypeCertificate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TypeCertificate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTypeCertificate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TypeCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TypeCertificate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TypeCertificate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TypeCertificate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTypeCertificate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Program on a TypeCertificate
//----------------------------------------------------------------------------
func AssignProgramToTypeCertificate( typeCertificateId uint64, programId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TypeCertificate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTypeCertificate(typeCertificateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TypeCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TypeCertificate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftProgram

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftProgram with a
		// matching programId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, programId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Program	to the TypeCertificate
			//----------------------------------------------------------------------------
			parentObj.Program = &childObj

			//----------------------------------------------------------------------------
			// save the TypeCertificate
			//----------------------------------------------------------------------------
			return UpdateTypeCertificate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Program", programId )
			return utils.RequestResult{false, msg, "assignProgram", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Program on a TypeCertificate
//----------------------------------------------------------------------------
func UnassignProgramFromTypeCertificate(typeCertificateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TypeCertificate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTypeCertificate(typeCertificateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TypeCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TypeCertificate)

		//----------------------------------------------------------------------------
		// assign an empty AircraftProgram to the Program
		//----------------------------------------------------------------------------
		parentObj.Program = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Program
		//----------------------------------------------------------------------------
		parentObj.ProgramId = nil;

		//----------------------------------------------------------------------------
		// save the TypeCertificate
		//----------------------------------------------------------------------------
		return UpdateTypeCertificate(parentObj)

	} else {
		return parentRequestResult
	}

}


