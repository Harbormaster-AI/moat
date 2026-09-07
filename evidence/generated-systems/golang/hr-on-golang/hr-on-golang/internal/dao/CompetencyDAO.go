package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CompetencyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCompetency - creates a new db entry
//----------------------------------------------------------------------------
func CreateCompetency(obj model.Competency)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Competency with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Competency", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCompetency", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCompetency - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCompetency(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Competency

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Competency with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Competency using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Competency using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCompetency", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCompetency - returns all
//----------------------------------------------------------------------------
func GetAllCompetency()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Competency

	//----------------------------------------------------------------------------
	// Request the ORM to find all Competency
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Competency" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Competency", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCompetency", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCompetency - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCompetency(obj model.Competency)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Competency using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Competency using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCompetency", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCompetency - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCompetency(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Competency with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCompetency(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Competency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Competency)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Competency using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Competency using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCompetency", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more jobProfilesIds as a JobProfiles to a Competency
//----------------------------------------------------------------------------
func AddJobProfilesToCompetency ( competencyId uint64, jobProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Competency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetency(competencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Competency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Competency)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobProfilesIds, ",")

		for _, jobProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobProfile
			// with a matching jobProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the JobProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfiles", jobProfilesId )
				return utils.RequestResult{false, msg, "unassignJobProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Competency from the gorm
		//----------------------------------------------------------------------------
		return GetCompetency(competencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more jobProfilesIds as a JobProfiles from a Competency
//----------------------------------------------------------------------------
func RemoveJobProfilesFromCompetency( competencyId uint64, jobProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Competency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetency(competencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Competency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Competency)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobProfilesIds, ",")

		for _, jobProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobProfile
			// with a matching jobProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove JobProfileObj from the JobProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfiles", jobProfilesId )
				return utils.RequestResult{false, msg, "removeJobProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Competency from the gorm
		//----------------------------------------------------------------------------
		return GetCompetency(competencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more competencyRatingsIds as a CompetencyRatings to a Competency
//----------------------------------------------------------------------------
func AddCompetencyRatingsToCompetency ( competencyId uint64, competencyRatingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Competency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetency(competencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Competency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Competency)

		// slice the ids on comma with no spaces
		ids := strings.Split( competencyRatingsIds, ",")

		for _, competencyRatingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompetencyRating

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompetencyRating
			// with a matching competencyRatingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competencyRatingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompetencyRatings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompetencyRatings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompetencyRatings", competencyRatingsId )
				return utils.RequestResult{false, msg, "unassignCompetencyRatings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Competency from the gorm
		//----------------------------------------------------------------------------
		return GetCompetency(competencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more competencyRatingsIds as a CompetencyRatings from a Competency
//----------------------------------------------------------------------------
func RemoveCompetencyRatingsFromCompetency( competencyId uint64, competencyRatingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Competency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompetency(competencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Competency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Competency)

		// slice the ids on comma with no spaces
		ids := strings.Split( competencyRatingsIds, ",")

		for _, competencyRatingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompetencyRating

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompetencyRating
			// with a matching competencyRatingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competencyRatingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CompetencyRatingObj from the CompetencyRatings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompetencyRatings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompetencyRatings", competencyRatingsId )
				return utils.RequestResult{false, msg, "removeCompetencyRatings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Competency from the gorm
		//----------------------------------------------------------------------------
		return GetCompetency(competencyId)

	} else {
		return parentRequestResult
	}
}

