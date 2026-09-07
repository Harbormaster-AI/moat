package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CandidateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCandidate - creates a new db entry
//----------------------------------------------------------------------------
func CreateCandidate(obj model.Candidate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Candidate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Candidate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCandidate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCandidate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCandidate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Candidate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Candidate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Candidate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Candidate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCandidate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCandidate - returns all
//----------------------------------------------------------------------------
func GetAllCandidate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Candidate

	//----------------------------------------------------------------------------
	// Request the ORM to find all Candidate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Candidate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Candidate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCandidate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCandidate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCandidate(obj model.Candidate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Candidate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Candidate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCandidate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCandidate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCandidate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCandidate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Candidate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Candidate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Candidate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCandidate", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more applicationsIds as a Applications to a Candidate
//----------------------------------------------------------------------------
func AddApplicationsToCandidate ( candidateId uint64, applicationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( applicationsIds, ",")

		for _, applicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobApplication

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobApplication
			// with a matching applicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , applicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Applications using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Applications").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Applications", applicationsId )
				return utils.RequestResult{false, msg, "unassignApplications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more applicationsIds as a Applications from a Candidate
//----------------------------------------------------------------------------
func RemoveApplicationsFromCandidate( candidateId uint64, applicationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( applicationsIds, ",")

		for _, applicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobApplication

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobApplication
			// with a matching applicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , applicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove JobApplicationObj from the Applications array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Applications").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Applications", applicationsId )
				return utils.RequestResult{false, msg, "removeApplications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more interviewsIds as a Interviews to a Candidate
//----------------------------------------------------------------------------
func AddInterviewsToCandidate ( candidateId uint64, interviewsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewsIds, ",")

		for _, interviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Interview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Interview
			// with a matching interviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Interviews using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviews").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviews", interviewsId )
				return utils.RequestResult{false, msg, "unassignInterviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more interviewsIds as a Interviews from a Candidate
//----------------------------------------------------------------------------
func RemoveInterviewsFromCandidate( candidateId uint64, interviewsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewsIds, ",")

		for _, interviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Interview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Interview
			// with a matching interviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InterviewObj from the Interviews array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviews").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviews", interviewsId )
				return utils.RequestResult{false, msg, "removeInterviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more offersIds as a Offers to a Candidate
//----------------------------------------------------------------------------
func AddOffersToCandidate ( candidateId uint64, offersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( offersIds, ",")

		for _, offersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Offer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Offer
			// with a matching offersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , offersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Offers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Offers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Offers", offersId )
				return utils.RequestResult{false, msg, "unassignOffers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more offersIds as a Offers from a Candidate
//----------------------------------------------------------------------------
func RemoveOffersFromCandidate( candidateId uint64, offersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( offersIds, ",")

		for _, offersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Offer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Offer
			// with a matching offersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , offersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OfferObj from the Offers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Offers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Offers", offersId )
				return utils.RequestResult{false, msg, "removeOffers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more documentsIds as a Documents to a Candidate
//----------------------------------------------------------------------------
func AddDocumentsToCandidate ( candidateId uint64, documentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Document

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Document
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
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more documentsIds as a Documents from a Candidate
//----------------------------------------------------------------------------
func RemoveDocumentsFromCandidate( candidateId uint64, documentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Candidate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCandidate(candidateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Candidate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Candidate)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Document

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Document
			// with a matching documentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , documentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DocumentObj from the Documents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Documents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Documents", documentsId )
				return utils.RequestResult{false, msg, "removeDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Candidate from the gorm
		//----------------------------------------------------------------------------
		return GetCandidate(candidateId)

	} else {
		return parentRequestResult
	}
}

