import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Contact} from '../models/Contact';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {UserService} from '../services/User.service';
import {ActivityService} from '../services/Activity.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {CampaignService} from '../services/Campaign.service';
import {NoteService} from '../services/Note.service';
import {EmailMessageService} from '../services/EmailMessage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ContactService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	contact : Contact;

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
	// add a Contact
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod) : Observable<any> {
		const uri_ = this.apiUrl + '/Contact/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		title: title,
      		email: email,
      		phone: phone,
      		mobile: mobile,
      		mailingAddress: mailingAddress,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
			PreferredContactMethod: PreferredContactMethod
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Contact/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		title: title,
      		email: email,
      		phone: phone,
      		mobile: mobile,
      		mailingAddress: mailingAddress,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
			PreferredContactMethod: PreferredContactMethod
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteContact(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Contact/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Contact
	// returns the results untouched as an Observable Contact
	// Contact model
	// delegates via URI
	//********************************************************************
	getContact(id) : Observable<Contact> {
		const uri_ = this.apiUrl + '/Contact/load/' + id;

		return this.http.get<Contact>(uri_);
	}
	
	//********************************************************************
	// gets all Contact
	// returns the results untouched as JSON representation of an
	// Observable array of Contact models
	// delegates via URI
	//********************************************************************
	getContacts() : Observable<Contact[]> {
		const uri_ = this.apiUrl + '/Contact/';

		return this
			.http.get<Contact[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( contactId, _organizationId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.contact.organization = tmp;

	// save the Contact
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( contactId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// assign Organization to null
	this.contact.organization = null;

	// save the Contact
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( contactId, _accountId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.contact.account = tmp;

	// save the Contact
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( contactId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// assign Account to null
	this.contact.account = null;

	// save the Contact
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( contactId, _ownerId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.contact.owner = tmp;

	// save the Contact
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Contact
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( contactId ): Observable<any> {

		// get the Contact from storage
		this.loadHelper( contactId );

	// assign Owner to null
	this.contact.owner = null;

	// save the Contact
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( contactId, activitiesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.contact.activities.indexOf(activity) == -1 )
		this.contact.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( contactId, activitiesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.contact.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.contact.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more opportunitiesIds as a Opportunities
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOpportunities( contactId, opportunitiesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = opportunitiesIds.split(',')

	// iterate over array of opportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.contact.opportunities.indexOf(opportunity) == -1 )
		this.contact.opportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more opportunitiesIds as a Opportunities
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOpportunities( contactId, opportunitiesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= opportunitiesIds.split(',');
	var opportunities 	= this.contact.opportunities;

	if ( opportunities != null && opportunitiesIds != null ) {

		// iterate over array of opportunities ids
		opportunities.forEach(function (obj) {
			if ( opportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.contact.opportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more casesIds as a Cases
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCases( contactId, casesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = casesIds.split(',')

	// iterate over array of cases ids
	idList.forEach(function (id) {
		// read the Case_
		var case_ = new Case_Service(this.http).getCase_(id);
		// add the Case_ if not already assigned
		if ( this.contact.cases.indexOf(case_) == -1 )
		this.contact.cases.push(case_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more casesIds as a Cases
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCases( contactId, casesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= casesIds.split(',');
	var cases 	= this.contact.cases;

	if ( cases != null && casesIds != null ) {

		// iterate over array of cases ids
		cases.forEach(function (obj) {
			if ( casesIds.indexOf(obj._id) > -1 ) {
				// remove the Case_
				this.contact.cases.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( contactId, campaignsIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.contact.campaigns.indexOf(campaign) == -1 )
		this.contact.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( contactId, campaignsIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.contact.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.contact.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notesIds as a Notes
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotes( contactId, notesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = notesIds.split(',')

	// iterate over array of notes ids
	idList.forEach(function (id) {
		// read the Note
		var note = new NoteService(this.http).getNote(id);
		// add the Note if not already assigned
		if ( this.contact.notes.indexOf(note) == -1 )
		this.contact.notes.push(note);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notesIds as a Notes
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotes( contactId, notesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= notesIds.split(',');
	var notes 	= this.contact.notes;

	if ( notes != null && notesIds != null ) {

		// iterate over array of notes ids
		notes.forEach(function (obj) {
			if ( notesIds.indexOf(obj._id) > -1 ) {
				// remove the Note
				this.contact.notes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more emailMessagesIds as a EmailMessages
	// to a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmailMessages( contactId, emailMessagesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );

	// split on a comma with no spaces
	var idList = emailMessagesIds.split(',')

	// iterate over array of emailMessages ids
	idList.forEach(function (id) {
		// read the EmailMessage
		var emailMessage = new EmailMessageService(this.http).getEmailMessage(id);
		// add the EmailMessage if not already assigned
		if ( this.contact.emailMessages.indexOf(emailMessage) == -1 )
		this.contact.emailMessages.push(emailMessage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more emailMessagesIds as a EmailMessages
	// from a Contact
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmailMessages( contactId, emailMessagesIds ): Observable<any> {

		// get the Contact
		this.loadHelper( contactId );


	// split on a comma with no spaces
	var idList 					= emailMessagesIds.split(',');
	var emailMessages 	= this.contact.emailMessages;

	if ( emailMessages != null && emailMessagesIds != null ) {

		// iterate over array of emailMessages ids
		emailMessages.forEach(function (obj) {
			if ( emailMessagesIds.indexOf(obj._id) > -1 ) {
				// remove the EmailMessage
				this.contact.emailMessages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Contact
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Contact/update/' + this.contact;

	return  this.http.post(uri_, this.contact );
}

	//********************************************************************
	// loadHelper - internal helper to load a Contact
	//********************************************************************	
	loadHelper( id ) {
		this.getContact(id)
			.subscribe((res : Contact) => {
				this.contact = res;
			});
	}
}