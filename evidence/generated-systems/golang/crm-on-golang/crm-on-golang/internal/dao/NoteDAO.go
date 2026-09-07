package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing NoteDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateNote - creates a new db entry
//----------------------------------------------------------------------------
func CreateNote(obj model.Note)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Note with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Note", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateNote", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetNote - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetNote(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Note

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Note with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Note using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Note using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetNote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllNote - returns all
//----------------------------------------------------------------------------
func GetAllNote()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Note

	//----------------------------------------------------------------------------
	// Request the ORM to find all Note
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Note" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Note", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllNote", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateNote - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateNote(obj model.Note)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Note using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Note using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateNote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteNote - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteNote(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetNote(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Note using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Note using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteNote", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Note
//----------------------------------------------------------------------------
func AssignOrganizationToNote( noteId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

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
			// assign the Organization	to the Note
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Note
//----------------------------------------------------------------------------
func UnassignOrganizationFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Note
//----------------------------------------------------------------------------
func AssignOwnerToNote( noteId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

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
			// assign the Owner	to the Note
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Note
//----------------------------------------------------------------------------
func UnassignOwnerFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Note
//----------------------------------------------------------------------------
func AssignAccountToNote( noteId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

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
			// assign the Account	to the Note
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Note
//----------------------------------------------------------------------------
func UnassignAccountFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contact on a Note
//----------------------------------------------------------------------------
func AssignContactToNote( noteId uint64, contactId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

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
			// assign the Contact	to the Note
			//----------------------------------------------------------------------------
			parentObj.Contact = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contact", contactId )
			return utils.RequestResult{false, msg, "assignContact", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contact on a Note
//----------------------------------------------------------------------------
func UnassignContactFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Contact to the Contact
		//----------------------------------------------------------------------------
		parentObj.Contact = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contact
		//----------------------------------------------------------------------------
		parentObj.ContactId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Opportunity on a Note
//----------------------------------------------------------------------------
func AssignOpportunityToNote( noteId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Opportunity

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Opportunity with a
		// matching opportunityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, opportunityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Opportunity	to the Note
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a Note
//----------------------------------------------------------------------------
func UnassignOpportunityFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Case on a Note
//----------------------------------------------------------------------------
func AssignCaseToNote( noteId uint64, caseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Case_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Case_ with a
		// matching caseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, caseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Case	to the Note
			//----------------------------------------------------------------------------
			parentObj.Case = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Case", caseId )
			return utils.RequestResult{false, msg, "assignCase", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Case on a Note
//----------------------------------------------------------------------------
func UnassignCaseFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Case_ to the Case
		//----------------------------------------------------------------------------
		parentObj.Case = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Case
		//----------------------------------------------------------------------------
		parentObj.CaseId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lead on a Note
//----------------------------------------------------------------------------
func AssignLeadToNote( noteId uint64, leadId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Lead

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lead with a
		// matching leadId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, leadId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lead	to the Note
			//----------------------------------------------------------------------------
			parentObj.Lead = &childObj

			//----------------------------------------------------------------------------
			// save the Note
			//----------------------------------------------------------------------------
			return UpdateNote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lead", leadId )
			return utils.RequestResult{false, msg, "assignLead", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lead on a Note
//----------------------------------------------------------------------------
func UnassignLeadFromNote(noteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Note with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNote(noteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Note so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Note)

		//----------------------------------------------------------------------------
		// assign an empty Lead to the Lead
		//----------------------------------------------------------------------------
		parentObj.Lead = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lead
		//----------------------------------------------------------------------------
		parentObj.LeadId = nil;

		//----------------------------------------------------------------------------
		// save the Note
		//----------------------------------------------------------------------------
		return UpdateNote(parentObj)

	} else {
		return parentRequestResult
	}

}


