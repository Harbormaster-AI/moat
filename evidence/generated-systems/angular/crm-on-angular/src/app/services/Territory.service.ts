import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Territory} from '../models/Territory';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {UserService} from '../services/User.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TerritoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	territory : Territory;

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
	// add a Territory
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTerritory(name, region, Organization, Accounts, Users, TerritoryType) : Observable<any> {
		const uri_ = this.apiUrl + '/Territory/create';
		const obj = {
			      		name: name,
      		region: region,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			TerritoryType: TerritoryType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Territory
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTerritory(name, region, Organization, Accounts, Users, TerritoryType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Territory/update/' + id;
		const obj = {
				      		name: name,
      		region: region,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			TerritoryType: TerritoryType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Territory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTerritory(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Territory/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Territory
	// returns the results untouched as an Observable Territory
	// Territory model
	// delegates via URI
	//********************************************************************
	getTerritory(id) : Observable<Territory> {
		const uri_ = this.apiUrl + '/Territory/load/' + id;

		return this.http.get<Territory>(uri_);
	}
	
	//********************************************************************
	// gets all Territory
	// returns the results untouched as JSON representation of an
	// Observable array of Territory models
	// delegates via URI
	//********************************************************************
	getTerritorys() : Observable<Territory[]> {
		const uri_ = this.apiUrl + '/Territory/';

		return this
			.http.get<Territory[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Territory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( territoryId, _organizationId ): Observable<any> {

		// get the Territory from storage
		this.loadHelper( territoryId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.territory.organization = tmp;

	// save the Territory
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Territory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( territoryId ): Observable<any> {

		// get the Territory from storage
		this.loadHelper( territoryId );

	// assign Organization to null
	this.territory.organization = null;

	// save the Territory
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a Territory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( territoryId, accountsIds ): Observable<any> {

		// get the Territory
		this.loadHelper( territoryId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.territory.accounts.indexOf(account) == -1 )
		this.territory.accounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a Territory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( territoryId, accountsIds ): Observable<any> {

		// get the Territory
		this.loadHelper( territoryId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.territory.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.territory.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more usersIds as a Users
	// to a Territory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsers( territoryId, usersIds ): Observable<any> {

		// get the Territory
		this.loadHelper( territoryId );

	// split on a comma with no spaces
	var idList = usersIds.split(',')

	// iterate over array of users ids
	idList.forEach(function (id) {
		// read the User
		var user = new UserService(this.http).getUser(id);
		// add the User if not already assigned
		if ( this.territory.users.indexOf(user) == -1 )
		this.territory.users.push(user);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usersIds as a Users
	// from a Territory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsers( territoryId, usersIds ): Observable<any> {

		// get the Territory
		this.loadHelper( territoryId );


	// split on a comma with no spaces
	var idList 					= usersIds.split(',');
	var users 	= this.territory.users;

	if ( users != null && usersIds != null ) {

		// iterate over array of users ids
		users.forEach(function (obj) {
			if ( usersIds.indexOf(obj._id) > -1 ) {
				// remove the User
				this.territory.users.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Territory
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Territory/update/' + this.territory;

	return  this.http.post(uri_, this.territory );
}

	//********************************************************************
	// loadHelper - internal helper to load a Territory
	//********************************************************************	
	loadHelper( id ) {
		this.getTerritory(id)
			.subscribe((res : Territory) => {
				this.territory = res;
			});
	}
}