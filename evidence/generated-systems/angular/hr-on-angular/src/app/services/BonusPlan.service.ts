import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BonusPlan} from '../models/BonusPlan';
import {CompensationPackageService} from '../services/CompensationPackage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BonusPlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	bonusPlan : BonusPlan;

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
	// add a BonusPlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBonusPlan(name, targetPercentage, CompensationPackages) : Observable<any> {
		const uri_ = this.apiUrl + '/BonusPlan/create';
		const obj = {
			      		name: name,
      		targetPercentage: targetPercentage,
			CompensationPackages: CompensationPackages != null && CompensationPackages.length > 0 ? CompensationPackages : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BonusPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBonusPlan(name, targetPercentage, CompensationPackages, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BonusPlan/update/' + id;
		const obj = {
				      		name: name,
      		targetPercentage: targetPercentage,
			CompensationPackages: CompensationPackages != null && CompensationPackages.length > 0 ? CompensationPackages : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BonusPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBonusPlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BonusPlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BonusPlan
	// returns the results untouched as an Observable BonusPlan
	// BonusPlan model
	// delegates via URI
	//********************************************************************
	getBonusPlan(id) : Observable<BonusPlan> {
		const uri_ = this.apiUrl + '/BonusPlan/load/' + id;

		return this.http.get<BonusPlan>(uri_);
	}
	
	//********************************************************************
	// gets all BonusPlan
	// returns the results untouched as JSON representation of an
	// Observable array of BonusPlan models
	// delegates via URI
	//********************************************************************
	getBonusPlans() : Observable<BonusPlan[]> {
		const uri_ = this.apiUrl + '/BonusPlan/';

		return this
			.http.get<BonusPlan[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more compensationPackagesIds as a CompensationPackages
	// to a BonusPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompensationPackages( bonusPlanId, compensationPackagesIds ): Observable<any> {

		// get the BonusPlan
		this.loadHelper( bonusPlanId );

	// split on a comma with no spaces
	var idList = compensationPackagesIds.split(',')

	// iterate over array of compensationPackages ids
	idList.forEach(function (id) {
		// read the CompensationPackage
		var compensationPackage = new CompensationPackageService(this.http).getCompensationPackage(id);
		// add the CompensationPackage if not already assigned
		if ( this.bonusPlan.compensationPackages.indexOf(compensationPackage) == -1 )
		this.bonusPlan.compensationPackages.push(compensationPackage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more compensationPackagesIds as a CompensationPackages
	// from a BonusPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompensationPackages( bonusPlanId, compensationPackagesIds ): Observable<any> {

		// get the BonusPlan
		this.loadHelper( bonusPlanId );


	// split on a comma with no spaces
	var idList 					= compensationPackagesIds.split(',');
	var compensationPackages 	= this.bonusPlan.compensationPackages;

	if ( compensationPackages != null && compensationPackagesIds != null ) {

		// iterate over array of compensationPackages ids
		compensationPackages.forEach(function (obj) {
			if ( compensationPackagesIds.indexOf(obj._id) > -1 ) {
				// remove the CompensationPackage
				this.bonusPlan.compensationPackages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BonusPlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BonusPlan/update/' + this.bonusPlan;

	return  this.http.post(uri_, this.bonusPlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a BonusPlan
	//********************************************************************	
	loadHelper( id ) {
		this.getBonusPlan(id)
			.subscribe((res : BonusPlan) => {
				this.bonusPlan = res;
			});
	}
}