import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Agent} from '../models/Agent';
import {DistributorService} from '../services/Distributor.service';
import {PolicyService} from '../services/Policy.service';
import {CustomerService} from '../services/Customer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AgentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	agent : Agent;

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
	// add a Agent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Agent/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		licenseId: licenseId,
      		Distributor: Distributor != null && Distributor.length > 0 ? Distributor : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Customers: Customers != null && Customers.length > 0 ? Customers : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Agent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Agent/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		licenseId: licenseId,
      		Distributor: Distributor != null && Distributor.length > 0 ? Distributor : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Customers: Customers != null && Customers.length > 0 ? Customers : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Agent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAgent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Agent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Agent
	// returns the results untouched as an Observable Agent
	// Agent model
	// delegates via URI
	//********************************************************************
	getAgent(id) : Observable<Agent> {
		const uri_ = this.apiUrl + '/Agent/load/' + id;

		return this.http.get<Agent>(uri_);
	}
	
	//********************************************************************
	// gets all Agent
	// returns the results untouched as JSON representation of an
	// Observable array of Agent models
	// delegates via URI
	//********************************************************************
	getAgents() : Observable<Agent[]> {
		const uri_ = this.apiUrl + '/Agent/';

		return this
			.http.get<Agent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Distributor on a Agent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDistributor( agentId, _distributorId ): Observable<any> {

		// get the Agent from storage
		this.loadHelper( agentId );

	// get the Distributor from storage
	var tmp 	= new DistributorService(this.http).getDistributor(_distributorId);

	// assign the Distributor
	this.agent.distributor = tmp;

	// save the Agent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Distributor on a Agent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDistributor( agentId ): Observable<any> {

		// get the Agent from storage
		this.loadHelper( agentId );

	// assign Distributor to null
	this.agent.distributor = null;

	// save the Agent
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Agent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( agentId, policiesIds ): Observable<any> {

		// get the Agent
		this.loadHelper( agentId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.agent.policies.indexOf(policy) == -1 )
		this.agent.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Agent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( agentId, policiesIds ): Observable<any> {

		// get the Agent
		this.loadHelper( agentId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.agent.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.agent.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more customersIds as a Customers
	// to a Agent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCustomers( agentId, customersIds ): Observable<any> {

		// get the Agent
		this.loadHelper( agentId );

	// split on a comma with no spaces
	var idList = customersIds.split(',')

	// iterate over array of customers ids
	idList.forEach(function (id) {
		// read the Customer
		var customer = new CustomerService(this.http).getCustomer(id);
		// add the Customer if not already assigned
		if ( this.agent.customers.indexOf(customer) == -1 )
		this.agent.customers.push(customer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more customersIds as a Customers
	// from a Agent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCustomers( agentId, customersIds ): Observable<any> {

		// get the Agent
		this.loadHelper( agentId );


	// split on a comma with no spaces
	var idList 					= customersIds.split(',');
	var customers 	= this.agent.customers;

	if ( customers != null && customersIds != null ) {

		// iterate over array of customers ids
		customers.forEach(function (obj) {
			if ( customersIds.indexOf(obj._id) > -1 ) {
				// remove the Customer
				this.agent.customers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Agent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Agent/update/' + this.agent;

	return  this.http.post(uri_, this.agent );
}

	//********************************************************************
	// loadHelper - internal helper to load a Agent
	//********************************************************************	
	loadHelper( id ) {
		this.getAgent(id)
			.subscribe((res : Agent) => {
				this.agent = res;
			});
	}
}