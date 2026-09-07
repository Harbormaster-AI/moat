import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DeviceCriterion} from '../models/DeviceCriterion';
import {TargetingProfileService} from '../services/TargetingProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DeviceCriterionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	deviceCriterion : DeviceCriterion;

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
	// add a DeviceCriterion
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator) : Observable<any> {
		const uri_ = this.apiUrl + '/DeviceCriterion/create';
		const obj = {
			      		TargetingProfile: TargetingProfile != null && TargetingProfile.length > 0 ? TargetingProfile : null,
      		DeviceType: DeviceType,
      		PlatformType: PlatformType,
			Operator: Operator
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DeviceCriterion
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DeviceCriterion/update/' + id;
		const obj = {
				      		TargetingProfile: TargetingProfile != null && TargetingProfile.length > 0 ? TargetingProfile : null,
      		DeviceType: DeviceType,
      		PlatformType: PlatformType,
			Operator: Operator
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DeviceCriterion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDeviceCriterion(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DeviceCriterion/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DeviceCriterion
	// returns the results untouched as an Observable DeviceCriterion
	// DeviceCriterion model
	// delegates via URI
	//********************************************************************
	getDeviceCriterion(id) : Observable<DeviceCriterion> {
		const uri_ = this.apiUrl + '/DeviceCriterion/load/' + id;

		return this.http.get<DeviceCriterion>(uri_);
	}
	
	//********************************************************************
	// gets all DeviceCriterion
	// returns the results untouched as JSON representation of an
	// Observable array of DeviceCriterion models
	// delegates via URI
	//********************************************************************
	getDeviceCriterions() : Observable<DeviceCriterion[]> {
		const uri_ = this.apiUrl + '/DeviceCriterion/';

		return this
			.http.get<DeviceCriterion[]>(uri_);
	}
	
			//********************************************************************
	// assigns a TargetingProfile on a DeviceCriterion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTargetingProfile( deviceCriterionId, _targetingProfileId ): Observable<any> {

		// get the DeviceCriterion from storage
		this.loadHelper( deviceCriterionId );

	// get the TargetingProfile from storage
	var tmp 	= new TargetingProfileService(this.http).getTargetingProfile(_targetingProfileId);

	// assign the TargetingProfile
	this.deviceCriterion.targetingProfile = tmp;

	// save the DeviceCriterion
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TargetingProfile on a DeviceCriterion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTargetingProfile( deviceCriterionId ): Observable<any> {

		// get the DeviceCriterion from storage
		this.loadHelper( deviceCriterionId );

	// assign TargetingProfile to null
	this.deviceCriterion.targetingProfile = null;

	// save the DeviceCriterion
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a DeviceCriterion
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DeviceCriterion/update/' + this.deviceCriterion;

	return  this.http.post(uri_, this.deviceCriterion );
}

	//********************************************************************
	// loadHelper - internal helper to load a DeviceCriterion
	//********************************************************************	
	loadHelper( id ) {
		this.getDeviceCriterion(id)
			.subscribe((res : DeviceCriterion) => {
				this.deviceCriterion = res;
			});
	}
}