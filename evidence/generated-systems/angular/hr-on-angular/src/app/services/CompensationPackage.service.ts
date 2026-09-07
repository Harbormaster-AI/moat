import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CompensationPackage} from '../models/CompensationPackage';
import {EmploymentContractService} from '../services/EmploymentContract.service';
import {SalaryComponentService} from '../services/SalaryComponent.service';
import {BonusPlanService} from '../services/BonusPlan.service';
import {EquityGrantService} from '../services/EquityGrant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CompensationPackageService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	compensationPackage : CompensationPackage;

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
	// add a CompensationPackage
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants) : Observable<any> {
		const uri_ = this.apiUrl + '/CompensationPackage/create';
		const obj = {
			      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		currency: currency,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
      		SalaryComponents: SalaryComponents != null && SalaryComponents.length > 0 ? SalaryComponents : null,
      		BonusPlans: BonusPlans != null && BonusPlans.length > 0 ? BonusPlans : null,
			EquityGrants: EquityGrants != null && EquityGrants.length > 0 ? EquityGrants : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CompensationPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CompensationPackage/update/' + id;
		const obj = {
				      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		currency: currency,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
      		SalaryComponents: SalaryComponents != null && SalaryComponents.length > 0 ? SalaryComponents : null,
      		BonusPlans: BonusPlans != null && BonusPlans.length > 0 ? BonusPlans : null,
			EquityGrants: EquityGrants != null && EquityGrants.length > 0 ? EquityGrants : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CompensationPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCompensationPackage(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CompensationPackage/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CompensationPackage
	// returns the results untouched as an Observable CompensationPackage
	// CompensationPackage model
	// delegates via URI
	//********************************************************************
	getCompensationPackage(id) : Observable<CompensationPackage> {
		const uri_ = this.apiUrl + '/CompensationPackage/load/' + id;

		return this.http.get<CompensationPackage>(uri_);
	}
	
	//********************************************************************
	// gets all CompensationPackage
	// returns the results untouched as JSON representation of an
	// Observable array of CompensationPackage models
	// delegates via URI
	//********************************************************************
	getCompensationPackages() : Observable<CompensationPackage[]> {
		const uri_ = this.apiUrl + '/CompensationPackage/';

		return this
			.http.get<CompensationPackage[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Contract on a CompensationPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContract( compensationPackageId, _contractId ): Observable<any> {

		// get the CompensationPackage from storage
		this.loadHelper( compensationPackageId );

	// get the EmploymentContract from storage
	var tmp 	= new EmploymentContractService(this.http).getEmploymentContract(_contractId);

	// assign the Contract
	this.compensationPackage.contract = tmp;

	// save the CompensationPackage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contract on a CompensationPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContract( compensationPackageId ): Observable<any> {

		// get the CompensationPackage from storage
		this.loadHelper( compensationPackageId );

	// assign Contract to null
	this.compensationPackage.contract = null;

	// save the CompensationPackage
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more salaryComponentsIds as a SalaryComponents
	// to a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSalaryComponents( compensationPackageId, salaryComponentsIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );

	// split on a comma with no spaces
	var idList = salaryComponentsIds.split(',')

	// iterate over array of salaryComponents ids
	idList.forEach(function (id) {
		// read the SalaryComponent
		var salaryComponent = new SalaryComponentService(this.http).getSalaryComponent(id);
		// add the SalaryComponent if not already assigned
		if ( this.compensationPackage.salaryComponents.indexOf(salaryComponent) == -1 )
		this.compensationPackage.salaryComponents.push(salaryComponent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more salaryComponentsIds as a SalaryComponents
	// from a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSalaryComponents( compensationPackageId, salaryComponentsIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );


	// split on a comma with no spaces
	var idList 					= salaryComponentsIds.split(',');
	var salaryComponents 	= this.compensationPackage.salaryComponents;

	if ( salaryComponents != null && salaryComponentsIds != null ) {

		// iterate over array of salaryComponents ids
		salaryComponents.forEach(function (obj) {
			if ( salaryComponentsIds.indexOf(obj._id) > -1 ) {
				// remove the SalaryComponent
				this.compensationPackage.salaryComponents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more bonusPlansIds as a BonusPlans
	// to a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBonusPlans( compensationPackageId, bonusPlansIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );

	// split on a comma with no spaces
	var idList = bonusPlansIds.split(',')

	// iterate over array of bonusPlans ids
	idList.forEach(function (id) {
		// read the BonusPlan
		var bonusPlan = new BonusPlanService(this.http).getBonusPlan(id);
		// add the BonusPlan if not already assigned
		if ( this.compensationPackage.bonusPlans.indexOf(bonusPlan) == -1 )
		this.compensationPackage.bonusPlans.push(bonusPlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more bonusPlansIds as a BonusPlans
	// from a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBonusPlans( compensationPackageId, bonusPlansIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );


	// split on a comma with no spaces
	var idList 					= bonusPlansIds.split(',');
	var bonusPlans 	= this.compensationPackage.bonusPlans;

	if ( bonusPlans != null && bonusPlansIds != null ) {

		// iterate over array of bonusPlans ids
		bonusPlans.forEach(function (obj) {
			if ( bonusPlansIds.indexOf(obj._id) > -1 ) {
				// remove the BonusPlan
				this.compensationPackage.bonusPlans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more equityGrantsIds as a EquityGrants
	// to a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEquityGrants( compensationPackageId, equityGrantsIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );

	// split on a comma with no spaces
	var idList = equityGrantsIds.split(',')

	// iterate over array of equityGrants ids
	idList.forEach(function (id) {
		// read the EquityGrant
		var equityGrant = new EquityGrantService(this.http).getEquityGrant(id);
		// add the EquityGrant if not already assigned
		if ( this.compensationPackage.equityGrants.indexOf(equityGrant) == -1 )
		this.compensationPackage.equityGrants.push(equityGrant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more equityGrantsIds as a EquityGrants
	// from a CompensationPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEquityGrants( compensationPackageId, equityGrantsIds ): Observable<any> {

		// get the CompensationPackage
		this.loadHelper( compensationPackageId );


	// split on a comma with no spaces
	var idList 					= equityGrantsIds.split(',');
	var equityGrants 	= this.compensationPackage.equityGrants;

	if ( equityGrants != null && equityGrantsIds != null ) {

		// iterate over array of equityGrants ids
		equityGrants.forEach(function (obj) {
			if ( equityGrantsIds.indexOf(obj._id) > -1 ) {
				// remove the EquityGrant
				this.compensationPackage.equityGrants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CompensationPackage
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CompensationPackage/update/' + this.compensationPackage;

	return  this.http.post(uri_, this.compensationPackage );
}

	//********************************************************************
	// loadHelper - internal helper to load a CompensationPackage
	//********************************************************************	
	loadHelper( id ) {
		this.getCompensationPackage(id)
			.subscribe((res : CompensationPackage) => {
				this.compensationPackage = res;
			});
	}
}