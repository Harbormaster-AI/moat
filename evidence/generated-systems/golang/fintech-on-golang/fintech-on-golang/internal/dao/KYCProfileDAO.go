package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing KYCProfileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateKYCProfile - creates a new db entry
//----------------------------------------------------------------------------
func CreateKYCProfile(obj model.KYCProfile)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a KYCProfile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a KYCProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateKYCProfile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetKYCProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetKYCProfile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.KYCProfile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a KYCProfile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a KYCProfile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a KYCProfile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetKYCProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllKYCProfile - returns all
//----------------------------------------------------------------------------
func GetAllKYCProfile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.KYCProfile

	//----------------------------------------------------------------------------
	// Request the ORM to find all KYCProfile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all KYCProfile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all KYCProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllKYCProfile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateKYCProfile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateKYCProfile(obj model.KYCProfile)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a KYCProfile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a KYCProfile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateKYCProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteKYCProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteKYCProfile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetKYCProfile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.KYCProfile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a KYCProfile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a KYCProfile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteKYCProfile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a KYCProfile
//----------------------------------------------------------------------------
func AssignCustomerToKYCProfile( kYCProfileId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the KYCProfile
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the KYCProfile
			//----------------------------------------------------------------------------
			return UpdateKYCProfile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a KYCProfile
//----------------------------------------------------------------------------
func UnassignCustomerFromKYCProfile(kYCProfileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the KYCProfile
		//----------------------------------------------------------------------------
		return UpdateKYCProfile(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more documentsIds as a Documents to a KYCProfile
//----------------------------------------------------------------------------
func AddDocumentsToKYCProfile ( kYCProfileId uint64, documentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KYCDocument

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KYCDocument
			// with a matching documentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , documentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Documents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Documents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Documents", documentsId )
				return utils.RequestResult{false, msg, "unassignDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more documentsIds as a Documents from a KYCProfile
//----------------------------------------------------------------------------
func RemoveDocumentsFromKYCProfile( kYCProfileId uint64, documentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KYCDocument

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KYCDocument
			// with a matching documentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , documentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove KYCDocumentObj from the Documents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Documents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Documents", documentsId )
				return utils.RequestResult{false, msg, "removeDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more screeningsIds as a Screenings to a KYCProfile
//----------------------------------------------------------------------------
func AddScreeningsToKYCProfile ( kYCProfileId uint64, screeningsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Screening

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Screening
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Screenings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "unassignScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more screeningsIds as a Screenings from a KYCProfile
//----------------------------------------------------------------------------
func RemoveScreeningsFromKYCProfile( kYCProfileId uint64, screeningsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Screening

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Screening
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ScreeningObj from the Screenings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "removeScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more addressesIds as a Addresses to a KYCProfile
//----------------------------------------------------------------------------
func AddAddressesToKYCProfile ( kYCProfileId uint64, addressesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( addressesIds, ",")

		for _, addressesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.VerifiedAddress

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a VerifiedAddress
			// with a matching addressesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , addressesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Addresses using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Addresses").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Addresses", addressesId )
				return utils.RequestResult{false, msg, "unassignAddresses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more addressesIds as a Addresses from a KYCProfile
//----------------------------------------------------------------------------
func RemoveAddressesFromKYCProfile( kYCProfileId uint64, addressesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the KYCProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKYCProfile(kYCProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KYCProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KYCProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( addressesIds, ",")

		for _, addressesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.VerifiedAddress

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a VerifiedAddress
			// with a matching addressesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , addressesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove VerifiedAddressObj from the Addresses array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Addresses").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Addresses", addressesId )
				return utils.RequestResult{false, msg, "removeAddresses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified KYCProfile from the gorm
		//----------------------------------------------------------------------------
		return GetKYCProfile(kYCProfileId)

	} else {
		return parentRequestResult
	}
}

