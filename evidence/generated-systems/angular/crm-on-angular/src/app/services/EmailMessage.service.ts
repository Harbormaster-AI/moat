import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EmailMessage} from '../models/EmailMessage';
import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {LeadService} from '../services/Lead.service';
import {Case_Service} from '../services/Case_.service';
import {OpportunityService} from '../services/Opportunity.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EmailMessageService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	emailMessage : EmailMessage;

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
	// add a EmailMessage
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/EmailMessage/create';
		const obj = {
			      		subject: subject,
      		body: body,
      		sentAt: sentAt,
      		messageId: messageId,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Direction: Direction,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EmailMessage/update/' + id;
		const obj = {
				      		subject: subject,
      		body: body,
      		sentAt: sentAt,
      		messageId: messageId,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Direction: Direction,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEmailMessage(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EmailMessage/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EmailMessage
	// returns the results untouched as an Observable EmailMessage
	// EmailMessage model
	// delegates via URI
	//********************************************************************
	getEmailMessage(id) : Observable<EmailMessage> {
		const uri_ = this.apiUrl + '/EmailMessage/load/' + id;

		return this.http.get<EmailMessage>(uri_);
	}
	
	//********************************************************************
	// gets all EmailMessage
	// returns the results untouched as JSON representation of an
	// Observable array of EmailMessage models
	// delegates via URI
	//********************************************************************
	getEmailMessages() : Observable<EmailMessage[]> {
		const uri_ = this.apiUrl + '/EmailMessage/';

		return this
			.http.get<EmailMessage[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( emailMessageId, _organizationId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.emailMessage.organization = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Organization to null
	this.emailMessage.organization = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( emailMessageId, _ownerId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.emailMessage.owner = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Owner to null
	this.emailMessage.owner = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( emailMessageId, _accountId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.emailMessage.account = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Account to null
	this.emailMessage.account = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contact on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContact( emailMessageId, _contactId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_contactId);

	// assign the Contact
	this.emailMessage.contact = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contact on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContact( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Contact to null
	this.emailMessage.contact = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lead on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLead( emailMessageId, _leadId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Lead from storage
	var tmp 	= new LeadService(this.http).getLead(_leadId);

	// assign the Lead
	this.emailMessage.lead = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lead on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLead( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Lead to null
	this.emailMessage.lead = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Case on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCase( emailMessageId, _caseId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Case_ from storage
	var tmp 	= new Case_Service(this.http).getCase_(_caseId);

	// assign the Case
	this.emailMessage.case = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Case on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCase( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Case to null
	this.emailMessage.case = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Opportunity on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( emailMessageId, _opportunityId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.emailMessage.opportunity = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Opportunity to null
	this.emailMessage.opportunity = null;

	// save the EmailMessage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Campaign on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( emailMessageId, _campaignId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.emailMessage.campaign = tmp;

	// save the EmailMessage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a EmailMessage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( emailMessageId ): Observable<any> {

		// get the EmailMessage from storage
		this.loadHelper( emailMessageId );

	// assign Campaign to null
	this.emailMessage.campaign = null;

	// save the EmailMessage
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a EmailMessage
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EmailMessage/update/' + this.emailMessage;

	return  this.http.post(uri_, this.emailMessage );
}

	//********************************************************************
	// loadHelper - internal helper to load a EmailMessage
	//********************************************************************	
	loadHelper( id ) {
		this.getEmailMessage(id)
			.subscribe((res : EmailMessage) => {
				this.emailMessage = res;
			});
	}
}