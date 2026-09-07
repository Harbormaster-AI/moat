package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing KYCDocumentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateKYCDocument - creates a new db entry
//----------------------------------------------------------------------------
func CreateKYCDocument(obj model.KYCDocument)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a KYCDocument with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a KYCDocument", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateKYCDocument", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetKYCDocument - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetKYCDocument(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.KYCDocument

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a KYCDocument with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a KYCDocument using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a KYCDocument using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetKYCDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllKYCDocument - returns all
//----------------------------------------------------------------------------
func GetAllKYCDocument()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.KYCDocument

	//----------------------------------------------------------------------------
	// Request the ORM to find all KYCDocument
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all KYCDocument" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all KYCDocument", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllKYCDocument", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateKYCDocument - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateKYCDocument(obj model.KYCDocument)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a KYCDocument using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a KYCDocument using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateKYCDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteKYCDocument - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteKYCDocument(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the KYCDocument with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetKYCDocument(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.KYCDocument)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a KYCDocument using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a KYCDocument using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteKYCDocument", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a KYCDocument
//----------------------------------------------------------------------------
func AssignKycProfileToKYCDocument( kYCDocumentId uint64, kycProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the KYCDocument with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCDocument(kYCDocumentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCDocument)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.KYCProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KYCProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the KYCDocument
			//----------------------------------------------------------------------------
			parentObj.KycProfile = &childObj

			//----------------------------------------------------------------------------
			// save the KYCDocument
			//----------------------------------------------------------------------------
			return UpdateKYCDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )
			return utils.RequestResult{false, msg, "assignKycProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a KYCDocument
//----------------------------------------------------------------------------
func UnassignKycProfileFromKYCDocument(kYCDocumentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KYCDocument with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCDocument(kYCDocumentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCDocument so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCDocument)

		//----------------------------------------------------------------------------
		// assign an empty KYCProfile to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the KYCDocument
		//----------------------------------------------------------------------------
		return UpdateKYCDocument(parentObj)

	} else {
		return parentRequestResult
	}

}


