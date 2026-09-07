import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SoftwareUpdate} from '../models/SoftwareUpdate';
import {MedicalDeviceService} from '../services/MedicalDevice.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SoftwareUpdateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	softwareUpdate : SoftwareUpdate;

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
	// add a SoftwareUpdate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSoftwareUpdate(version, appliedDate, Device, UpdateType) : Observable<any> {
		const uri_ = this.apiUrl + '/SoftwareUpdate/create';
		const obj = {
			      		version: version,
      		appliedDate: appliedDate,
      		Device: Device != null && Device.length > 0 ? Device : null,
			UpdateType: UpdateType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SoftwareUpdate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSoftwareUpdate(version, appliedDate, Device, UpdateType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SoftwareUpdate/update/' + id;
		const obj = {
				      		version: version,
      		appliedDate: appliedDate,
      		Device: Device != null && Device.length > 0 ? Device : null,
			UpdateType: UpdateType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SoftwareUpdate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSoftwareUpdate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SoftwareUpdate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SoftwareUpdate
	// returns the results untouched as an Observable SoftwareUpdate
	// SoftwareUpdate model
	// delegates via URI
	//********************************************************************
	getSoftwareUpdate(id) : Observable<SoftwareUpdate> {
		const uri_ = this.apiUrl + '/SoftwareUpdate/load/' + id;

		return this.http.get<SoftwareUpdate>(uri_);
	}
	
	//********************************************************************
	// gets all SoftwareUpdate
	// returns the results untouched as JSON representation of an
	// Observable array of SoftwareUpdate models
	// delegates via URI
	//********************************************************************
	getSoftwareUpdates() : Observable<SoftwareUpdate[]> {
		const uri_ = this.apiUrl + '/SoftwareUpdate/';

		return this
			.http.get<SoftwareUpdate[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Device on a SoftwareUpdate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDevice( softwareUpdateId, _deviceId ): Observable<any> {

		// get the SoftwareUpdate from storage
		this.loadHelper( softwareUpdateId );

	// get the MedicalDevice from storage
	var tmp 	= new MedicalDeviceService(this.http).getMedicalDevice(_deviceId);

	// assign the Device
	this.softwareUpdate.device = tmp;

	// save the SoftwareUpdate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Device on a SoftwareUpdate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDevice( softwareUpdateId ): Observable<any> {

		// get the SoftwareUpdate from storage
		this.loadHelper( softwareUpdateId );

	// assign Device to null
	this.softwareUpdate.device = null;

	// save the SoftwareUpdate
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SoftwareUpdate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SoftwareUpdate/update/' + this.softwareUpdate;

	return  this.http.post(uri_, this.softwareUpdate );
}

	//********************************************************************
	// loadHelper - internal helper to load a SoftwareUpdate
	//********************************************************************	
	loadHelper( id ) {
		this.getSoftwareUpdate(id)
			.subscribe((res : SoftwareUpdate) => {
				this.softwareUpdate = res;
			});
	}
}