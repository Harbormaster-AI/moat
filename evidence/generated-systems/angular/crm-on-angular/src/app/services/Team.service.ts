import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Team} from '../models/Team';
import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {AccountService} from '../services/Account.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TeamService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	team : Team;

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
	// add a Team
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTeam(name, Organization, Users, Accounts, Opportunities, Cases, Campaigns, TeamType) : Observable<any> {
		const uri_ = this.apiUrl + '/Team/create';
		const obj = {
			      		name: name,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			TeamType: TeamType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTeam(name, Organization, Users, Accounts, Opportunities, Cases, Campaigns, TeamType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Team/update/' + id;
		const obj = {
				      		name: name,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Opportunities: Opportunities != null && Opportunities.length > 0 ? Opportunities : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			TeamType: TeamType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTeam(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Team/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Team
	// returns the results untouched as an Observable Team
	// Team model
	// delegates via URI
	//********************************************************************
	getTeam(id) : Observable<Team> {
		const uri_ = this.apiUrl + '/Team/load/' + id;

		return this.http.get<Team>(uri_);
	}
	
	//********************************************************************
	// gets all Team
	// returns the results untouched as JSON representation of an
	// Observable array of Team models
	// delegates via URI
	//********************************************************************
	getTeams() : Observable<Team[]> {
		const uri_ = this.apiUrl + '/Team/';

		return this
			.http.get<Team[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( teamId, _organizationId ): Observable<any> {

		// get the Team from storage
		this.loadHelper( teamId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.team.organization = tmp;

	// save the Team
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( teamId ): Observable<any> {

		// get the Team from storage
		this.loadHelper( teamId );

	// assign Organization to null
	this.team.organization = null;

	// save the Team
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more usersIds as a Users
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsers( teamId, usersIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = usersIds.split(',')

	// iterate over array of users ids
	idList.forEach(function (id) {
		// read the User
		var user = new UserService(this.http).getUser(id);
		// add the User if not already assigned
		if ( this.team.users.indexOf(user) == -1 )
		this.team.users.push(user);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usersIds as a Users
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsers( teamId, usersIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= usersIds.split(',');
	var users 	= this.team.users;

	if ( users != null && usersIds != null ) {

		// iterate over array of users ids
		users.forEach(function (obj) {
			if ( usersIds.indexOf(obj._id) > -1 ) {
				// remove the User
				this.team.users.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( teamId, accountsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.team.accounts.indexOf(account) == -1 )
		this.team.accounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( teamId, accountsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.team.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.team.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more opportunitiesIds as a Opportunities
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOpportunities( teamId, opportunitiesIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = opportunitiesIds.split(',')

	// iterate over array of opportunities ids
	idList.forEach(function (id) {
		// read the Opportunity
		var opportunity = new OpportunityService(this.http).getOpportunity(id);
		// add the Opportunity if not already assigned
		if ( this.team.opportunities.indexOf(opportunity) == -1 )
		this.team.opportunities.push(opportunity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more opportunitiesIds as a Opportunities
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOpportunities( teamId, opportunitiesIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= opportunitiesIds.split(',');
	var opportunities 	= this.team.opportunities;

	if ( opportunities != null && opportunitiesIds != null ) {

		// iterate over array of opportunities ids
		opportunities.forEach(function (obj) {
			if ( opportunitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Opportunity
				this.team.opportunities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more casesIds as a Cases
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCases( teamId, casesIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = casesIds.split(',')

	// iterate over array of cases ids
	idList.forEach(function (id) {
		// read the Case_
		var case_ = new Case_Service(this.http).getCase_(id);
		// add the Case_ if not already assigned
		if ( this.team.cases.indexOf(case_) == -1 )
		this.team.cases.push(case_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more casesIds as a Cases
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCases( teamId, casesIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= casesIds.split(',');
	var cases 	= this.team.cases;

	if ( cases != null && casesIds != null ) {

		// iterate over array of cases ids
		cases.forEach(function (obj) {
			if ( casesIds.indexOf(obj._id) > -1 ) {
				// remove the Case_
				this.team.cases.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( teamId, campaignsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.team.campaigns.indexOf(campaign) == -1 )
		this.team.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( teamId, campaignsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.team.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.team.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Team
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Team/update/' + this.team;

	return  this.http.post(uri_, this.team );
}

	//********************************************************************
	// loadHelper - internal helper to load a Team
	//********************************************************************	
	loadHelper( id ) {
		this.getTeam(id)
			.subscribe((res : Team) => {
				this.team = res;
			});
	}
}