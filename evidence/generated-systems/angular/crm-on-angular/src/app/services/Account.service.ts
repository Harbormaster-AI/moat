import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Account} from '../models/Account';
import {OrganizationService} from '../services/Organization.service';
import {ContactService} from '../services/Contact.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {UserService} from '../services/User.service';
import {TerritoryService} from '../services/Territory.service';
import {ActivityService} from '../services/Activity.service';
import {CampaignService} from '../services/Campaign.service';
import {QuoteService} from '../services/Quote.service';
import {OrderService} from '../services/Order.service';
import {ContractService} from '../services/Contract.service';
import {NoteService} from '../services/Note.service';
import {EmailMessageService} from '../services/EmailMessage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	account : Account;

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
	// add a Account
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage) : Observable<any> {
		const uri_ = this.apiUrl + '/Account/create';
		const obj = {
			      		name: name,
      		accountNumber: accountNumber,
      		industry: industry,
      		billingAddress: billingAddress,
      		shippingAddress: shippingAddress,
      		website: website,
      		phone: phone,
      		asActive: asActive,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ParentAccount: ParentAccount != null && ParentAccount.length > 0 ? ParentAccount : null,
      		ChildAccounts: ChildAccounts != null && ChildAccounts.length > 0 ? ChildAccounts : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Territory: Territory != null && Territory.length > 0 ? Territory : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		AccountType: AccountType,
			LifecycleStage: LifecycleStage
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Account/update/' + id;
		const obj = {
				      		name: name,
      		accountNumber: accountNumber,
      		industry: industry,
      		billingAddress: billingAddress,
      		shippingAddress: shippingAddress,
      		website: website,
      		phone: phone,
      		asActive: asActive,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ParentAccount: ParentAccount != null && ParentAccount.length > 0 ? ParentAccount : null,
      		ChildAccounts: ChildAccounts != null && ChildAccounts.length > 0 ? ChildAccounts : null,
      		Contacts: Contacts != null && Contacts.length > 0 ? Contacts : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Territory: Territory != null && Territory.length > 0 ? Territory : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Notes: Notes != null && Notes.length > 0 ? Notes : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		AccountType: AccountType,
			LifecycleStage: LifecycleStage
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Account/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Account
	// returns the results untouched as an Observable Account
	// Account model
	// delegates via URI
	//********************************************************************
	getAccount(id) : Observable<Account> {
		const uri_ = this.apiUrl + '/Account/load/' + id;

		return this.http.get<Account>(uri_);
	}
	
	//********************************************************************
	// gets all Account
	// returns the results untouched as JSON representation of an
	// Observable array of Account models
	// delegates via URI
	//********************************************************************
	getAccounts() : Observable<Account[]> {
		const uri_ = this.apiUrl + '/Account/';

		return this
			.http.get<Account[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( accountId, _organizationId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.account.organization = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign Organization to null
	this.account.organization = null;

	// save the Account
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ParentAccount on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParentAccount( accountId, _parentAccountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_parentAccountId);

	// assign the ParentAccount
	this.account.parentAccount = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ParentAccount on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParentAccount( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign ParentAccount to null
	this.account.parentAccount = null;

	// save the Account
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( accountId, _ownerId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.account.owner = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign Owner to null
	this.account.owner = null;

	// save the Account
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Territory on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTerritory( accountId, _territoryId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the Territory from storage
	var tmp 	= new TerritoryService(this.http).getTerritory(_territoryId);

	// assign the Territory
	this.account.territory = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Territory on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTerritory( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign Territory to null
	this.account.territory = null;

	// save the Account
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more childAccountsIds as a ChildAccounts
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChildAccounts( accountId, childAccountsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = childAccountsIds.split(',')

	// iterate over array of childAccounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.account.childAccounts.indexOf(account) == -1 )
		this.account.childAccounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more childAccountsIds as a ChildAccounts
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChildAccounts( accountId, childAccountsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= childAccountsIds.split(',');
	var childAccounts 	= this.account.childAccounts;

	if ( childAccounts != null && childAccountsIds != null ) {

		// iterate over array of childAccounts ids
		childAccounts.forEach(function (obj) {
			if ( childAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.account.childAccounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contactsIds as a Contacts
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContacts( accountId, contactsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = contactsIds.split(',')

	// iterate over array of contacts ids
	idList.forEach(function (id) {
		// read the Contact
		var contact = new ContactService(this.http).getContact(id);
		// add the Contact if not already assigned
		if ( this.account.contacts.indexOf(contact) == -1 )
		this.account.contacts.push(contact);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contactsIds as a Contacts
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContacts( accountId, contactsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= contactsIds.split(',');
	var contacts 	= this.account.contacts;

	if ( contacts != null && contactsIds != null ) {

		// iterate over array of contacts ids
		contacts.forEach(function (obj) {
			if ( contactsIds.indexOf(obj._id) > -1 ) {
				// remove the Contact
				this.account.contacts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more opportunitiesIds as a Opportunities
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOpportunities( accountId, opportunitiesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = opportunitiesIds.split(',')

	// iterate over array of opportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.account.opportunities.indexOf(opportunity) == -1 )
		this.account.opportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more opportunitiesIds as a Opportunities
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOpportunities( accountId, opportunitiesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= opportunitiesIds.split(',');
	var opportunities 	= this.account.opportunities;

	if ( opportunities != null && opportunitiesIds != null ) {

		// iterate over array of opportunities ids
		opportunities.forEach(function (obj) {
			if ( opportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.account.opportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more casesIds as a Cases
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCases( accountId, casesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = casesIds.split(',')

	// iterate over array of cases ids
	idList.forEach(function (id) {
		// read the Case_
		var case_ = new Case_Service(this.http).getCase_(id);
		// add the Case_ if not already assigned
		if ( this.account.cases.indexOf(case_) == -1 )
		this.account.cases.push(case_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more casesIds as a Cases
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCases( accountId, casesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= casesIds.split(',');
	var cases 	= this.account.cases;

	if ( cases != null && casesIds != null ) {

		// iterate over array of cases ids
		cases.forEach(function (obj) {
			if ( casesIds.indexOf(obj._id) > -1 ) {
				// remove the Case_
				this.account.cases.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( accountId, activitiesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.account.activities.indexOf(activity) == -1 )
		this.account.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( accountId, activitiesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.account.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.account.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( accountId, campaignsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.account.campaigns.indexOf(campaign) == -1 )
		this.account.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( accountId, campaignsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.account.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.account.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( accountId, quotesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.account.quotes.indexOf(quote) == -1 )
		this.account.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( accountId, quotesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.account.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.account.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( accountId, ordersIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the Order
		var order = new OrderService(this.http).getOrder(id);
		// add the Order if not already assigned
		if ( this.account.orders.indexOf(order) == -1 )
		this.account.orders.push(order);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( accountId, ordersIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.account.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the Order
				this.account.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( accountId, contractsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the Contract
		var contract = new ContractService(this.http).getContract(id);
		// add the Contract if not already assigned
		if ( this.account.contracts.indexOf(contract) == -1 )
		this.account.contracts.push(contract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( accountId, contractsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.account.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the Contract
				this.account.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notesIds as a Notes
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotes( accountId, notesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = notesIds.split(',')

	// iterate over array of notes ids
	idList.forEach(function (id) {
		// read the Note
		var note = new NoteService(this.http).getNote(id);
		// add the Note if not already assigned
		if ( this.account.notes.indexOf(note) == -1 )
		this.account.notes.push(note);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notesIds as a Notes
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotes( accountId, notesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= notesIds.split(',');
	var notes 	= this.account.notes;

	if ( notes != null && notesIds != null ) {

		// iterate over array of notes ids
		notes.forEach(function (obj) {
			if ( notesIds.indexOf(obj._id) > -1 ) {
				// remove the Note
				this.account.notes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more emailMessagesIds as a EmailMessages
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmailMessages( accountId, emailMessagesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = emailMessagesIds.split(',')

	// iterate over array of emailMessages ids
	idList.forEach(function (id) {
		// read the EmailMessage
		var emailMessage = new EmailMessageService(this.http).getEmailMessage(id);
		// add the EmailMessage if not already assigned
		if ( this.account.emailMessages.indexOf(emailMessage) == -1 )
		this.account.emailMessages.push(emailMessage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more emailMessagesIds as a EmailMessages
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmailMessages( accountId, emailMessagesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= emailMessagesIds.split(',');
	var emailMessages 	= this.account.emailMessages;

	if ( emailMessages != null && emailMessagesIds != null ) {

		// iterate over array of emailMessages ids
		emailMessages.forEach(function (obj) {
			if ( emailMessagesIds.indexOf(obj._id) > -1 ) {
				// remove the EmailMessage
				this.account.emailMessages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Account
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Account/update/' + this.account;

	return  this.http.post(uri_, this.account );
}

	//********************************************************************
	// loadHelper - internal helper to load a Account
	//********************************************************************	
	loadHelper( id ) {
		this.getAccount(id)
			.subscribe((res : Account) => {
				this.account = res;
			});
	}
}