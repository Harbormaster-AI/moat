import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {User} from '../models/User';
import {AgencyService} from '../services/Agency.service';
import {TeamService} from '../services/Team.service';
import {AdAccountService} from '../services/AdAccount.service';
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
	addUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role) : Observable<any> {
		const uri_ = this.apiUrl + '/User/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		email: email,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
			Role: Role
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUser(firstName, lastName, email, Agency, Teams, AdAccounts, Role, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/User/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		email: email,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
			Role: Role
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
	// assigns a Agency on a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAgency( userId, _agencyId ): Observable<any> {

		// get the User from storage
		this.loadHelper( userId );

	// get the Agency from storage
	var tmp 	= new AgencyService(this.http).getAgency(_agencyId);

	// assign the Agency
	this.user.agency = tmp;

	// save the User
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Agency on a User
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAgency( userId ): Observable<any> {

		// get the User from storage
		this.loadHelper( userId );

	// assign Agency to null
	this.user.agency = null;

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
	// adds one or more adAccountsIds as a AdAccounts
	// to a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdAccounts( userId, adAccountsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );

	// split on a comma with no spaces
	var idList = adAccountsIds.split(',')

	// iterate over array of adAccounts ids
	idList.forEach(function (id) {
		// read the AdAccount
		var adAccount = new AdAccountService(this.http).getAdAccount(id);
		// add the AdAccount if not already assigned
		if ( this.user.adAccounts.indexOf(adAccount) == -1 )
		this.user.adAccounts.push(adAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adAccountsIds as a AdAccounts
	// from a User
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdAccounts( userId, adAccountsIds ): Observable<any> {

		// get the User
		this.loadHelper( userId );


	// split on a comma with no spaces
	var idList 					= adAccountsIds.split(',');
	var adAccounts 	= this.user.adAccounts;

	if ( adAccounts != null && adAccountsIds != null ) {

		// iterate over array of adAccounts ids
		adAccounts.forEach(function (obj) {
			if ( adAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the AdAccount
				this.user.adAccounts.pop(obj);
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