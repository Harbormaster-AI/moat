import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Policy} from '../models/Policy';
import {InsurerService} from '../services/Insurer.service';
import {CustomerService} from '../services/Customer.service';
import {InsuranceProductService} from '../services/InsuranceProduct.service';
import {AgentService} from '../services/Agent.service';
import {PolicyCoverageService} from '../services/PolicyCoverage.service';
import {InsuredObjectService} from '../services/InsuredObject.service';
import {EndorsementService} from '../services/Endorsement.service';
import {BillingAccountService} from '../services/BillingAccount.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import {ClaimService} from '../services/Claim.service';
import {ReinsuranceAgreementService} from '../services/ReinsuranceAgreement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	policy : Policy;

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
	// add a Policy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan) : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/create';
		const obj = {
			      		policyNumber: policyNumber,
      		effectivePeriod: effectivePeriod,
      		totalPremium: totalPremium,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		Agent: Agent != null && Agent.length > 0 ? Agent : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
      		Endorsements: Endorsements != null && Endorsements.length > 0 ? Endorsements : null,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Beneficiaries: Beneficiaries != null && Beneficiaries.length > 0 ? Beneficiaries : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ReinsuranceAgreements: ReinsuranceAgreements != null && ReinsuranceAgreements.length > 0 ? ReinsuranceAgreements : null,
      		Status: Status,
			PaymentPlan: PaymentPlan
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Policy/update/' + id;
		const obj = {
				      		policyNumber: policyNumber,
      		effectivePeriod: effectivePeriod,
      		totalPremium: totalPremium,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		Agent: Agent != null && Agent.length > 0 ? Agent : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
      		Endorsements: Endorsements != null && Endorsements.length > 0 ? Endorsements : null,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Beneficiaries: Beneficiaries != null && Beneficiaries.length > 0 ? Beneficiaries : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ReinsuranceAgreements: ReinsuranceAgreements != null && ReinsuranceAgreements.length > 0 ? ReinsuranceAgreements : null,
      		Status: Status,
			PaymentPlan: PaymentPlan
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Policy
	// returns the results untouched as an Observable Policy
	// Policy model
	// delegates via URI
	//********************************************************************
	getPolicy(id) : Observable<Policy> {
		const uri_ = this.apiUrl + '/Policy/load/' + id;

		return this.http.get<Policy>(uri_);
	}
	
	//********************************************************************
	// gets all Policy
	// returns the results untouched as JSON representation of an
	// Observable array of Policy models
	// delegates via URI
	//********************************************************************
	getPolicys() : Observable<Policy[]> {
		const uri_ = this.apiUrl + '/Policy/';

		return this
			.http.get<Policy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Insurer on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsurer( policyId, _insurerId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the Insurer from storage
	var tmp 	= new InsurerService(this.http).getInsurer(_insurerId);

	// assign the Insurer
	this.policy.insurer = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Insurer on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsurer( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Insurer to null
	this.policy.insurer = null;

	// save the Policy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Customer on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( policyId, _customerId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.policy.customer = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Customer to null
	this.policy.customer = null;

	// save the Policy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( policyId, _productId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the InsuranceProduct from storage
	var tmp 	= new InsuranceProductService(this.http).getInsuranceProduct(_productId);

	// assign the Product
	this.policy.product = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Product to null
	this.policy.product = null;

	// save the Policy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Agent on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAgent( policyId, _agentId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the Agent from storage
	var tmp 	= new AgentService(this.http).getAgent(_agentId);

	// assign the Agent
	this.policy.agent = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Agent on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAgent( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Agent to null
	this.policy.agent = null;

	// save the Policy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a BillingAccount on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBillingAccount( policyId, _billingAccountId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the BillingAccount from storage
	var tmp 	= new BillingAccountService(this.http).getBillingAccount(_billingAccountId);

	// assign the BillingAccount
	this.policy.billingAccount = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BillingAccount on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBillingAccount( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign BillingAccount to null
	this.policy.billingAccount = null;

	// save the Policy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more coveragesIds as a Coverages
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCoverages( policyId, coveragesIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = coveragesIds.split(',')

	// iterate over array of coverages ids
	idList.forEach(function (id) {
		// read the PolicyCoverage
		var policyCoverage = new PolicyCoverageService(this.http).getPolicyCoverage(id);
		// add the PolicyCoverage if not already assigned
		if ( this.policy.coverages.indexOf(policyCoverage) == -1 )
		this.policy.coverages.push(policyCoverage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more coveragesIds as a Coverages
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCoverages( policyId, coveragesIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= coveragesIds.split(',');
	var coverages 	= this.policy.coverages;

	if ( coverages != null && coveragesIds != null ) {

		// iterate over array of coverages ids
		coverages.forEach(function (obj) {
			if ( coveragesIds.indexOf(obj._id) > -1 ) {
				// remove the PolicyCoverage
				this.policy.coverages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more insuredObjectsIds as a InsuredObjects
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsuredObjects( policyId, insuredObjectsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = insuredObjectsIds.split(',')

	// iterate over array of insuredObjects ids
	idList.forEach(function (id) {
		// read the InsuredObject
		var insuredObject = new InsuredObjectService(this.http).getInsuredObject(id);
		// add the InsuredObject if not already assigned
		if ( this.policy.insuredObjects.indexOf(insuredObject) == -1 )
		this.policy.insuredObjects.push(insuredObject);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insuredObjectsIds as a InsuredObjects
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsuredObjects( policyId, insuredObjectsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= insuredObjectsIds.split(',');
	var insuredObjects 	= this.policy.insuredObjects;

	if ( insuredObjects != null && insuredObjectsIds != null ) {

		// iterate over array of insuredObjects ids
		insuredObjects.forEach(function (obj) {
			if ( insuredObjectsIds.indexOf(obj._id) > -1 ) {
				// remove the InsuredObject
				this.policy.insuredObjects.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more endorsementsIds as a Endorsements
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEndorsements( policyId, endorsementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = endorsementsIds.split(',')

	// iterate over array of endorsements ids
	idList.forEach(function (id) {
		// read the Endorsement
		var endorsement = new EndorsementService(this.http).getEndorsement(id);
		// add the Endorsement if not already assigned
		if ( this.policy.endorsements.indexOf(endorsement) == -1 )
		this.policy.endorsements.push(endorsement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more endorsementsIds as a Endorsements
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEndorsements( policyId, endorsementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= endorsementsIds.split(',');
	var endorsements 	= this.policy.endorsements;

	if ( endorsements != null && endorsementsIds != null ) {

		// iterate over array of endorsements ids
		endorsements.forEach(function (obj) {
			if ( endorsementsIds.indexOf(obj._id) > -1 ) {
				// remove the Endorsement
				this.policy.endorsements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more beneficiariesIds as a Beneficiaries
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBeneficiaries( policyId, beneficiariesIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = beneficiariesIds.split(',')

	// iterate over array of beneficiaries ids
	idList.forEach(function (id) {
		// read the Beneficiary
		var beneficiary = new BeneficiaryService(this.http).getBeneficiary(id);
		// add the Beneficiary if not already assigned
		if ( this.policy.beneficiaries.indexOf(beneficiary) == -1 )
		this.policy.beneficiaries.push(beneficiary);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more beneficiariesIds as a Beneficiaries
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBeneficiaries( policyId, beneficiariesIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= beneficiariesIds.split(',');
	var beneficiaries 	= this.policy.beneficiaries;

	if ( beneficiaries != null && beneficiariesIds != null ) {

		// iterate over array of beneficiaries ids
		beneficiaries.forEach(function (obj) {
			if ( beneficiariesIds.indexOf(obj._id) > -1 ) {
				// remove the Beneficiary
				this.policy.beneficiaries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( policyId, claimsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.policy.claims.indexOf(claim) == -1 )
		this.policy.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( policyId, claimsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.policy.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.policy.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReinsuranceAgreements( policyId, reinsuranceAgreementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = reinsuranceAgreementsIds.split(',')

	// iterate over array of reinsuranceAgreements ids
	idList.forEach(function (id) {
		// read the ReinsuranceAgreement
		var reinsuranceAgreement = new ReinsuranceAgreementService(this.http).getReinsuranceAgreement(id);
		// add the ReinsuranceAgreement if not already assigned
		if ( this.policy.reinsuranceAgreements.indexOf(reinsuranceAgreement) == -1 )
		this.policy.reinsuranceAgreements.push(reinsuranceAgreement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReinsuranceAgreements( policyId, reinsuranceAgreementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= reinsuranceAgreementsIds.split(',');
	var reinsuranceAgreements 	= this.policy.reinsuranceAgreements;

	if ( reinsuranceAgreements != null && reinsuranceAgreementsIds != null ) {

		// iterate over array of reinsuranceAgreements ids
		reinsuranceAgreements.forEach(function (obj) {
			if ( reinsuranceAgreementsIds.indexOf(obj._id) > -1 ) {
				// remove the ReinsuranceAgreement
				this.policy.reinsuranceAgreements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Policy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Policy/update/' + this.policy;

	return  this.http.post(uri_, this.policy );
}

	//********************************************************************
	// loadHelper - internal helper to load a Policy
	//********************************************************************	
	loadHelper( id ) {
		this.getPolicy(id)
			.subscribe((res : Policy) => {
				this.policy = res;
			});
	}
}