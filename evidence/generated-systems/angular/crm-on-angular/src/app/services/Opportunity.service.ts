import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Opportunity} from '../models/Opportunity';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {UserService} from '../services/User.service';
import {ContactService} from '../services/Contact.service';
import {OpportunityLineItemService} from '../services/OpportunityLineItem.service';
import {OpportunityStageHistoryService} from '../services/OpportunityStageHistory.service';
import {QuoteService} from '../services/Quote.service';
import {OrderService} from '../services/Order.service';
import {CampaignService} from '../services/Campaign.service';
import {ActivityService} from '../services/Activity.service';
import {TeamService} from '../services/Team.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OpportunityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	opportunity : Opportunity;

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
	// add a Opportunity
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory) : Observable<any> {
		const uri_ = this.apiUrl + '/Opportunity/create';
		const obj = {
			      		name: name,
      		amount: amount,
      		closeDate: closeDate,
      		probability: probability,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		StageHistory: StageHistory != null && StageHistory.length > 0 ? StageHistory : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Stage: Stage,
      		Type: Type,
			ForecastCategory: ForecastCategory
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Opportunity/update/' + id;
		const obj = {
				      		name: name,
      		amount: amount,
      		closeDate: closeDate,
      		probability: probability,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		StageHistory: StageHistory != null && StageHistory.length > 0 ? StageHistory : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Stage: Stage,
      		Type: Type,
			ForecastCategory: ForecastCategory
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOpportunity(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Opportunity/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Opportunity
	// returns the results untouched as an Observable Opportunity
	// Opportunity model
	// delegates via URI
	//********************************************************************
	getOpportunity(id) : Observable<Opportunity> {
		const uri_ = this.apiUrl + '/Opportunity/load/' + id;

		return this.http.get<Opportunity>(uri_);
	}
	
	//********************************************************************
	// gets all Opportunity
	// returns the results untouched as JSON representation of an
	// Observable array of Opportunity models
	// delegates via URI
	//********************************************************************
	getOpportunitys() : Observable<Opportunity[]> {
		const uri_ = this.apiUrl + '/Opportunity/';

		return this
			.http.get<Opportunity[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( opportunityId, _organizationId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.opportunity.organization = tmp;

	// save the Opportunity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( opportunityId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// assign Organization to null
	this.opportunity.organization = null;

	// save the Opportunity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( opportunityId, _accountId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.opportunity.account = tmp;

	// save the Opportunity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( opportunityId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// assign Account to null
	this.opportunity.account = null;

	// save the Opportunity
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( opportunityId, _ownerId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.opportunity.owner = tmp;

	// save the Opportunity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Opportunity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( opportunityId ): Observable<any> {

		// get the Opportunity from storage
		this.loadHelper( opportunityId );

	// assign Owner to null
	this.opportunity.owner = null;

	// save the Opportunity
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more contactsIds as a Contacts
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContacts( opportunityId, contactsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = contactsIds.split(',')

	// iterate over array of contacts ids
	idList.forEach(function (id) {
		// read the Contact
		var contact = new ContactService(this.http).getContact(id);
		// add the Contact if not already assigned
		if ( this.opportunity.contacts.indexOf(contact) == -1 )
		this.opportunity.contacts.push(contact);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contactsIds as a Contacts
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContacts( opportunityId, contactsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= contactsIds.split(',');
	var contacts 	= this.opportunity.contacts;

	if ( contacts != null && contactsIds != null ) {

		// iterate over array of contacts ids
		contacts.forEach(function (obj) {
			if ( contactsIds.indexOf(obj._id) > -1 ) {
				// remove the Contact
				this.opportunity.contacts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more lineItemsIds as a LineItems
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLineItems( opportunityId, lineItemsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = lineItemsIds.split(',')

	// iterate over array of lineItems ids
	idList.forEach(function (id) {
		// read the OpportunityLineItem
		var opportunityLineItem = new OpportunityLineItemService(this.http).getOpportunityLineItem(id);
		// add the OpportunityLineItem if not already assigned
		if ( this.opportunity.lineItems.indexOf(opportunityLineItem) == -1 )
		this.opportunity.lineItems.push(opportunityLineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lineItemsIds as a LineItems
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLineItems( opportunityId, lineItemsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= lineItemsIds.split(',');
	var lineItems 	= this.opportunity.lineItems;

	if ( lineItems != null && lineItemsIds != null ) {

		// iterate over array of lineItems ids
		lineItems.forEach(function (obj) {
			if ( lineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the OpportunityLineItem
				this.opportunity.lineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more stageHistoryIds as a StageHistory
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addStageHistory( opportunityId, stageHistoryIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = stageHistoryIds.split(',')

	// iterate over array of stageHistory ids
	idList.forEach(function (id) {
		// read the OpportunityStageHistory
		var opportunityStageHistory = new OpportunityStageHistoryService(this.http).getOpportunityStageHistory(id);
		// add the OpportunityStageHistory if not already assigned
		if ( this.opportunity.stageHistory.indexOf(opportunityStageHistory) == -1 )
		this.opportunity.stageHistory.push(opportunityStageHistory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more stageHistoryIds as a StageHistory
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeStageHistory( opportunityId, stageHistoryIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= stageHistoryIds.split(',');
	var stageHistory 	= this.opportunity.stageHistory;

	if ( stageHistory != null && stageHistoryIds != null ) {

		// iterate over array of stageHistory ids
		stageHistory.forEach(function (obj) {
			if ( stageHistoryIds.indexOf(obj._id) > -1 ) {
				// remove the OpportunityStageHistory
				this.opportunity.stageHistory.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( opportunityId, quotesIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.opportunity.quotes.indexOf(quote) == -1 )
		this.opportunity.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( opportunityId, quotesIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.opportunity.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.opportunity.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( opportunityId, ordersIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the Order
		var order = new OrderService(this.http).getOrder(id);
		// add the Order if not already assigned
		if ( this.opportunity.orders.indexOf(order) == -1 )
		this.opportunity.orders.push(order);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( opportunityId, ordersIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.opportunity.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the Order
				this.opportunity.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( opportunityId, campaignsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.opportunity.campaigns.indexOf(campaign) == -1 )
		this.opportunity.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( opportunityId, campaignsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.opportunity.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.opportunity.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( opportunityId, activitiesIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.opportunity.activities.indexOf(activity) == -1 )
		this.opportunity.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( opportunityId, activitiesIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.opportunity.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.opportunity.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more teamsIds as a Teams
	// to a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTeams( opportunityId, teamsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );

	// split on a comma with no spaces
	var idList = teamsIds.split(',')

	// iterate over array of teams ids
	idList.forEach(function (id) {
		// read the Team
		var team = new TeamService(this.http).getTeam(id);
		// add the Team if not already assigned
		if ( this.opportunity.teams.indexOf(team) == -1 )
		this.opportunity.teams.push(team);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more teamsIds as a Teams
	// from a Opportunity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTeams( opportunityId, teamsIds ): Observable<any> {

		// get the Opportunity
		this.loadHelper( opportunityId );


	// split on a comma with no spaces
	var idList 					= teamsIds.split(',');
	var teams 	= this.opportunity.teams;

	if ( teams != null && teamsIds != null ) {

		// iterate over array of teams ids
		teams.forEach(function (obj) {
			if ( teamsIds.indexOf(obj._id) > -1 ) {
				// remove the Team
				this.opportunity.teams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Opportunity
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Opportunity/update/' + this.opportunity;

	return  this.http.post(uri_, this.opportunity );
}

	//********************************************************************
	// loadHelper - internal helper to load a Opportunity
	//********************************************************************	
	loadHelper( id ) {
		this.getOpportunity(id)
			.subscribe((res : Opportunity) => {
				this.opportunity = res;
			});
	}
}