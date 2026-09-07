package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing Case_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCase_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateCase_(obj model.Case_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Case_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Case_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCase_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCase_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCase_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Case_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Case_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Case_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Case_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCase_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCase_ - returns all
//----------------------------------------------------------------------------
func GetAllCase_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Case_

	//----------------------------------------------------------------------------
	// Request the ORM to find all Case_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Case_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Case_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCase_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCase_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCase_(obj model.Case_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Case_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Case_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCase_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCase_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCase_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCase_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Case_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Case_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCase_", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Case_
//----------------------------------------------------------------------------
func AssignOrganizationToCase_( case_Id uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the Case_
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Case_
			//----------------------------------------------------------------------------
			return UpdateCase_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Case_
//----------------------------------------------------------------------------
func UnassignOrganizationFromCase_(case_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Case_
		//----------------------------------------------------------------------------
		return UpdateCase_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Case_
//----------------------------------------------------------------------------
func AssignAccountToCase_( case_Id uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the Case_
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Case_
			//----------------------------------------------------------------------------
			return UpdateCase_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Case_
//----------------------------------------------------------------------------
func UnassignAccountFromCase_(case_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Case_
		//----------------------------------------------------------------------------
		return UpdateCase_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contact on a Case_
//----------------------------------------------------------------------------
func AssignContactToCase_( case_Id uint64, contactId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Contact

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Contact with a
		// matching contactId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, contactId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Contact	to the Case_
			//----------------------------------------------------------------------------
			parentObj.Contact = &childObj

			//----------------------------------------------------------------------------
			// save the Case_
			//----------------------------------------------------------------------------
			return UpdateCase_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contact", contactId )
			return utils.RequestResult{false, msg, "assignContact", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contact on a Case_
//----------------------------------------------------------------------------
func UnassignContactFromCase_(case_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// assign an empty Contact to the Contact
		//----------------------------------------------------------------------------
		parentObj.Contact = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contact
		//----------------------------------------------------------------------------
		parentObj.ContactId = nil;

		//----------------------------------------------------------------------------
		// save the Case_
		//----------------------------------------------------------------------------
		return UpdateCase_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Case_
//----------------------------------------------------------------------------
func AssignOwnerToCase_( case_Id uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.User

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a User with a
		// matching ownerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ownerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Owner	to the Case_
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Case_
			//----------------------------------------------------------------------------
			return UpdateCase_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Case_
//----------------------------------------------------------------------------
func UnassignOwnerFromCase_(case_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Case_
		//----------------------------------------------------------------------------
		return UpdateCase_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Team on a Case_
//----------------------------------------------------------------------------
func AssignTeamToCase_( case_Id uint64, teamId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Team

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Team with a
		// matching teamId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, teamId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Team	to the Case_
			//----------------------------------------------------------------------------
			parentObj.Team = &childObj

			//----------------------------------------------------------------------------
			// save the Case_
			//----------------------------------------------------------------------------
			return UpdateCase_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Team", teamId )
			return utils.RequestResult{false, msg, "assignTeam", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Team on a Case_
//----------------------------------------------------------------------------
func UnassignTeamFromCase_(case_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		//----------------------------------------------------------------------------
		// assign an empty Team to the Team
		//----------------------------------------------------------------------------
		parentObj.Team = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Team
		//----------------------------------------------------------------------------
		parentObj.TeamId = nil;

		//----------------------------------------------------------------------------
		// save the Case_
		//----------------------------------------------------------------------------
		return UpdateCase_(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Case_
//----------------------------------------------------------------------------
func AddActivitiesToCase_ ( case_Id uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Activities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "unassignActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Case_
//----------------------------------------------------------------------------
func RemoveActivitiesFromCase_( case_Id uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ActivityObj from the Activities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "removeActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more caseCommentsIds as a CaseComments to a Case_
//----------------------------------------------------------------------------
func AddCaseCommentsToCase_ ( case_Id uint64, caseCommentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( caseCommentsIds, ",")

		for _, caseCommentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching caseCommentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , caseCommentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CaseComments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CaseComments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CaseComments", caseCommentsId )
				return utils.RequestResult{false, msg, "unassignCaseComments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more caseCommentsIds as a CaseComments from a Case_
//----------------------------------------------------------------------------
func RemoveCaseCommentsFromCase_( case_Id uint64, caseCommentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( caseCommentsIds, ",")

		for _, caseCommentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching caseCommentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , caseCommentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove NoteObj from the CaseComments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CaseComments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CaseComments", caseCommentsId )
				return utils.RequestResult{false, msg, "removeCaseComments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more emailsIds as a Emails to a Case_
//----------------------------------------------------------------------------
func AddEmailsToCase_ ( case_Id uint64, emailsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailsIds, ",")

		for _, emailsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Emails using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Emails").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Emails", emailsId )
				return utils.RequestResult{false, msg, "unassignEmails", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more emailsIds as a Emails from a Case_
//----------------------------------------------------------------------------
func RemoveEmailsFromCase_( case_Id uint64, emailsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailsIds, ",")

		for _, emailsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmailMessageObj from the Emails array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Emails").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Emails", emailsId )
				return utils.RequestResult{false, msg, "removeEmails", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more relatedOpportunitiesIds as a RelatedOpportunities to a Case_
//----------------------------------------------------------------------------
func AddRelatedOpportunitiesToCase_ ( case_Id uint64, relatedOpportunitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedOpportunitiesIds, ",")

		for _, relatedOpportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching relatedOpportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedOpportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedOpportunities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedOpportunities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedOpportunities", relatedOpportunitiesId )
				return utils.RequestResult{false, msg, "unassignRelatedOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedOpportunitiesIds as a RelatedOpportunities from a Case_
//----------------------------------------------------------------------------
func RemoveRelatedOpportunitiesFromCase_( case_Id uint64, relatedOpportunitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Case_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCase_(case_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Case_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Case_)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedOpportunitiesIds, ",")

		for _, relatedOpportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching relatedOpportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedOpportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityObj from the RelatedOpportunities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedOpportunities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedOpportunities", relatedOpportunitiesId )
				return utils.RequestResult{false, msg, "removeRelatedOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Case_ from the gorm
		//----------------------------------------------------------------------------
		return GetCase_(case_Id)

	} else {
		return parentRequestResult
	}
}

