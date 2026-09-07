import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Campaign} from '../models/Campaign';
import {OrganizationService} from '../services/Organization.service';
import {CampaignMemberService} from '../services/CampaignMember.service';
import {OpportunityService} from '../services/Opportunity.service';
import {AccountService} from '../services/Account.service';
import {LeadService} from '../services/Lead.service';
import {ContactService} from '../services/Contact.service';
import {TeamService} from '../services/Team.service';
import {ActivityService} from '../services/Activity.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CampaignService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	campaign : Campaign;

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
	// add a Campaign
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type) : Observable<any> {
		const uri_ = this.apiUrl + '/Campaign/create';
		const obj = {
			      		name: name,
      		startDate: startDate,
      		endDate: endDate,
      		budget: budget,
      		actualCost: actualCost,
      		expectedRevenue: expectedRevenue,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ParentCampaign: ParentCampaign != null && ParentCampaign.length > 0 ? ParentCampaign : null,
      		ChildCampaigns: ChildCampaigns != null && ChildCampaigns.length > 0 ? ChildCampaigns : null,
      		Members: Members != null && Members.length > 0 ? Members : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Leads: Leads != null && Leads.length > 0 ? Leads : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Status: Status,
			Type: Type
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Campaign/update/' + id;
		const obj = {
				      		name: name,
      		startDate: startDate,
      		endDate: endDate,
      		budget: budget,
      		actualCost: actualCost,
      		expectedRevenue: expectedRevenue,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ParentCampaign: ParentCampaign != null && ParentCampaign.length > 0 ? ParentCampaign : null,
      		ChildCampaigns: ChildCampaigns != null && ChildCampaigns.length > 0 ? ChildCampaigns : null,
      		Members: Members != null && Members.length > 0 ? Members : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Leads: Leads != null && Leads.length > 0 ? Leads : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Status: Status,
			Type: Type
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCampaign(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Campaign/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Campaign
	// returns the results untouched as an Observable Campaign
	// Campaign model
	// delegates via URI
	//********************************************************************
	getCampaign(id) : Observable<Campaign> {
		const uri_ = this.apiUrl + '/Campaign/load/' + id;

		return this.http.get<Campaign>(uri_);
	}
	
	//********************************************************************
	// gets all Campaign
	// returns the results untouched as JSON representation of an
	// Observable array of Campaign models
	// delegates via URI
	//********************************************************************
	getCampaigns() : Observable<Campaign[]> {
		const uri_ = this.apiUrl + '/Campaign/';

		return this
			.http.get<Campaign[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( campaignId, _organizationId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.campaign.organization = tmp;

	// save the Campaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( campaignId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// assign Organization to null
	this.campaign.organization = null;

	// save the Campaign
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ParentCampaign on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParentCampaign( campaignId, _parentCampaignId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_parentCampaignId);

	// assign the ParentCampaign
	this.campaign.parentCampaign = tmp;

	// save the Campaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ParentCampaign on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParentCampaign( campaignId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// assign ParentCampaign to null
	this.campaign.parentCampaign = null;

	// save the Campaign
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more childCampaignsIds as a ChildCampaigns
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChildCampaigns( campaignId, childCampaignsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = childCampaignsIds.split(',')

	// iterate over array of childCampaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.campaign.childCampaigns.indexOf(campaign) == -1 )
		this.campaign.childCampaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more childCampaignsIds as a ChildCampaigns
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChildCampaigns( campaignId, childCampaignsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= childCampaignsIds.split(',');
	var childCampaigns 	= this.campaign.childCampaigns;

	if ( childCampaigns != null && childCampaignsIds != null ) {

		// iterate over array of childCampaigns ids
		childCampaigns.forEach(function (obj) {
			if ( childCampaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.campaign.childCampaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more membersIds as a Members
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMembers( campaignId, membersIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = membersIds.split(',')

	// iterate over array of members ids
	idList.forEach(function (id) {
		// read the CampaignMember
		var campaignMember = new CampaignMemberService(this.http).getCampaignMember(id);
		// add the CampaignMember if not already assigned
		if ( this.campaign.members.indexOf(campaignMember) == -1 )
		this.campaign.members.push(campaignMember);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more membersIds as a Members
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMembers( campaignId, membersIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= membersIds.split(',');
	var members 	= this.campaign.members;

	if ( members != null && membersIds != null ) {

		// iterate over array of members ids
		members.forEach(function (obj) {
			if ( membersIds.indexOf(obj._id) > -1 ) {
				// remove the CampaignMember
				this.campaign.members.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more opportunitiesIds as a Opportunities
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOpportunities( campaignId, opportunitiesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = opportunitiesIds.split(',')

	// iterate over array of opportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.campaign.opportunities.indexOf(opportunity) == -1 )
		this.campaign.opportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more opportunitiesIds as a Opportunities
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOpportunities( campaignId, opportunitiesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= opportunitiesIds.split(',');
	var opportunities 	= this.campaign.opportunities;

	if ( opportunities != null && opportunitiesIds != null ) {

		// iterate over array of opportunities ids
		opportunities.forEach(function (obj) {
			if ( opportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.campaign.opportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( campaignId, accountsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.campaign.accounts.indexOf(account) == -1 )
		this.campaign.accounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( campaignId, accountsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.campaign.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.campaign.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more leadsIds as a Leads
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLeads( campaignId, leadsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = leadsIds.split(',')

	// iterate over array of leads ids
	idList.forEach(function (id) {
		// read the Lead
		var lead = new LeadService(this.http).getLead(id);
		// add the Lead if not already assigned
		if ( this.campaign.leads.indexOf(lead) == -1 )
		this.campaign.leads.push(lead);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more leadsIds as a Leads
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLeads( campaignId, leadsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= leadsIds.split(',');
	var leads 	= this.campaign.leads;

	if ( leads != null && leadsIds != null ) {

		// iterate over array of leads ids
		leads.forEach(function (obj) {
			if ( leadsIds.indexOf(obj._id) > -1 ) {
				// remove the Lead
				this.campaign.leads.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contactsIds as a Contacts
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContacts( campaignId, contactsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = contactsIds.split(',')

	// iterate over array of contacts ids
	idList.forEach(function (id) {
		// read the Contact
		var contact = new ContactService(this.http).getContact(id);
		// add the Contact if not already assigned
		if ( this.campaign.contacts.indexOf(contact) == -1 )
		this.campaign.contacts.push(contact);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contactsIds as a Contacts
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContacts( campaignId, contactsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= contactsIds.split(',');
	var contacts 	= this.campaign.contacts;

	if ( contacts != null && contactsIds != null ) {

		// iterate over array of contacts ids
		contacts.forEach(function (obj) {
			if ( contactsIds.indexOf(obj._id) > -1 ) {
				// remove the Contact
				this.campaign.contacts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more teamsIds as a Teams
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTeams( campaignId, teamsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = teamsIds.split(',')

	// iterate over array of teams ids
	idList.forEach(function (id) {
		// read the Team
		var team = new TeamService(this.http).getTeam(id);
		// add the Team if not already assigned
		if ( this.campaign.teams.indexOf(team) == -1 )
		this.campaign.teams.push(team);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more teamsIds as a Teams
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTeams( campaignId, teamsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= teamsIds.split(',');
	var teams 	= this.campaign.teams;

	if ( teams != null && teamsIds != null ) {

		// iterate over array of teams ids
		teams.forEach(function (obj) {
			if ( teamsIds.indexOf(obj._id) > -1 ) {
				// remove the Team
				this.campaign.teams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( campaignId, activitiesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.campaign.activities.indexOf(activity) == -1 )
		this.campaign.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( campaignId, activitiesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.campaign.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.campaign.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Campaign
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Campaign/update/' + this.campaign;

	return  this.http.post(uri_, this.campaign );
}

	//********************************************************************
	// loadHelper - internal helper to load a Campaign
	//********************************************************************	
	loadHelper( id ) {
		this.getCampaign(id)
			.subscribe((res : Campaign) => {
				this.campaign = res;
			});
	}
}