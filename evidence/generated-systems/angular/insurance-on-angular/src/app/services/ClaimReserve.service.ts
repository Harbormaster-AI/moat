import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ClaimReserve} from '../models/ClaimReserve';
import {ClaimService} from '../services/Claim.service';
import {ExposureService} from '../services/Exposure.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClaimReserveService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	claimReserve : ClaimReserve;

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
	// add a ClaimReserve
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ClaimReserve/create';
		const obj = {
			      		amount: amount,
      		setDate: setDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		ReserveType: ReserveType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ClaimReserve/update/' + id;
		const obj = {
				      		amount: amount,
      		setDate: setDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		ReserveType: ReserveType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClaimReserve(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ClaimReserve/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ClaimReserve
	// returns the results untouched as an Observable ClaimReserve
	// ClaimReserve model
	// delegates via URI
	//********************************************************************
	getClaimReserve(id) : Observable<ClaimReserve> {
		const uri_ = this.apiUrl + '/ClaimReserve/load/' + id;

		return this.http.get<ClaimReserve>(uri_);
	}
	
	//********************************************************************
	// gets all ClaimReserve
	// returns the results untouched as JSON representation of an
	// Observable array of ClaimReserve models
	// delegates via URI
	//********************************************************************
	getClaimReserves() : Observable<ClaimReserve[]> {
		const uri_ = this.apiUrl + '/ClaimReserve/';

		return this
			.http.get<ClaimReserve[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Claim on a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( claimReserveId, _claimId ): Observable<any> {

		// get the ClaimReserve from storage
		this.loadHelper( claimReserveId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.claimReserve.claim = tmp;

	// save the ClaimReserve
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( claimReserveId ): Observable<any> {

		// get the ClaimReserve from storage
		this.loadHelper( claimReserveId );

	// assign Claim to null
	this.claimReserve.claim = null;

	// save the ClaimReserve
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Exposure on a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignExposure( claimReserveId, _exposureId ): Observable<any> {

		// get the ClaimReserve from storage
		this.loadHelper( claimReserveId );

	// get the Exposure from storage
	var tmp 	= new ExposureService(this.http).getExposure(_exposureId);

	// assign the Exposure
	this.claimReserve.exposure = tmp;

	// save the ClaimReserve
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Exposure on a ClaimReserve
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignExposure( claimReserveId ): Observable<any> {

		// get the ClaimReserve from storage
		this.loadHelper( claimReserveId );

	// assign Exposure to null
	this.claimReserve.exposure = null;

	// save the ClaimReserve
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ClaimReserve
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ClaimReserve/update/' + this.claimReserve;

	return  this.http.post(uri_, this.claimReserve );
}

	//********************************************************************
	// loadHelper - internal helper to load a ClaimReserve
	//********************************************************************	
	loadHelper( id ) {
		this.getClaimReserve(id)
			.subscribe((res : ClaimReserve) => {
				this.claimReserve = res;
			});
	}
}