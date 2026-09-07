import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Team} from '../models/Team';
import {AgencyService} from '../services/Agency.service';
import {UserService} from '../services/User.service';
import {AdAccountService} from '../services/AdAccount.service';
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
	addTeam(name, Agency, Users, AdAccounts) : Observable<any> {
		const uri_ = this.apiUrl + '/Team/create';
		const obj = {
			      		name: name,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTeam(name, Agency, Users, AdAccounts, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Team/update/' + id;
		const obj = {
				      		name: name,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null
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
	// assigns a Agency on a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAgency( teamId, _agencyId ): Observable<any> {

		// get the Team from storage
		this.loadHelper( teamId );

	// get the Agency from storage
	var tmp 	= new AgencyService(this.http).getAgency(_agencyId);

	// assign the Agency
	this.team.agency = tmp;

	// save the Team
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Agency on a Team
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAgency( teamId ): Observable<any> {

		// get the Team from storage
		this.loadHelper( teamId );

	// assign Agency to null
	this.team.agency = null;

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
	// adds one or more adAccountsIds as a AdAccounts
	// to a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdAccounts( teamId, adAccountsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );

	// split on a comma with no spaces
	var idList = adAccountsIds.split(',')

	// iterate over array of adAccounts ids
	idList.forEach(function (id) {
		// read the AdAccount
		var adAccount = new AdAccountService(this.http).getAdAccount(id);
		// add the AdAccount if not already assigned
		if ( this.team.adAccounts.indexOf(adAccount) == -1 )
		this.team.adAccounts.push(adAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adAccountsIds as a AdAccounts
	// from a Team
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdAccounts( teamId, adAccountsIds ): Observable<any> {

		// get the Team
		this.loadHelper( teamId );


	// split on a comma with no spaces
	var idList 					= adAccountsIds.split(',');
	var adAccounts 	= this.team.adAccounts;

	if ( adAccounts != null && adAccountsIds != null ) {

		// iterate over array of adAccounts ids
		adAccounts.forEach(function (obj) {
			if ( adAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the AdAccount
				this.team.adAccounts.pop(obj);
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