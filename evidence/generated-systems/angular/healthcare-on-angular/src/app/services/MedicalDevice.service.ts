import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MedicalDevice} from '../models/MedicalDevice';
import {PatientService} from '../services/Patient.service';
import {ObservationService} from '../services/Observation.service';
import {SoftwareUpdateService} from '../services/SoftwareUpdate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MedicalDeviceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	medicalDevice : MedicalDevice;

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
	// add a MedicalDevice
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/MedicalDevice/create';
		const obj = {
			      		udi: udi,
      		manufacturer: manufacturer,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		SoftwareUpdates: SoftwareUpdates != null && SoftwareUpdates.length > 0 ? SoftwareUpdates : null,
      		DeviceType: DeviceType,
			ConnectivityStatus: ConnectivityStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MedicalDevice
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MedicalDevice/update/' + id;
		const obj = {
				      		udi: udi,
      		manufacturer: manufacturer,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		SoftwareUpdates: SoftwareUpdates != null && SoftwareUpdates.length > 0 ? SoftwareUpdates : null,
      		DeviceType: DeviceType,
			ConnectivityStatus: ConnectivityStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MedicalDevice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMedicalDevice(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MedicalDevice/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MedicalDevice
	// returns the results untouched as an Observable MedicalDevice
	// MedicalDevice model
	// delegates via URI
	//********************************************************************
	getMedicalDevice(id) : Observable<MedicalDevice> {
		const uri_ = this.apiUrl + '/MedicalDevice/load/' + id;

		return this.http.get<MedicalDevice>(uri_);
	}
	
	//********************************************************************
	// gets all MedicalDevice
	// returns the results untouched as JSON representation of an
	// Observable array of MedicalDevice models
	// delegates via URI
	//********************************************************************
	getMedicalDevices() : Observable<MedicalDevice[]> {
		const uri_ = this.apiUrl + '/MedicalDevice/';

		return this
			.http.get<MedicalDevice[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a MedicalDevice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( medicalDeviceId, _patientId ): Observable<any> {

		// get the MedicalDevice from storage
		this.loadHelper( medicalDeviceId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.medicalDevice.patient = tmp;

	// save the MedicalDevice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a MedicalDevice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( medicalDeviceId ): Observable<any> {

		// get the MedicalDevice from storage
		this.loadHelper( medicalDeviceId );

	// assign Patient to null
	this.medicalDevice.patient = null;

	// save the MedicalDevice
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more observationsIds as a Observations
	// to a MedicalDevice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObservations( medicalDeviceId, observationsIds ): Observable<any> {

		// get the MedicalDevice
		this.loadHelper( medicalDeviceId );

	// split on a comma with no spaces
	var idList = observationsIds.split(',')

	// iterate over array of observations ids
	idList.forEach(function (id) {
		// read the Observation
		var observation = new ObservationService(this.http).getObservation(id);
		// add the Observation if not already assigned
		if ( this.medicalDevice.observations.indexOf(observation) == -1 )
		this.medicalDevice.observations.push(observation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more observationsIds as a Observations
	// from a MedicalDevice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObservations( medicalDeviceId, observationsIds ): Observable<any> {

		// get the MedicalDevice
		this.loadHelper( medicalDeviceId );


	// split on a comma with no spaces
	var idList 					= observationsIds.split(',');
	var observations 	= this.medicalDevice.observations;

	if ( observations != null && observationsIds != null ) {

		// iterate over array of observations ids
		observations.forEach(function (obj) {
			if ( observationsIds.indexOf(obj._id) > -1 ) {
				// remove the Observation
				this.medicalDevice.observations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more softwareUpdatesIds as a SoftwareUpdates
	// to a MedicalDevice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSoftwareUpdates( medicalDeviceId, softwareUpdatesIds ): Observable<any> {

		// get the MedicalDevice
		this.loadHelper( medicalDeviceId );

	// split on a comma with no spaces
	var idList = softwareUpdatesIds.split(',')

	// iterate over array of softwareUpdates ids
	idList.forEach(function (id) {
		// read the SoftwareUpdate
		var softwareUpdate = new SoftwareUpdateService(this.http).getSoftwareUpdate(id);
		// add the SoftwareUpdate if not already assigned
		if ( this.medicalDevice.softwareUpdates.indexOf(softwareUpdate) == -1 )
		this.medicalDevice.softwareUpdates.push(softwareUpdate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more softwareUpdatesIds as a SoftwareUpdates
	// from a MedicalDevice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSoftwareUpdates( medicalDeviceId, softwareUpdatesIds ): Observable<any> {

		// get the MedicalDevice
		this.loadHelper( medicalDeviceId );


	// split on a comma with no spaces
	var idList 					= softwareUpdatesIds.split(',');
	var softwareUpdates 	= this.medicalDevice.softwareUpdates;

	if ( softwareUpdates != null && softwareUpdatesIds != null ) {

		// iterate over array of softwareUpdates ids
		softwareUpdates.forEach(function (obj) {
			if ( softwareUpdatesIds.indexOf(obj._id) > -1 ) {
				// remove the SoftwareUpdate
				this.medicalDevice.softwareUpdates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MedicalDevice
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MedicalDevice/update/' + this.medicalDevice;

	return  this.http.post(uri_, this.medicalDevice );
}

	//********************************************************************
	// loadHelper - internal helper to load a MedicalDevice
	//********************************************************************	
	loadHelper( id ) {
		this.getMedicalDevice(id)
			.subscribe((res : MedicalDevice) => {
				this.medicalDevice = res;
			});
	}
}