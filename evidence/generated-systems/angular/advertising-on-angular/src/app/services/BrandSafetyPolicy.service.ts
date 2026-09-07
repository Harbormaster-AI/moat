import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BrandSafetyPolicy} from '../models/BrandSafetyPolicy';
import {TargetingProfileService} from '../services/TargetingProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BrandSafetyPolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	brandSafetyPolicy : BrandSafetyPolicy;

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
	// add a BrandSafetyPolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold) : Observable<any> {
		const uri_ = this.apiUrl + '/BrandSafetyPolicy/create';
		const obj = {
			      		TargetingProfiles: TargetingProfiles != null && TargetingProfiles.length > 0 ? TargetingProfiles : null,
      		Level: Level,
			ContentRatingThreshold: ContentRatingThreshold
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BrandSafetyPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BrandSafetyPolicy/update/' + id;
		const obj = {
				      		TargetingProfiles: TargetingProfiles != null && TargetingProfiles.length > 0 ? TargetingProfiles : null,
      		Level: Level,
			ContentRatingThreshold: ContentRatingThreshold
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BrandSafetyPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBrandSafetyPolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BrandSafetyPolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BrandSafetyPolicy
	// returns the results untouched as an Observable BrandSafetyPolicy
	// BrandSafetyPolicy model
	// delegates via URI
	//********************************************************************
	getBrandSafetyPolicy(id) : Observable<BrandSafetyPolicy> {
		const uri_ = this.apiUrl + '/BrandSafetyPolicy/load/' + id;

		return this.http.get<BrandSafetyPolicy>(uri_);
	}
	
	//********************************************************************
	// gets all BrandSafetyPolicy
	// returns the results untouched as JSON representation of an
	// Observable array of BrandSafetyPolicy models
	// delegates via URI
	//********************************************************************
	getBrandSafetyPolicys() : Observable<BrandSafetyPolicy[]> {
		const uri_ = this.apiUrl + '/BrandSafetyPolicy/';

		return this
			.http.get<BrandSafetyPolicy[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more targetingProfilesIds as a TargetingProfiles
	// to a BrandSafetyPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTargetingProfiles( brandSafetyPolicyId, targetingProfilesIds ): Observable<any> {

		// get the BrandSafetyPolicy
		this.loadHelper( brandSafetyPolicyId );

	// split on a comma with no spaces
	var idList = targetingProfilesIds.split(',')

	// iterate over array of targetingProfiles ids
	idList.forEach(function (id) {
		// read the TargetingProfile
		var targetingProfile = new TargetingProfileService(this.http).getTargetingProfile(id);
		// add the TargetingProfile if not already assigned
		if ( this.brandSafetyPolicy.targetingProfiles.indexOf(targetingProfile) == -1 )
		this.brandSafetyPolicy.targetingProfiles.push(targetingProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more targetingProfilesIds as a TargetingProfiles
	// from a BrandSafetyPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTargetingProfiles( brandSafetyPolicyId, targetingProfilesIds ): Observable<any> {

		// get the BrandSafetyPolicy
		this.loadHelper( brandSafetyPolicyId );


	// split on a comma with no spaces
	var idList 					= targetingProfilesIds.split(',');
	var targetingProfiles 	= this.brandSafetyPolicy.targetingProfiles;

	if ( targetingProfiles != null && targetingProfilesIds != null ) {

		// iterate over array of targetingProfiles ids
		targetingProfiles.forEach(function (obj) {
			if ( targetingProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the TargetingProfile
				this.brandSafetyPolicy.targetingProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BrandSafetyPolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BrandSafetyPolicy/update/' + this.brandSafetyPolicy;

	return  this.http.post(uri_, this.brandSafetyPolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a BrandSafetyPolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getBrandSafetyPolicy(id)
			.subscribe((res : BrandSafetyPolicy) => {
				this.brandSafetyPolicy = res;
			});
	}
}