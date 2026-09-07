package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CompetencyRatingDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCompetencyRating - creates a new db entry
//----------------------------------------------------------------------------
func CreateCompetencyRating(obj model.CompetencyRating)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CompetencyRating with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CompetencyRating", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCompetencyRating", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCompetencyRating - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCompetencyRating(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CompetencyRating

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CompetencyRating with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CompetencyRating using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CompetencyRating using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCompetencyRating", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCompetencyRating - returns all
//----------------------------------------------------------------------------
func GetAllCompetencyRating()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CompetencyRating

	//----------------------------------------------------------------------------
	// Request the ORM to find all CompetencyRating
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CompetencyRating" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CompetencyRating", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCompetencyRating", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCompetencyRating - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCompetencyRating(obj model.CompetencyRating)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CompetencyRating using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CompetencyRating using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCompetencyRating", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCompetencyRating - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCompetencyRating(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CompetencyRating with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCompetencyRating(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompetencyRating so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CompetencyRating)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CompetencyRating using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CompetencyRating using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCompetencyRating", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Review on a CompetencyRating
//----------------------------------------------------------------------------
func AssignReviewToCompetencyRating( competencyRatingId uint64, reviewId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CompetencyRating with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetencyRating(competencyRatingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompetencyRating so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompetencyRating)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PerformanceReview

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PerformanceReview with a
		// matching reviewId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, reviewId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Review	to the CompetencyRating
			//----------------------------------------------------------------------------
			parentObj.Review = &childObj

			//----------------------------------------------------------------------------
			// save the CompetencyRating
			//----------------------------------------------------------------------------
			return UpdateCompetencyRating(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Review", reviewId )
			return utils.RequestResult{false, msg, "assignReview", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Review on a CompetencyRating
//----------------------------------------------------------------------------
func UnassignReviewFromCompetencyRating(competencyRatingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompetencyRating with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetencyRating(competencyRatingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompetencyRating so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompetencyRating)

		//----------------------------------------------------------------------------
		// assign an empty PerformanceReview to the Review
		//----------------------------------------------------------------------------
		parentObj.Review = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Review
		//----------------------------------------------------------------------------
		parentObj.ReviewId = nil;

		//----------------------------------------------------------------------------
		// save the CompetencyRating
		//----------------------------------------------------------------------------
		return UpdateCompetencyRating(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Competency on a CompetencyRating
//----------------------------------------------------------------------------
func AssignCompetencyToCompetencyRating( competencyRatingId uint64, competencyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CompetencyRating with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetencyRating(competencyRatingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompetencyRating so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompetencyRating)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Competency

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Competency with a
		// matching competencyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, competencyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Competency	to the CompetencyRating
			//----------------------------------------------------------------------------
			parentObj.Competency = &childObj

			//----------------------------------------------------------------------------
			// save the CompetencyRating
			//----------------------------------------------------------------------------
			return UpdateCompetencyRating(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Competency", competencyId )
			return utils.RequestResult{false, msg, "assignCompetency", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Competency on a CompetencyRating
//----------------------------------------------------------------------------
func UnassignCompetencyFromCompetencyRating(competencyRatingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompetencyRating with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetencyRating(competencyRatingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompetencyRating so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompetencyRating)

		//----------------------------------------------------------------------------
		// assign an empty Competency to the Competency
		//----------------------------------------------------------------------------
		parentObj.Competency = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Competency
		//----------------------------------------------------------------------------
		parentObj.CompetencyId = nil;

		//----------------------------------------------------------------------------
		// save the CompetencyRating
		//----------------------------------------------------------------------------
		return UpdateCompetencyRating(parentObj)

	} else {
		return parentRequestResult
	}

}


