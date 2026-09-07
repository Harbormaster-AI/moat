import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Insurer} from '../models/Insurer';
import {InsuranceProductService} from '../services/InsuranceProduct.service';
import {DistributorService} from '../services/Distributor.service';
import {PolicyService} from '../services/Policy.service';
import {ClaimService} from '../services/Claim.service';
import {ReinsuranceAgreementService} from '../services/ReinsuranceAgreement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsurerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insurer : Insurer;

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
	// add a Insurer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements) : Observable<any> {
		const uri_ = this.apiUrl + '/Insurer/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		domicileCountry: domicileCountry,
      		naicNumber: naicNumber,
      		website: website,
      		Products: Products != null && Products.length > 0 ? Products : null,
      		DistributionPartners: DistributionPartners != null && DistributionPartners.length > 0 ? DistributionPartners : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
			ReinsuranceAgreements: ReinsuranceAgreements != null && ReinsuranceAgreements.length > 0 ? ReinsuranceAgreements : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Insurer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Insurer/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		domicileCountry: domicileCountry,
      		naicNumber: naicNumber,
      		website: website,
      		Products: Products != null && Products.length > 0 ? Products : null,
      		DistributionPartners: DistributionPartners != null && DistributionPartners.length > 0 ? DistributionPartners : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
			ReinsuranceAgreements: ReinsuranceAgreements != null && ReinsuranceAgreements.length > 0 ? ReinsuranceAgreements : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Insurer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsurer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Insurer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Insurer
	// returns the results untouched as an Observable Insurer
	// Insurer model
	// delegates via URI
	//********************************************************************
	getInsurer(id) : Observable<Insurer> {
		const uri_ = this.apiUrl + '/Insurer/load/' + id;

		return this.http.get<Insurer>(uri_);
	}
	
	//********************************************************************
	// gets all Insurer
	// returns the results untouched as JSON representation of an
	// Observable array of Insurer models
	// delegates via URI
	//********************************************************************
	getInsurers() : Observable<Insurer[]> {
		const uri_ = this.apiUrl + '/Insurer/';

		return this
			.http.get<Insurer[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more productsIds as a Products
	// to a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProducts( insurerId, productsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );

	// split on a comma with no spaces
	var idList = productsIds.split(',')

	// iterate over array of products ids
	idList.forEach(function (id) {
		// read the InsuranceProduct
		var insuranceProduct = new InsuranceProductService(this.http).getInsuranceProduct(id);
		// add the InsuranceProduct if not already assigned
		if ( this.insurer.products.indexOf(insuranceProduct) == -1 )
		this.insurer.products.push(insuranceProduct);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productsIds as a Products
	// from a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProducts( insurerId, productsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );


	// split on a comma with no spaces
	var idList 					= productsIds.split(',');
	var products 	= this.insurer.products;

	if ( products != null && productsIds != null ) {

		// iterate over array of products ids
		products.forEach(function (obj) {
			if ( productsIds.indexOf(obj._id) > -1 ) {
				// remove the InsuranceProduct
				this.insurer.products.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more distributionPartnersIds as a DistributionPartners
	// to a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDistributionPartners( insurerId, distributionPartnersIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );

	// split on a comma with no spaces
	var idList = distributionPartnersIds.split(',')

	// iterate over array of distributionPartners ids
	idList.forEach(function (id) {
		// read the Distributor
		var distributor = new DistributorService(this.http).getDistributor(id);
		// add the Distributor if not already assigned
		if ( this.insurer.distributionPartners.indexOf(distributor) == -1 )
		this.insurer.distributionPartners.push(distributor);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more distributionPartnersIds as a DistributionPartners
	// from a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDistributionPartners( insurerId, distributionPartnersIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );


	// split on a comma with no spaces
	var idList 					= distributionPartnersIds.split(',');
	var distributionPartners 	= this.insurer.distributionPartners;

	if ( distributionPartners != null && distributionPartnersIds != null ) {

		// iterate over array of distributionPartners ids
		distributionPartners.forEach(function (obj) {
			if ( distributionPartnersIds.indexOf(obj._id) > -1 ) {
				// remove the Distributor
				this.insurer.distributionPartners.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( insurerId, policiesIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.insurer.policies.indexOf(policy) == -1 )
		this.insurer.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( insurerId, policiesIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.insurer.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.insurer.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( insurerId, claimsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.insurer.claims.indexOf(claim) == -1 )
		this.insurer.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( insurerId, claimsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.insurer.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.insurer.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements
	// to a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReinsuranceAgreements( insurerId, reinsuranceAgreementsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );

	// split on a comma with no spaces
	var idList = reinsuranceAgreementsIds.split(',')

	// iterate over array of reinsuranceAgreements ids
	idList.forEach(function (id) {
		// read the ReinsuranceAgreement
		var reinsuranceAgreement = new ReinsuranceAgreementService(this.http).getReinsuranceAgreement(id);
		// add the ReinsuranceAgreement if not already assigned
		if ( this.insurer.reinsuranceAgreements.indexOf(reinsuranceAgreement) == -1 )
		this.insurer.reinsuranceAgreements.push(reinsuranceAgreement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements
	// from a Insurer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReinsuranceAgreements( insurerId, reinsuranceAgreementsIds ): Observable<any> {

		// get the Insurer
		this.loadHelper( insurerId );


	// split on a comma with no spaces
	var idList 					= reinsuranceAgreementsIds.split(',');
	var reinsuranceAgreements 	= this.insurer.reinsuranceAgreements;

	if ( reinsuranceAgreements != null && reinsuranceAgreementsIds != null ) {

		// iterate over array of reinsuranceAgreements ids
		reinsuranceAgreements.forEach(function (obj) {
			if ( reinsuranceAgreementsIds.indexOf(obj._id) > -1 ) {
				// remove the ReinsuranceAgreement
				this.insurer.reinsuranceAgreements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Insurer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Insurer/update/' + this.insurer;

	return  this.http.post(uri_, this.insurer );
}

	//********************************************************************
	// loadHelper - internal helper to load a Insurer
	//********************************************************************	
	loadHelper( id ) {
		this.getInsurer(id)
			.subscribe((res : Insurer) => {
				this.insurer = res;
			});
	}
}