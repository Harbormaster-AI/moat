import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Agency} from '../models/Agency';
import {AdvertiserService} from '../services/Advertiser.service';
import {TeamService} from '../services/Team.service';
import {UserService} from '../services/User.service';
import {InsertionOrderService} from '../services/InsertionOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AgencyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	agency : Agency;

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
	// add a Agency
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders) : Observable<any> {
		const uri_ = this.apiUrl + '/Agency/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Advertisers: Advertisers != null && Advertisers.length > 0 ? Advertisers : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			InsertionOrders: InsertionOrders != null && InsertionOrders.length > 0 ? InsertionOrders : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Agency
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Agency/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Advertisers: Advertisers != null && Advertisers.length > 0 ? Advertisers : null,
      		Teams: Teams != null && Teams.length > 0 ? Teams : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
			InsertionOrders: InsertionOrders != null && InsertionOrders.length > 0 ? InsertionOrders : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Agency
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAgency(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Agency/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Agency
	// returns the results untouched as an Observable Agency
	// Agency model
	// delegates via URI
	//********************************************************************
	getAgency(id) : Observable<Agency> {
		const uri_ = this.apiUrl + '/Agency/load/' + id;

		return this.http.get<Agency>(uri_);
	}
	
	//********************************************************************
	// gets all Agency
	// returns the results untouched as JSON representation of an
	// Observable array of Agency models
	// delegates via URI
	//********************************************************************
	getAgencys() : Observable<Agency[]> {
		const uri_ = this.apiUrl + '/Agency/';

		return this
			.http.get<Agency[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more advertisersIds as a Advertisers
	// to a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdvertisers( agencyId, advertisersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );

	// split on a comma with no spaces
	var idList = advertisersIds.split(',')

	// iterate over array of advertisers ids
	idList.forEach(function (id) {
		// read the Advertiser
		var advertiser = new AdvertiserService(this.http).getAdvertiser(id);
		// add the Advertiser if not already assigned
		if ( this.agency.advertisers.indexOf(advertiser) == -1 )
		this.agency.advertisers.push(advertiser);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more advertisersIds as a Advertisers
	// from a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdvertisers( agencyId, advertisersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );


	// split on a comma with no spaces
	var idList 					= advertisersIds.split(',');
	var advertisers 	= this.agency.advertisers;

	if ( advertisers != null && advertisersIds != null ) {

		// iterate over array of advertisers ids
		advertisers.forEach(function (obj) {
			if ( advertisersIds.indexOf(obj._id) > -1 ) {
				// remove the Advertiser
				this.agency.advertisers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more teamsIds as a Teams
	// to a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTeams( agencyId, teamsIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );

	// split on a comma with no spaces
	var idList = teamsIds.split(',')

	// iterate over array of teams ids
	idList.forEach(function (id) {
		// read the Team
		var team = new TeamService(this.http).getTeam(id);
		// add the Team if not already assigned
		if ( this.agency.teams.indexOf(team) == -1 )
		this.agency.teams.push(team);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more teamsIds as a Teams
	// from a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTeams( agencyId, teamsIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );


	// split on a comma with no spaces
	var idList 					= teamsIds.split(',');
	var teams 	= this.agency.teams;

	if ( teams != null && teamsIds != null ) {

		// iterate over array of teams ids
		teams.forEach(function (obj) {
			if ( teamsIds.indexOf(obj._id) > -1 ) {
				// remove the Team
				this.agency.teams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more usersIds as a Users
	// to a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsers( agencyId, usersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );

	// split on a comma with no spaces
	var idList = usersIds.split(',')

	// iterate over array of users ids
	idList.forEach(function (id) {
		// read the User
		var user = new UserService(this.http).getUser(id);
		// add the User if not already assigned
		if ( this.agency.users.indexOf(user) == -1 )
		this.agency.users.push(user);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usersIds as a Users
	// from a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsers( agencyId, usersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );


	// split on a comma with no spaces
	var idList 					= usersIds.split(',');
	var users 	= this.agency.users;

	if ( users != null && usersIds != null ) {

		// iterate over array of users ids
		users.forEach(function (obj) {
			if ( usersIds.indexOf(obj._id) > -1 ) {
				// remove the User
				this.agency.users.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more insertionOrdersIds as a InsertionOrders
	// to a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsertionOrders( agencyId, insertionOrdersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );

	// split on a comma with no spaces
	var idList = insertionOrdersIds.split(',')

	// iterate over array of insertionOrders ids
	idList.forEach(function (id) {
		// read the InsertionOrder
		var insertionOrder = new InsertionOrderService(this.http).getInsertionOrder(id);
		// add the InsertionOrder if not already assigned
		if ( this.agency.insertionOrders.indexOf(insertionOrder) == -1 )
		this.agency.insertionOrders.push(insertionOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insertionOrdersIds as a InsertionOrders
	// from a Agency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsertionOrders( agencyId, insertionOrdersIds ): Observable<any> {

		// get the Agency
		this.loadHelper( agencyId );


	// split on a comma with no spaces
	var idList 					= insertionOrdersIds.split(',');
	var insertionOrders 	= this.agency.insertionOrders;

	if ( insertionOrders != null && insertionOrdersIds != null ) {

		// iterate over array of insertionOrders ids
		insertionOrders.forEach(function (obj) {
			if ( insertionOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the InsertionOrder
				this.agency.insertionOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Agency
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Agency/update/' + this.agency;

	return  this.http.post(uri_, this.agency );
}

	//********************************************************************
	// loadHelper - internal helper to load a Agency
	//********************************************************************	
	loadHelper( id ) {
		this.getAgency(id)
			.subscribe((res : Agency) => {
				this.agency = res;
			});
	}
}