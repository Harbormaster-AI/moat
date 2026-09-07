import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Contract} from '../models/Contract';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {UserService} from '../services/User.service';
import {OrderService} from '../services/Order.service';
import {Case_Service} from '../services/Case_.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ContractService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	contract : Contract;

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
	// add a Contract
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addContract(contractNumber, startDate, endDate, renewalTermMonths, autoRenew, Organization, Account, Owner, Orders, Cases, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Contract/create';
		const obj = {
			      		contractNumber: contractNumber,
      		startDate: startDate,
      		endDate: endDate,
      		renewalTermMonths: renewalTermMonths,
      		autoRenew: autoRenew,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateContract(contractNumber, startDate, endDate, renewalTermMonths, autoRenew, Organization, Account, Owner, Orders, Cases, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Contract/update/' + id;
		const obj = {
				      		contractNumber: contractNumber,
      		startDate: startDate,
      		endDate: endDate,
      		renewalTermMonths: renewalTermMonths,
      		autoRenew: autoRenew,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Cases: Cases != null && Cases.length > 0 ? Cases : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteContract(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Contract/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Contract
	// returns the results untouched as an Observable Contract
	// Contract model
	// delegates via URI
	//********************************************************************
	getContract(id) : Observable<Contract> {
		const uri_ = this.apiUrl + '/Contract/load/' + id;

		return this.http.get<Contract>(uri_);
	}
	
	//********************************************************************
	// gets all Contract
	// returns the results untouched as JSON representation of an
	// Observable array of Contract models
	// delegates via URI
	//********************************************************************
	getContracts() : Observable<Contract[]> {
		const uri_ = this.apiUrl + '/Contract/';

		return this
			.http.get<Contract[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( contractId, _organizationId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.contract.organization = tmp;

	// save the Contract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( contractId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// assign Organization to null
	this.contract.organization = null;

	// save the Contract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( contractId, _accountId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.contract.account = tmp;

	// save the Contract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( contractId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// assign Account to null
	this.contract.account = null;

	// save the Contract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( contractId, _ownerId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.contract.owner = tmp;

	// save the Contract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( contractId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// assign Owner to null
	this.contract.owner = null;

	// save the Contract
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( contractId, ordersIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the Order
		var order = new OrderService(this.http).getOrder(id);
		// add the Order if not already assigned
		if ( this.contract.orders.indexOf(order) == -1 )
		this.contract.orders.push(order);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( contractId, ordersIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.contract.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the Order
				this.contract.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more casesIds as a Cases
	// to a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCases( contractId, casesIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );

	// split on a comma with no spaces
	var idList = casesIds.split(',')

	// iterate over array of cases ids
	idList.forEach(function (id) {
		// read the Case_
		var case_ = new Case_Service(this.http).getCase_(id);
		// add the Case_ if not already assigned
		if ( this.contract.cases.indexOf(case_) == -1 )
		this.contract.cases.push(case_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more casesIds as a Cases
	// from a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCases( contractId, casesIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );


	// split on a comma with no spaces
	var idList 					= casesIds.split(',');
	var cases 	= this.contract.cases;

	if ( cases != null && casesIds != null ) {

		// iterate over array of cases ids
		cases.forEach(function (obj) {
			if ( casesIds.indexOf(obj._id) > -1 ) {
				// remove the Case_
				this.contract.cases.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Contract
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Contract/update/' + this.contract;

	return  this.http.post(uri_, this.contract );
}

	//********************************************************************
	// loadHelper - internal helper to load a Contract
	//********************************************************************	
	loadHelper( id ) {
		this.getContract(id)
			.subscribe((res : Contract) => {
				this.contract = res;
			});
	}
}