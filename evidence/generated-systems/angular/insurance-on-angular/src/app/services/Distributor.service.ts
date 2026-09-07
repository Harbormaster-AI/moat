import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Distributor} from '../models/Distributor';
import {InsurerService} from '../services/Insurer.service';
import {AgentService} from '../services/Agent.service';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DistributorService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	distributor : Distributor;

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
	// add a Distributor
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType) : Observable<any> {
		const uri_ = this.apiUrl + '/Distributor/create';
		const obj = {
			      		name: name,
      		licenseNumber: licenseNumber,
      		region: region,
      		Insurers: Insurers != null && Insurers.length > 0 ? Insurers : null,
      		Agents: Agents != null && Agents.length > 0 ? Agents : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
			DistributorType: DistributorType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Distributor
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Distributor/update/' + id;
		const obj = {
				      		name: name,
      		licenseNumber: licenseNumber,
      		region: region,
      		Insurers: Insurers != null && Insurers.length > 0 ? Insurers : null,
      		Agents: Agents != null && Agents.length > 0 ? Agents : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
			DistributorType: DistributorType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Distributor
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDistributor(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Distributor/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Distributor
	// returns the results untouched as an Observable Distributor
	// Distributor model
	// delegates via URI
	//********************************************************************
	getDistributor(id) : Observable<Distributor> {
		const uri_ = this.apiUrl + '/Distributor/load/' + id;

		return this.http.get<Distributor>(uri_);
	}
	
	//********************************************************************
	// gets all Distributor
	// returns the results untouched as JSON representation of an
	// Observable array of Distributor models
	// delegates via URI
	//********************************************************************
	getDistributors() : Observable<Distributor[]> {
		const uri_ = this.apiUrl + '/Distributor/';

		return this
			.http.get<Distributor[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more insurersIds as a Insurers
	// to a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsurers( distributorId, insurersIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );

	// split on a comma with no spaces
	var idList = insurersIds.split(',')

	// iterate over array of insurers ids
	idList.forEach(function (id) {
		// read the Insurer
		var insurer = new InsurerService(this.http).getInsurer(id);
		// add the Insurer if not already assigned
		if ( this.distributor.insurers.indexOf(insurer) == -1 )
		this.distributor.insurers.push(insurer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insurersIds as a Insurers
	// from a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsurers( distributorId, insurersIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );


	// split on a comma with no spaces
	var idList 					= insurersIds.split(',');
	var insurers 	= this.distributor.insurers;

	if ( insurers != null && insurersIds != null ) {

		// iterate over array of insurers ids
		insurers.forEach(function (obj) {
			if ( insurersIds.indexOf(obj._id) > -1 ) {
				// remove the Insurer
				this.distributor.insurers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more agentsIds as a Agents
	// to a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAgents( distributorId, agentsIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );

	// split on a comma with no spaces
	var idList = agentsIds.split(',')

	// iterate over array of agents ids
	idList.forEach(function (id) {
		// read the Agent
		var agent = new AgentService(this.http).getAgent(id);
		// add the Agent if not already assigned
		if ( this.distributor.agents.indexOf(agent) == -1 )
		this.distributor.agents.push(agent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more agentsIds as a Agents
	// from a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAgents( distributorId, agentsIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );


	// split on a comma with no spaces
	var idList 					= agentsIds.split(',');
	var agents 	= this.distributor.agents;

	if ( agents != null && agentsIds != null ) {

		// iterate over array of agents ids
		agents.forEach(function (obj) {
			if ( agentsIds.indexOf(obj._id) > -1 ) {
				// remove the Agent
				this.distributor.agents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( distributorId, policiesIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.distributor.policies.indexOf(policy) == -1 )
		this.distributor.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Distributor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( distributorId, policiesIds ): Observable<any> {

		// get the Distributor
		this.loadHelper( distributorId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.distributor.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.distributor.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Distributor
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Distributor/update/' + this.distributor;

	return  this.http.post(uri_, this.distributor );
}

	//********************************************************************
	// loadHelper - internal helper to load a Distributor
	//********************************************************************	
	loadHelper( id ) {
		this.getDistributor(id)
			.subscribe((res : Distributor) => {
				this.distributor = res;
			});
	}
}