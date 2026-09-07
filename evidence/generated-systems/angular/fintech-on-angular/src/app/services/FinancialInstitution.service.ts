import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FinancialInstitution} from '../models/FinancialInstitution';
import {BranchService} from '../services/Branch.service';
import {CustomerService} from '../services/Customer.service';
import {ProductOfferingService} from '../services/ProductOffering.service';
import {PaymentProcessorService} from '../services/PaymentProcessor.service';
import {CompliancePolicyService} from '../services/CompliancePolicy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FinancialInstitutionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	financialInstitution : FinancialInstitution;

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
	// add a FinancialInstitution
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies) : Observable<any> {
		const uri_ = this.apiUrl + '/FinancialInstitution/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		countryOfIncorporation: countryOfIncorporation,
      		bic: bic,
      		website: website,
      		Branches: Branches != null && Branches.length > 0 ? Branches : null,
      		Customers: Customers != null && Customers.length > 0 ? Customers : null,
      		ProductOfferings: ProductOfferings != null && ProductOfferings.length > 0 ? ProductOfferings : null,
      		PaymentProcessors: PaymentProcessors != null && PaymentProcessors.length > 0 ? PaymentProcessors : null,
			CompliancePolicies: CompliancePolicies != null && CompliancePolicies.length > 0 ? CompliancePolicies : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FinancialInstitution
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FinancialInstitution/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		countryOfIncorporation: countryOfIncorporation,
      		bic: bic,
      		website: website,
      		Branches: Branches != null && Branches.length > 0 ? Branches : null,
      		Customers: Customers != null && Customers.length > 0 ? Customers : null,
      		ProductOfferings: ProductOfferings != null && ProductOfferings.length > 0 ? ProductOfferings : null,
      		PaymentProcessors: PaymentProcessors != null && PaymentProcessors.length > 0 ? PaymentProcessors : null,
			CompliancePolicies: CompliancePolicies != null && CompliancePolicies.length > 0 ? CompliancePolicies : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FinancialInstitution
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFinancialInstitution(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FinancialInstitution/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FinancialInstitution
	// returns the results untouched as an Observable FinancialInstitution
	// FinancialInstitution model
	// delegates via URI
	//********************************************************************
	getFinancialInstitution(id) : Observable<FinancialInstitution> {
		const uri_ = this.apiUrl + '/FinancialInstitution/load/' + id;

		return this.http.get<FinancialInstitution>(uri_);
	}
	
	//********************************************************************
	// gets all FinancialInstitution
	// returns the results untouched as JSON representation of an
	// Observable array of FinancialInstitution models
	// delegates via URI
	//********************************************************************
	getFinancialInstitutions() : Observable<FinancialInstitution[]> {
		const uri_ = this.apiUrl + '/FinancialInstitution/';

		return this
			.http.get<FinancialInstitution[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more branchesIds as a Branches
	// to a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBranches( financialInstitutionId, branchesIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );

	// split on a comma with no spaces
	var idList = branchesIds.split(',')

	// iterate over array of branches ids
	idList.forEach(function (id) {
		// read the Branch
		var branch = new BranchService(this.http).getBranch(id);
		// add the Branch if not already assigned
		if ( this.financialInstitution.branches.indexOf(branch) == -1 )
		this.financialInstitution.branches.push(branch);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more branchesIds as a Branches
	// from a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBranches( financialInstitutionId, branchesIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );


	// split on a comma with no spaces
	var idList 					= branchesIds.split(',');
	var branches 	= this.financialInstitution.branches;

	if ( branches != null && branchesIds != null ) {

		// iterate over array of branches ids
		branches.forEach(function (obj) {
			if ( branchesIds.indexOf(obj._id) > -1 ) {
				// remove the Branch
				this.financialInstitution.branches.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more customersIds as a Customers
	// to a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCustomers( financialInstitutionId, customersIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );

	// split on a comma with no spaces
	var idList = customersIds.split(',')

	// iterate over array of customers ids
	idList.forEach(function (id) {
		// read the Customer
		var customer = new CustomerService(this.http).getCustomer(id);
		// add the Customer if not already assigned
		if ( this.financialInstitution.customers.indexOf(customer) == -1 )
		this.financialInstitution.customers.push(customer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more customersIds as a Customers
	// from a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCustomers( financialInstitutionId, customersIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );


	// split on a comma with no spaces
	var idList 					= customersIds.split(',');
	var customers 	= this.financialInstitution.customers;

	if ( customers != null && customersIds != null ) {

		// iterate over array of customers ids
		customers.forEach(function (obj) {
			if ( customersIds.indexOf(obj._id) > -1 ) {
				// remove the Customer
				this.financialInstitution.customers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more productOfferingsIds as a ProductOfferings
	// to a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProductOfferings( financialInstitutionId, productOfferingsIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );

	// split on a comma with no spaces
	var idList = productOfferingsIds.split(',')

	// iterate over array of productOfferings ids
	idList.forEach(function (id) {
		// read the ProductOffering
		var productOffering = new ProductOfferingService(this.http).getProductOffering(id);
		// add the ProductOffering if not already assigned
		if ( this.financialInstitution.productOfferings.indexOf(productOffering) == -1 )
		this.financialInstitution.productOfferings.push(productOffering);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productOfferingsIds as a ProductOfferings
	// from a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProductOfferings( financialInstitutionId, productOfferingsIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );


	// split on a comma with no spaces
	var idList 					= productOfferingsIds.split(',');
	var productOfferings 	= this.financialInstitution.productOfferings;

	if ( productOfferings != null && productOfferingsIds != null ) {

		// iterate over array of productOfferings ids
		productOfferings.forEach(function (obj) {
			if ( productOfferingsIds.indexOf(obj._id) > -1 ) {
				// remove the ProductOffering
				this.financialInstitution.productOfferings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more paymentProcessorsIds as a PaymentProcessors
	// to a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPaymentProcessors( financialInstitutionId, paymentProcessorsIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );

	// split on a comma with no spaces
	var idList = paymentProcessorsIds.split(',')

	// iterate over array of paymentProcessors ids
	idList.forEach(function (id) {
		// read the PaymentProcessor
		var paymentProcessor = new PaymentProcessorService(this.http).getPaymentProcessor(id);
		// add the PaymentProcessor if not already assigned
		if ( this.financialInstitution.paymentProcessors.indexOf(paymentProcessor) == -1 )
		this.financialInstitution.paymentProcessors.push(paymentProcessor);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentProcessorsIds as a PaymentProcessors
	// from a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePaymentProcessors( financialInstitutionId, paymentProcessorsIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );


	// split on a comma with no spaces
	var idList 					= paymentProcessorsIds.split(',');
	var paymentProcessors 	= this.financialInstitution.paymentProcessors;

	if ( paymentProcessors != null && paymentProcessorsIds != null ) {

		// iterate over array of paymentProcessors ids
		paymentProcessors.forEach(function (obj) {
			if ( paymentProcessorsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentProcessor
				this.financialInstitution.paymentProcessors.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more compliancePoliciesIds as a CompliancePolicies
	// to a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompliancePolicies( financialInstitutionId, compliancePoliciesIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );

	// split on a comma with no spaces
	var idList = compliancePoliciesIds.split(',')

	// iterate over array of compliancePolicies ids
	idList.forEach(function (id) {
		// read the CompliancePolicy
		var compliancePolicy = new CompliancePolicyService(this.http).getCompliancePolicy(id);
		// add the CompliancePolicy if not already assigned
		if ( this.financialInstitution.compliancePolicies.indexOf(compliancePolicy) == -1 )
		this.financialInstitution.compliancePolicies.push(compliancePolicy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more compliancePoliciesIds as a CompliancePolicies
	// from a FinancialInstitution
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompliancePolicies( financialInstitutionId, compliancePoliciesIds ): Observable<any> {

		// get the FinancialInstitution
		this.loadHelper( financialInstitutionId );


	// split on a comma with no spaces
	var idList 					= compliancePoliciesIds.split(',');
	var compliancePolicies 	= this.financialInstitution.compliancePolicies;

	if ( compliancePolicies != null && compliancePoliciesIds != null ) {

		// iterate over array of compliancePolicies ids
		compliancePolicies.forEach(function (obj) {
			if ( compliancePoliciesIds.indexOf(obj._id) > -1 ) {
				// remove the CompliancePolicy
				this.financialInstitution.compliancePolicies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a FinancialInstitution
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FinancialInstitution/update/' + this.financialInstitution;

	return  this.http.post(uri_, this.financialInstitution );
}

	//********************************************************************
	// loadHelper - internal helper to load a FinancialInstitution
	//********************************************************************	
	loadHelper( id ) {
		this.getFinancialInstitution(id)
			.subscribe((res : FinancialInstitution) => {
				this.financialInstitution = res;
			});
	}
}