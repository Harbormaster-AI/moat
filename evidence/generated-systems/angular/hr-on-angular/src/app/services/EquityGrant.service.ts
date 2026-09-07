import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EquityGrant} from '../models/EquityGrant';
import {CompensationPackageService} from '../services/CompensationPackage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EquityGrantService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	equityGrant : EquityGrant;

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
	// add a EquityGrant
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType) : Observable<any> {
		const uri_ = this.apiUrl + '/EquityGrant/create';
		const obj = {
			      		grantId: grantId,
      		grantedUnits: grantedUnits,
      		vestingStart: vestingStart,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
			GrantType: GrantType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EquityGrant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EquityGrant/update/' + id;
		const obj = {
				      		grantId: grantId,
      		grantedUnits: grantedUnits,
      		vestingStart: vestingStart,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
			GrantType: GrantType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EquityGrant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEquityGrant(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EquityGrant/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EquityGrant
	// returns the results untouched as an Observable EquityGrant
	// EquityGrant model
	// delegates via URI
	//********************************************************************
	getEquityGrant(id) : Observable<EquityGrant> {
		const uri_ = this.apiUrl + '/EquityGrant/load/' + id;

		return this.http.get<EquityGrant>(uri_);
	}
	
	//********************************************************************
	// gets all EquityGrant
	// returns the results untouched as JSON representation of an
	// Observable array of EquityGrant models
	// delegates via URI
	//********************************************************************
	getEquityGrants() : Observable<EquityGrant[]> {
		const uri_ = this.apiUrl + '/EquityGrant/';

		return this
			.http.get<EquityGrant[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CompensationPackage on a EquityGrant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCompensationPackage( equityGrantId, _compensationPackageId ): Observable<any> {

		// get the EquityGrant from storage
		this.loadHelper( equityGrantId );

	// get the CompensationPackage from storage
	var tmp 	= new CompensationPackageService(this.http).getCompensationPackage(_compensationPackageId);

	// assign the CompensationPackage
	this.equityGrant.compensationPackage = tmp;

	// save the EquityGrant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CompensationPackage on a EquityGrant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCompensationPackage( equityGrantId ): Observable<any> {

		// get the EquityGrant from storage
		this.loadHelper( equityGrantId );

	// assign CompensationPackage to null
	this.equityGrant.compensationPackage = null;

	// save the EquityGrant
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a EquityGrant
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EquityGrant/update/' + this.equityGrant;

	return  this.http.post(uri_, this.equityGrant );
}

	//********************************************************************
	// loadHelper - internal helper to load a EquityGrant
	//********************************************************************	
	loadHelper( id ) {
		this.getEquityGrant(id)
			.subscribe((res : EquityGrant) => {
				this.equityGrant = res;
			});
	}
}