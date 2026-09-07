import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Case_} from '../models/Case_';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {UserService} from '../services/User.service';
import {TeamService} from '../services/Team.service';
import {ActivityService} from '../services/Activity.service';
import {NoteService} from '../services/Note.service';
import {EmailMessageService} from '../services/EmailMessage.service';
import {OpportunityService} from '../services/Opportunity.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class Case_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	case_ : Case_;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a Case_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity) : Observable<any> {
		const uri_ = this.apiUrl + '/Case_/create';
		const obj = {
			      		caseNumber: caseNumber,
      		subject: subject,
      		description: description,
      		slaDue: slaDue,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Team: Team != null && Team.length > 0 ? Team : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		CaseComments: CaseComments != null && CaseComments.length > 0 ? CaseComments : null,
      		Emails: Emails != null && Emails.length > 0 ? Emails : null,
      		RelatedOpportunities: RelatedOpportunities != null && RelatedOpportunities.length > 0 ? RelatedOpportunities : null,
      		Status: Status,
      		Priority: Priority,
      		Origin: Origin,
			Severity: Severity
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCase_(caseNumber, subject, description, slaDue, Organization, Account, Contact, Owner, Team, Activities, CaseComments, Emails, RelatedOpportunities, Status, Priority, Origin, Severity, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Case_/update/' + id;
		const obj = {
				      		caseNumber: caseNumber,
      		subject: subject,
      		description: description,
      		slaDue: slaDue,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Team: Team != null && Team.length > 0 ? Team : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		CaseComments: CaseComments != null && CaseComments.length > 0 ? CaseComments : null,
      		Emails: Emails != null && Emails.length > 0 ? Emails : null,
      		RelatedOpportunities: RelatedOpportunities != null && RelatedOpportunities.length > 0 ? RelatedOpportunities : null,
      		Status: Status,
      		Priority: Priority,
      		Origin: Origin,
			Severity: Severity
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCase_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Case_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Case_
	// returns the results untouched as an Observable Case_
	// Case_ model
	// delegates via URI
	//********************************************************************
	getCase_(id) : Observable<Case_> {
		const uri_ = this.apiUrl + '/Case_/load/' + id;

		return this.http.get<Case_>(uri_);
	}
	
	//********************************************************************
	// gets all Case_
	// returns the results untouched as JSON representation of an
	// Observable array of Case_ models
	// delegates via URI
	//********************************************************************
	getCase_s() : Observable<Case_[]> {
		const uri_ = this.apiUrl + '/Case_/';

		return this
			.http.get<Case_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( case_Id, _organizationId ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.case_.organization = tmp;

	// save the Case_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( case_Id ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// assign Organization to null
	this.case_.organization = null;

	// save the Case_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( case_Id, _accountId ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.case_.account = tmp;

	// save the Case_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( case_Id ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// assign Account to null
	this.case_.account = null;

	// save the Case_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contact on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContact( case_Id, _contactId ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_contactId);

	// assign the Contact
	this.case_.contact = tmp;

	// save the Case_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contact on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContact( case_Id ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// assign Contact to null
	this.case_.contact = null;

	// save the Case_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( case_Id, _ownerId ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.case_.owner = tmp;

	// save the Case_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( case_Id ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// assign Owner to null
	this.case_.owner = null;

	// save the Case_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Team on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTeam( case_Id, _teamId ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// get the Team from storage
	var tmp 	= new TeamService(this.http).getTeam(_teamId);

	// assign the Team
	this.case_.team = tmp;

	// save the Case_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Team on a Case_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTeam( case_Id ): Observable<any> {

		// get the Case_ from storage
		this.loadHelper( case_Id );

	// assign Team to null
	this.case_.team = null;

	// save the Case_
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( case_Id, activitiesIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.case_.activities.indexOf(activity) == -1 )
		this.case_.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( case_Id, activitiesIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.case_.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.case_.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more caseCommentsIds as a CaseComments
	// to a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCaseComments( case_Id, caseCommentsIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );

	// split on a comma with no spaces
	var idList = caseCommentsIds.split(',')

	// iterate over array of caseComments ids
	idList.forEach(function (id) {
		// read the Note
		var note = new NoteService(this.http).getNote(id);
		// add the Note if not already assigned
		if ( this.case_.caseComments.indexOf(note) == -1 )
		this.case_.caseComments.push(note);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more caseCommentsIds as a CaseComments
	// from a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCaseComments( case_Id, caseCommentsIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );


	// split on a comma with no spaces
	var idList 					= caseCommentsIds.split(',');
	var caseComments 	= this.case_.caseComments;

	if ( caseComments != null && caseCommentsIds != null ) {

		// iterate over array of caseComments ids
		caseComments.forEach(function (obj) {
			if ( caseCommentsIds.indexOf(obj._id) > -1 ) {
				// remove the Note
				this.case_.caseComments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more emailsIds as a Emails
	// to a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmails( case_Id, emailsIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );

	// split on a comma with no spaces
	var idList = emailsIds.split(',')

	// iterate over array of emails ids
	idList.forEach(function (id) {
		// read the EmailMessage
		var emailMessage = new EmailMessageService(this.http).getEmailMessage(id);
		// add the EmailMessage if not already assigned
		if ( this.case_.emails.indexOf(emailMessage) == -1 )
		this.case_.emails.push(emailMessage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more emailsIds as a Emails
	// from a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmails( case_Id, emailsIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );


	// split on a comma with no spaces
	var idList 					= emailsIds.split(',');
	var emails 	= this.case_.emails;

	if ( emails != null && emailsIds != null ) {

		// iterate over array of emails ids
		emails.forEach(function (obj) {
			if ( emailsIds.indexOf(obj._id) > -1 ) {
				// remove the EmailMessage
				this.case_.emails.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more relatedOpportunitiesIds as a RelatedOpportunities
	// to a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedOpportunities( case_Id, relatedOpportunitiesIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );

	// split on a comma with no spaces
	var idList = relatedOpportunitiesIds.split(',')

	// iterate over array of relatedOpportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.case_.relatedOpportunities.indexOf(opportunity) == -1 )
		this.case_.relatedOpportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedOpportunitiesIds as a RelatedOpportunities
	// from a Case_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedOpportunities( case_Id, relatedOpportunitiesIds ): Observable<any> {

		// get the Case_
		this.loadHelper( case_Id );


	// split on a comma with no spaces
	var idList 					= relatedOpportunitiesIds.split(',');
	var relatedOpportunities 	= this.case_.relatedOpportunities;

	if ( relatedOpportunities != null && relatedOpportunitiesIds != null ) {

		// iterate over array of relatedOpportunities ids
		relatedOpportunities.forEach(function (obj) {
			if ( relatedOpportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.case_.relatedOpportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Case_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Case_/update/' + this.case_;

	return  this.http.post(uri_, this.case_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a Case_
	//********************************************************************	
	loadHelper( id ) {
		this.getCase_(id)
			.subscribe((res : Case_) => {
				this.case_ = res;
			});
	}
}