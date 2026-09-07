import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Lead} from '../models/Lead';
import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {ActivityService} from '../services/Activity.service';
import {CampaignService} from '../services/Campaign.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {OpportunityService} from '../services/Opportunity.service';
import {NoteService} from '../services/Note.service';
import {EmailMessageService} from '../services/EmailMessage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LeadService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	lead : Lead;

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
	// add a Lead
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating) : Observable<any> {
		const uri_ = this.apiUrl + '/Lead/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		company: company,
      		email: email,
      		phone: phone,
      		converted: converted,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		ConvertedAccount: ConvertedAccount != null && ConvertedAccount.length > 0 ? ConvertedAccount : null,
      		ConvertedContact: ConvertedContact != null && ConvertedContact.length > 0 ? ConvertedContact : null,
      		ConvertedOpportunity: ConvertedOpportunity != null && ConvertedOpportunity.length > 0 ? ConvertedOpportunity : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		Status: Status,
      		Source: Source,
			Rating: Rating
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Lead/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		company: company,
      		email: email,
      		phone: phone,
      		converted: converted,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		ConvertedAccount: ConvertedAccount != null && ConvertedAccount.length > 0 ? ConvertedAccount : null,
      		ConvertedContact: ConvertedContact != null && ConvertedContact.length > 0 ? ConvertedContact : null,
      		ConvertedOpportunity: ConvertedOpportunity != null && ConvertedOpportunity.length > 0 ? ConvertedOpportunity : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		Status: Status,
      		Source: Source,
			Rating: Rating
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLead(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Lead/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Lead
	// returns the results untouched as an Observable Lead
	// Lead model
	// delegates via URI
	//********************************************************************
	getLead(id) : Observable<Lead> {
		const uri_ = this.apiUrl + '/Lead/load/' + id;

		return this.http.get<Lead>(uri_);
	}
	
	//********************************************************************
	// gets all Lead
	// returns the results untouched as JSON representation of an
	// Observable array of Lead models
	// delegates via URI
	//********************************************************************
	getLeads() : Observable<Lead[]> {
		const uri_ = this.apiUrl + '/Lead/';

		return this
			.http.get<Lead[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( leadId, _organizationId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.lead.organization = tmp;

	// save the Lead
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( leadId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// assign Organization to null
	this.lead.organization = null;

	// save the Lead
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( leadId, _ownerId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.lead.owner = tmp;

	// save the Lead
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( leadId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// assign Owner to null
	this.lead.owner = null;

	// save the Lead
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ConvertedAccount on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConvertedAccount( leadId, _convertedAccountId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_convertedAccountId);

	// assign the ConvertedAccount
	this.lead.convertedAccount = tmp;

	// save the Lead
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConvertedAccount on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConvertedAccount( leadId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// assign ConvertedAccount to null
	this.lead.convertedAccount = null;

	// save the Lead
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ConvertedContact on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConvertedContact( leadId, _convertedContactId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_convertedContactId);

	// assign the ConvertedContact
	this.lead.convertedContact = tmp;

	// save the Lead
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConvertedContact on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConvertedContact( leadId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// assign ConvertedContact to null
	this.lead.convertedContact = null;

	// save the Lead
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ConvertedOpportunity on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConvertedOpportunity( leadId, _convertedOpportunityId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_convertedOpportunityId);

	// assign the ConvertedOpportunity
	this.lead.convertedOpportunity = tmp;

	// save the Lead
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConvertedOpportunity on a Lead
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConvertedOpportunity( leadId ): Observable<any> {

		// get the Lead from storage
		this.loadHelper( leadId );

	// assign ConvertedOpportunity to null
	this.lead.convertedOpportunity = null;

	// save the Lead
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( leadId, activitiesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.lead.activities.indexOf(activity) == -1 )
		this.lead.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( leadId, activitiesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.lead.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.lead.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( leadId, campaignsIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.lead.campaigns.indexOf(campaign) == -1 )
		this.lead.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( leadId, campaignsIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.lead.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.lead.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notesIds as a Notes
	// to a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotes( leadId, notesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );

	// split on a comma with no spaces
	var idList = notesIds.split(',')

	// iterate over array of notes ids
	idList.forEach(function (id) {
		// read the Note
		var note = new NoteService(this.http).getNote(id);
		// add the Note if not already assigned
		if ( this.lead.notes.indexOf(note) == -1 )
		this.lead.notes.push(note);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notesIds as a Notes
	// from a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotes( leadId, notesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );


	// split on a comma with no spaces
	var idList 					= notesIds.split(',');
	var notes 	= this.lead.notes;

	if ( notes != null && notesIds != null ) {

		// iterate over array of notes ids
		notes.forEach(function (obj) {
			if ( notesIds.indexOf(obj._id) > -1 ) {
				// remove the Note
				this.lead.notes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more emailMessagesIds as a EmailMessages
	// to a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmailMessages( leadId, emailMessagesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );

	// split on a comma with no spaces
	var idList = emailMessagesIds.split(',')

	// iterate over array of emailMessages ids
	idList.forEach(function (id) {
		// read the EmailMessage
		var emailMessage = new EmailMessageService(this.http).getEmailMessage(id);
		// add the EmailMessage if not already assigned
		if ( this.lead.emailMessages.indexOf(emailMessage) == -1 )
		this.lead.emailMessages.push(emailMessage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more emailMessagesIds as a EmailMessages
	// from a Lead
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmailMessages( leadId, emailMessagesIds ): Observable<any> {

		// get the Lead
		this.loadHelper( leadId );


	// split on a comma with no spaces
	var idList 					= emailMessagesIds.split(',');
	var emailMessages 	= this.lead.emailMessages;

	if ( emailMessages != null && emailMessagesIds != null ) {

		// iterate over array of emailMessages ids
		emailMessages.forEach(function (obj) {
			if ( emailMessagesIds.indexOf(obj._id) > -1 ) {
				// remove the EmailMessage
				this.lead.emailMessages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Lead
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Lead/update/' + this.lead;

	return  this.http.post(uri_, this.lead );
}

	//********************************************************************
	// loadHelper - internal helper to load a Lead
	//********************************************************************	
	loadHelper( id ) {
		this.getLead(id)
			.subscribe((res : Lead) => {
				this.lead = res;
			});
	}
}