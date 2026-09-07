import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Activity} from '../models/Activity';
import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {LeadService} from '../services/Lead.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ActivityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	activity : Activity;

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
	// add a Activity
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/Activity/create';
		const obj = {
			      		subject: subject,
      		dueDate: dueDate,
      		startAt: startAt,
      		endAt: endAt,
      		location: location,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		ActivityType: ActivityType,
      		Status: Status,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Activity/update/' + id;
		const obj = {
				      		subject: subject,
      		dueDate: dueDate,
      		startAt: startAt,
      		endAt: endAt,
      		location: location,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		ActivityType: ActivityType,
      		Status: Status,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteActivity(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Activity/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Activity
	// returns the results untouched as an Observable Activity
	// Activity model
	// delegates via URI
	//********************************************************************
	getActivity(id) : Observable<Activity> {
		const uri_ = this.apiUrl + '/Activity/load/' + id;

		return this.http.get<Activity>(uri_);
	}
	
	//********************************************************************
	// gets all Activity
	// returns the results untouched as JSON representation of an
	// Observable array of Activity models
	// delegates via URI
	//********************************************************************
	getActivitys() : Observable<Activity[]> {
		const uri_ = this.apiUrl + '/Activity/';

		return this
			.http.get<Activity[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( activityId, _organizationId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.activity.organization = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Organization to null
	this.activity.organization = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( activityId, _ownerId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.activity.owner = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Owner to null
	this.activity.owner = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( activityId, _accountId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.activity.account = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Account to null
	this.activity.account = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contact on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContact( activityId, _contactId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_contactId);

	// assign the Contact
	this.activity.contact = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contact on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContact( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Contact to null
	this.activity.contact = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lead on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLead( activityId, _leadId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Lead from storage
	var tmp 	= new LeadService(this.http).getLead(_leadId);

	// assign the Lead
	this.activity.lead = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lead on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLead( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Lead to null
	this.activity.lead = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Opportunity on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( activityId, _opportunityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.activity.opportunity = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Opportunity to null
	this.activity.opportunity = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Case on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCase( activityId, _caseId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Case_ from storage
	var tmp 	= new Case_Service(this.http).getCase_(_caseId);

	// assign the Case
	this.activity.case = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Case on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCase( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Case to null
	this.activity.case = null;

	// save the Activity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Campaign on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( activityId, _campaignId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.activity.campaign = tmp;

	// save the Activity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a Activity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( activityId ): Observable<any> {

		// get the Activity from storage
		this.loadHelper( activityId );

	// assign Campaign to null
	this.activity.campaign = null;

	// save the Activity
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Activity
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Activity/update/' + this.activity;

	return  this.http.post(uri_, this.activity );
}

	//********************************************************************
	// loadHelper - internal helper to load a Activity
	//********************************************************************	
	loadHelper( id ) {
		this.getActivity(id)
			.subscribe((res : Activity) => {
				this.activity = res;
			});
	}
}