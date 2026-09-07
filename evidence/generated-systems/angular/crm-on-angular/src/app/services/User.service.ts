import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {User} from '../models/User';
import {OrganizationService} from '../services/Organization.service';
import {TeamService} from '../services/Team.service';
import {ActivityService} from '../services/Activity.service';
import {AccountService} from '../services/Account.service';
import {LeadService} from '../services/Lead.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {QuoteService} from '../services/Quote.service';
import {OrderService} from '../services/Order.service';
import {ContractService} from '../services/Contract.service';
import {EmailMessageService} from '../services/EmailMessage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class UserService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	user : User;

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
	// add a User
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addUser(username, fullName, email, locale, Organization, Teams, Activities, OwnedAccounts, OwnedLeads, OwnedOpportunities, OwnedCases, Quotes, Orders, Contracts, EmailMessages, Role, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/User/create';
		const obj = {
			      		username: username,
      		fullName: fullName,
      		email: email,
      		locale: locale,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		OwnedAccounts: OwnedAccounts != null && OwnedAccounts.length > 0 ? OwnedAccounts : null,
      		OwnedLeads: OwnedLeads != null && OwnedLeads.length > 0 ? OwnedLeads : null,
      		OwnedOpportunities: OwnedOpportunities != null && OwnedOpportunities.length > 0 ? OwnedOpportunities : null,
      		OwnedCases: OwnedCases != null && OwnedCases.length > 0 ? OwnedCases : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		Role: Role,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUser(username, fullName, email, locale, Organization, Teams, Activities, OwnedAccounts, OwnedLeads, OwnedOpportunities, OwnedCases, Quotes, Orders, Contracts, EmailMessages, Role, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/User/update/' + id;
		const obj = {
				      		username: username,
      		fullName: fullName,
      		email: email,
      		locale: locale,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Activities: Activities != null && Activities.length > 0 ? Activities : null,
      		OwnedAccounts: OwnedAccounts != null && OwnedAccounts.length > 0 ? OwnedAccounts : null,
      		OwnedLeads: OwnedLeads != null && OwnedLeads.length > 0 ? OwnedLeads : null,
      		OwnedOpportunities: OwnedOpportunities != null && OwnedOpportunities.length > 0 ? OwnedOpportunities : null,
      		OwnedCases: OwnedCases != null && OwnedCases.length > 0 ? OwnedCases : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		EmailMessages: EmailMessages != null && EmailMessages.length > 0 ? EmailMessages : null,
      		Role: Role,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteUser(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/User/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a User
	// returns the results untouched as an Observable User
	// User model
	// delegates via URI
	//********************************************************************
	getUser(id) : Observable<User> {
		const uri_ = this.apiUrl + '/User/load/' + id;

		return this.http.get<User>(uri_);
	}
	
	//********************************************************************
	// gets all User
	// returns the results untouched as JSON representation of an
	// Observable array of User models
	// delegates via URI
	//********************************************************************
	getUsers() : Observable<User[]> {
		const uri_ = this.apiUrl + '/User/';

		return this
			.http.get<User[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( userId, _organizationId ): Observable<any> {

		// get the User from storage
		this.loadHelper( userId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.user.organization = tmp;

	// save the User
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( userId ): Observable<any> {

		// get the User from storage
		this.loadHelper( userId );

	// assign Organization to null
	this.user.organization = null;

	// save the User
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more teamsIds as a Teams
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTeams( userId, teamsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = teamsIds.split(',')

	// iterate over array of teams ids
	idList.forEach(function (id) {
		// read the Team
		var team = new TeamService(this.http).getTeam(id);
		// add the Team if not already assigned
		if ( this.user.teams.indexOf(team) == -1 )
		this.user.teams.push(team);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more teamsIds as a Teams
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTeams( userId, teamsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= teamsIds.split(',');
	var teams 	= this.user.teams;

	if ( teams != null && teamsIds != null ) {

		// iterate over array of teams ids
		teams.forEach(function (obj) {
			if ( teamsIds.indexOf(obj._id) > -1 ) {
				// remove the Team
				this.user.teams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more activitiesIds as a Activities
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addActivities( userId, activitiesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = activitiesIds.split(',')

	// iterate over array of activities ids
	idList.forEach(function (id) {
		// read the Activity
		var activity = new ActivityService(this.http).getActivity(id);
		// add the Activity if not already assigned
		if ( this.user.activities.indexOf(activity) == -1 )
		this.user.activities.push(activity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more activitiesIds as a Activities
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeActivities( userId, activitiesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= activitiesIds.split(',');
	var activities 	= this.user.activities;

	if ( activities != null && activitiesIds != null ) {

		// iterate over array of activities ids
		activities.forEach(function (obj) {
			if ( activitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Activity
				this.user.activities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ownedAccountsIds as a OwnedAccounts
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwnedAccounts( userId, ownedAccountsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = ownedAccountsIds.split(',')

	// iterate over array of ownedAccounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.user.ownedAccounts.indexOf(account) == -1 )
		this.user.ownedAccounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownedAccountsIds as a OwnedAccounts
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwnedAccounts( userId, ownedAccountsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= ownedAccountsIds.split(',');
	var ownedAccounts 	= this.user.ownedAccounts;

	if ( ownedAccounts != null && ownedAccountsIds != null ) {

		// iterate over array of ownedAccounts ids
		ownedAccounts.forEach(function (obj) {
			if ( ownedAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.user.ownedAccounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ownedLeadsIds as a OwnedLeads
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwnedLeads( userId, ownedLeadsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = ownedLeadsIds.split(',')

	// iterate over array of ownedLeads ids
	idList.forEach(function (id) {
		// read the Lead
		var lead = new LeadService(this.http).getLead(id);
		// add the Lead if not already assigned
		if ( this.user.ownedLeads.indexOf(lead) == -1 )
		this.user.ownedLeads.push(lead);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownedLeadsIds as a OwnedLeads
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwnedLeads( userId, ownedLeadsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= ownedLeadsIds.split(',');
	var ownedLeads 	= this.user.ownedLeads;

	if ( ownedLeads != null && ownedLeadsIds != null ) {

		// iterate over array of ownedLeads ids
		ownedLeads.forEach(function (obj) {
			if ( ownedLeadsIds.indexOf(obj._id) > -1 ) {
				// remove the Lead
				this.user.ownedLeads.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ownedOpportunitiesIds as a OwnedOpportunities
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwnedOpportunities( userId, ownedOpportunitiesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = ownedOpportunitiesIds.split(',')

	// iterate over array of ownedOpportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.user.ownedOpportunities.indexOf(opportunity) == -1 )
		this.user.ownedOpportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownedOpportunitiesIds as a OwnedOpportunities
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwnedOpportunities( userId, ownedOpportunitiesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= ownedOpportunitiesIds.split(',');
	var ownedOpportunities 	= this.user.ownedOpportunities;

	if ( ownedOpportunities != null && ownedOpportunitiesIds != null ) {

		// iterate over array of ownedOpportunities ids
		ownedOpportunities.forEach(function (obj) {
			if ( ownedOpportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.user.ownedOpportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ownedCasesIds as a OwnedCases
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwnedCases( userId, ownedCasesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = ownedCasesIds.split(',')

	// iterate over array of ownedCases ids
	idList.forEach(function (id) {
		// read the Case_
		var case_ = new Case_Service(this.http).getCase_(id);
		// add the Case_ if not already assigned
		if ( this.user.ownedCases.indexOf(case_) == -1 )
		this.user.ownedCases.push(case_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownedCasesIds as a OwnedCases
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwnedCases( userId, ownedCasesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= ownedCasesIds.split(',');
	var ownedCases 	= this.user.ownedCases;

	if ( ownedCases != null && ownedCasesIds != null ) {

		// iterate over array of ownedCases ids
		ownedCases.forEach(function (obj) {
			if ( ownedCasesIds.indexOf(obj._id) > -1 ) {
				// remove the Case_
				this.user.ownedCases.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( userId, quotesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.user.quotes.indexOf(quote) == -1 )
		this.user.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( userId, quotesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.user.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.user.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( userId, ordersIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the Order
		var order = new OrderService(this.http).getOrder(id);
		// add the Order if not already assigned
		if ( this.user.orders.indexOf(order) == -1 )
		this.user.orders.push(order);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( userId, ordersIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.user.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the Order
				this.user.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( userId, contractsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the Contract
		var contract = new ContractService(this.http).getContract(id);
		// add the Contract if not already assigned
		if ( this.user.contracts.indexOf(contract) == -1 )
		this.user.contracts.push(contract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( userId, contractsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.user.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the Contract
				this.user.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more emailMessagesIds as a EmailMessages
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmailMessages( userId, emailMessagesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = emailMessagesIds.split(',')

	// iterate over array of emailMessages ids
	idList.forEach(function (id) {
		// read the EmailMessage
		var emailMessage = new EmailMessageService(this.http).getEmailMessage(id);
		// add the EmailMessage if not already assigned
		if ( this.user.emailMessages.indexOf(emailMessage) == -1 )
		this.user.emailMessages.push(emailMessage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more emailMessagesIds as a EmailMessages
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmailMessages( userId, emailMessagesIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= emailMessagesIds.split(',');
	var emailMessages 	= this.user.emailMessages;

	if ( emailMessages != null && emailMessagesIds != null ) {

		// iterate over array of emailMessages ids
		emailMessages.forEach(function (obj) {
			if ( emailMessagesIds.indexOf(obj._id) > -1 ) {
				// remove the EmailMessage
				this.user.emailMessages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a User
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/User/update/' + this.user;

	return  this.http.post(uri_, this.user );
}

	//********************************************************************
	// loadHelper - internal helper to load a User
	//********************************************************************	
	loadHelper( id ) {
		this.getUser(id)
			.subscribe((res : User) => {
				this.user = res;
			});
	}
}