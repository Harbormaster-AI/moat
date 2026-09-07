import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Organization} from '../models/Organization';
import {UserService} from '../services/User.service';
import {AccountService} from '../services/Account.service';
import {TeamService} from '../services/Team.service';
import {TerritoryService} from '../services/Territory.service';
import {ProductService} from '../services/Product.service';
import {PriceBookService} from '../services/PriceBook.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OrganizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	organization : Organization;

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
	// add a Organization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOrganization(name, defaultCurrency, defaultLocale, website, Users, Accounts, Teams, Territories, Products, PriceBooks, Campaigns) : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/create';
		const obj = {
			      		name: name,
      		defaultCurrency: defaultCurrency,
      		defaultLocale: defaultLocale,
      		website: website,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Territories: Territories != null && Territories.length > 0 ? Territories : null,
      		Products: Products != null && Products.length > 0 ? Products : null,
      		PriceBooks: PriceBooks != null && PriceBooks.length > 0 ? PriceBooks : null,
			Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOrganization(name, defaultCurrency, defaultLocale, website, Users, Accounts, Teams, Territories, Products, PriceBooks, Campaigns, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Organization/update/' + id;
		const obj = {
				      		name: name,
      		defaultCurrency: defaultCurrency,
      		defaultLocale: defaultLocale,
      		website: website,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Territories: Territories != null && Territories.length > 0 ? Territories : null,
      		Products: Products != null && Products.length > 0 ? Products : null,
      		PriceBooks: PriceBooks != null && PriceBooks.length > 0 ? PriceBooks : null,
			Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOrganization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Organization
	// returns the results untouched as an Observable Organization
	// Organization model
	// delegates via URI
	//********************************************************************
	getOrganization(id) : Observable<Organization> {
		const uri_ = this.apiUrl + '/Organization/load/' + id;

		return this.http.get<Organization>(uri_);
	}
	
	//********************************************************************
	// gets all Organization
	// returns the results untouched as JSON representation of an
	// Observable array of Organization models
	// delegates via URI
	//********************************************************************
	getOrganizations() : Observable<Organization[]> {
		const uri_ = this.apiUrl + '/Organization/';

		return this
			.http.get<Organization[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more usersIds as a Users
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsers( organizationId, usersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = usersIds.split(',')

	// iterate over array of users ids
	idList.forEach(function (id) {
		// read the User
		var user = new UserService(this.http).getUser(id);
		// add the User if not already assigned
		if ( this.organization.users.indexOf(user) == -1 )
		this.organization.users.push(user);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usersIds as a Users
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsers( organizationId, usersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= usersIds.split(',');
	var users 	= this.organization.users;

	if ( users != null && usersIds != null ) {

		// iterate over array of users ids
		users.forEach(function (obj) {
			if ( usersIds.indexOf(obj._id) > -1 ) {
				// remove the User
				this.organization.users.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( organizationId, accountsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.organization.accounts.indexOf(account) == -1 )
		this.organization.accounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( organizationId, accountsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.organization.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.organization.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more teamsIds as a Teams
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTeams( organizationId, teamsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = teamsIds.split(',')

	// iterate over array of teams ids
	idList.forEach(function (id) {
		// read the Team
		var team = new TeamService(this.http).getTeam(id);
		// add the Team if not already assigned
		if ( this.organization.teams.indexOf(team) == -1 )
		this.organization.teams.push(team);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more teamsIds as a Teams
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTeams( organizationId, teamsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= teamsIds.split(',');
	var teams 	= this.organization.teams;

	if ( teams != null && teamsIds != null ) {

		// iterate over array of teams ids
		teams.forEach(function (obj) {
			if ( teamsIds.indexOf(obj._id) > -1 ) {
				// remove the Team
				this.organization.teams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more territoriesIds as a Territories
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTerritories( organizationId, territoriesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = territoriesIds.split(',')

	// iterate over array of territories ids
	idList.forEach(function (id) {
		// read the Territory
		var territory = new TerritoryService(this.http).getTerritory(id);
		// add the Territory if not already assigned
		if ( this.organization.territories.indexOf(territory) == -1 )
		this.organization.territories.push(territory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more territoriesIds as a Territories
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTerritories( organizationId, territoriesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= territoriesIds.split(',');
	var territories 	= this.organization.territories;

	if ( territories != null && territoriesIds != null ) {

		// iterate over array of territories ids
		territories.forEach(function (obj) {
			if ( territoriesIds.indexOf(obj._id) > -1 ) {
				// remove the Territory
				this.organization.territories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more productsIds as a Products
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProducts( organizationId, productsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = productsIds.split(',')

	// iterate over array of products ids
	idList.forEach(function (id) {
		// read the Product
		var product = new ProductService(this.http).getProduct(id);
		// add the Product if not already assigned
		if ( this.organization.products.indexOf(product) == -1 )
		this.organization.products.push(product);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productsIds as a Products
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProducts( organizationId, productsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= productsIds.split(',');
	var products 	= this.organization.products;

	if ( products != null && productsIds != null ) {

		// iterate over array of products ids
		products.forEach(function (obj) {
			if ( productsIds.indexOf(obj._id) > -1 ) {
				// remove the Product
				this.organization.products.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more priceBooksIds as a PriceBooks
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPriceBooks( organizationId, priceBooksIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = priceBooksIds.split(',')

	// iterate over array of priceBooks ids
	idList.forEach(function (id) {
		// read the PriceBook
		var priceBook = new PriceBookService(this.http).getPriceBook(id);
		// add the PriceBook if not already assigned
		if ( this.organization.priceBooks.indexOf(priceBook) == -1 )
		this.organization.priceBooks.push(priceBook);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more priceBooksIds as a PriceBooks
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePriceBooks( organizationId, priceBooksIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= priceBooksIds.split(',');
	var priceBooks 	= this.organization.priceBooks;

	if ( priceBooks != null && priceBooksIds != null ) {

		// iterate over array of priceBooks ids
		priceBooks.forEach(function (obj) {
			if ( priceBooksIds.indexOf(obj._id) > -1 ) {
				// remove the PriceBook
				this.organization.priceBooks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( organizationId, campaignsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.organization.campaigns.indexOf(campaign) == -1 )
		this.organization.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( organizationId, campaignsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.organization.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.organization.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Organization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Organization/update/' + this.organization;

	return  this.http.post(uri_, this.organization );
}

	//********************************************************************
	// loadHelper - internal helper to load a Organization
	//********************************************************************	
	loadHelper( id ) {
		this.getOrganization(id)
			.subscribe((res : Organization) => {
				this.organization = res;
			});
	}
}