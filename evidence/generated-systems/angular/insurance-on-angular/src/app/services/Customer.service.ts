import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Customer} from '../models/Customer';
import {ApplicationService} from '../services/Application.service';
import {PolicyService} from '../services/Policy.service';
import {ClaimService} from '../services/Claim.service';
import {AgentService} from '../services/Agent.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CustomerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	customer : Customer;

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
	// add a Customer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCustomer(firstName, lastName, organizationName, taxId, dateOfBirth, primaryAddress, Applications, Policies, Claims, Agents, Beneficiaries, CustomerType) : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		organizationName: organizationName,
      		taxId: taxId,
      		dateOfBirth: dateOfBirth,
      		primaryAddress: primaryAddress,
      		Applications: Applications != null && Applications.length > 0 ? Applications : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Agents: Agents != null && Agents.length > 0 ? Agents : null,
      		Beneficiaries: Beneficiaries != null && Beneficiaries.length > 0 ? Beneficiaries : null,
			CustomerType: CustomerType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCustomer(firstName, lastName, organizationName, taxId, dateOfBirth, primaryAddress, Applications, Policies, Claims, Agents, Beneficiaries, CustomerType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Customer/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		organizationName: organizationName,
      		taxId: taxId,
      		dateOfBirth: dateOfBirth,
      		primaryAddress: primaryAddress,
      		Applications: Applications != null && Applications.length > 0 ? Applications : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Agents: Agents != null && Agents.length > 0 ? Agents : null,
      		Beneficiaries: Beneficiaries != null && Beneficiaries.length > 0 ? Beneficiaries : null,
			CustomerType: CustomerType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCustomer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Customer
	// returns the results untouched as an Observable Customer
	// Customer model
	// delegates via URI
	//********************************************************************
	getCustomer(id) : Observable<Customer> {
		const uri_ = this.apiUrl + '/Customer/load/' + id;

		return this.http.get<Customer>(uri_);
	}
	
	//********************************************************************
	// gets all Customer
	// returns the results untouched as JSON representation of an
	// Observable array of Customer models
	// delegates via URI
	//********************************************************************
	getCustomers() : Observable<Customer[]> {
		const uri_ = this.apiUrl + '/Customer/';

		return this
			.http.get<Customer[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more applicationsIds as a Applications
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApplications( customerId, applicationsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = applicationsIds.split(',')

	// iterate over array of applications ids
	idList.forEach(function (id) {
		// read the Application
		var application = new ApplicationService(this.http).getApplication(id);
		// add the Application if not already assigned
		if ( this.customer.applications.indexOf(application) == -1 )
		this.customer.applications.push(application);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more applicationsIds as a Applications
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApplications( customerId, applicationsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= applicationsIds.split(',');
	var applications 	= this.customer.applications;

	if ( applications != null && applicationsIds != null ) {

		// iterate over array of applications ids
		applications.forEach(function (obj) {
			if ( applicationsIds.indexOf(obj._id) > -1 ) {
				// remove the Application
				this.customer.applications.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( customerId, policiesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.customer.policies.indexOf(policy) == -1 )
		this.customer.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( customerId, policiesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.customer.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.customer.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( customerId, claimsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.customer.claims.indexOf(claim) == -1 )
		this.customer.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( customerId, claimsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.customer.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.customer.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more agentsIds as a Agents
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAgents( customerId, agentsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = agentsIds.split(',')

	// iterate over array of agents ids
	idList.forEach(function (id) {
		// read the Agent
		var agent = new AgentService(this.http).getAgent(id);
		// add the Agent if not already assigned
		if ( this.customer.agents.indexOf(agent) == -1 )
		this.customer.agents.push(agent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more agentsIds as a Agents
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAgents( customerId, agentsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= agentsIds.split(',');
	var agents 	= this.customer.agents;

	if ( agents != null && agentsIds != null ) {

		// iterate over array of agents ids
		agents.forEach(function (obj) {
			if ( agentsIds.indexOf(obj._id) > -1 ) {
				// remove the Agent
				this.customer.agents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more beneficiariesIds as a Beneficiaries
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBeneficiaries( customerId, beneficiariesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = beneficiariesIds.split(',')

	// iterate over array of beneficiaries ids
	idList.forEach(function (id) {
		// read the Beneficiary
		var beneficiary = new BeneficiaryService(this.http).getBeneficiary(id);
		// add the Beneficiary if not already assigned
		if ( this.customer.beneficiaries.indexOf(beneficiary) == -1 )
		this.customer.beneficiaries.push(beneficiary);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more beneficiariesIds as a Beneficiaries
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBeneficiaries( customerId, beneficiariesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= beneficiariesIds.split(',');
	var beneficiaries 	= this.customer.beneficiaries;

	if ( beneficiaries != null && beneficiariesIds != null ) {

		// iterate over array of beneficiaries ids
		beneficiaries.forEach(function (obj) {
			if ( beneficiariesIds.indexOf(obj._id) > -1 ) {
				// remove the Beneficiary
				this.customer.beneficiaries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Customer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Customer/update/' + this.customer;

	return  this.http.post(uri_, this.customer );
}

	//********************************************************************
	// loadHelper - internal helper to load a Customer
	//********************************************************************	
	loadHelper( id ) {
		this.getCustomer(id)
			.subscribe((res : Customer) => {
				this.customer = res;
			});
	}
}