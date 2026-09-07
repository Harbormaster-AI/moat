import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Screening} from '../models/Screening';
import {KYCProfileService} from '../services/KYCProfile.service';
import {ComplianceAlertService} from '../services/ComplianceAlert.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ScreeningService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	screening : Screening;

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
	// add a Screening
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addScreening(score, screenedAt, KycProfile, Alerts, ScreeningType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Screening/create';
		const obj = {
			      		score: score,
      		screenedAt: screenedAt,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		ScreeningType: ScreeningType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateScreening(score, screenedAt, KycProfile, Alerts, ScreeningType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Screening/update/' + id;
		const obj = {
				      		score: score,
      		screenedAt: screenedAt,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		ScreeningType: ScreeningType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteScreening(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Screening/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Screening
	// returns the results untouched as an Observable Screening
	// Screening model
	// delegates via URI
	//********************************************************************
	getScreening(id) : Observable<Screening> {
		const uri_ = this.apiUrl + '/Screening/load/' + id;

		return this.http.get<Screening>(uri_);
	}
	
	//********************************************************************
	// gets all Screening
	// returns the results untouched as JSON representation of an
	// Observable array of Screening models
	// delegates via URI
	//********************************************************************
	getScreenings() : Observable<Screening[]> {
		const uri_ = this.apiUrl + '/Screening/';

		return this
			.http.get<Screening[]>(uri_);
	}
	
			//********************************************************************
	// assigns a KycProfile on a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignKycProfile( screeningId, _kycProfileId ): Observable<any> {

		// get the Screening from storage
		this.loadHelper( screeningId );

	// get the KYCProfile from storage
	var tmp 	= new KYCProfileService(this.http).getKYCProfile(_kycProfileId);

	// assign the KycProfile
	this.screening.kycProfile = tmp;

	// save the Screening
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a KycProfile on a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignKycProfile( screeningId ): Observable<any> {

		// get the Screening from storage
		this.loadHelper( screeningId );

	// assign KycProfile to null
	this.screening.kycProfile = null;

	// save the Screening
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a Screening
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( screeningId, alertsIds ): Observable<any> {

		// get the Screening
		this.loadHelper( screeningId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the ComplianceAlert
		var complianceAlert = new ComplianceAlertService(this.http).getComplianceAlert(id);
		// add the ComplianceAlert if not already assigned
		if ( this.screening.alerts.indexOf(complianceAlert) == -1 )
		this.screening.alerts.push(complianceAlert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a Screening
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( screeningId, alertsIds ): Observable<any> {

		// get the Screening
		this.loadHelper( screeningId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.screening.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceAlert
				this.screening.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Screening
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Screening/update/' + this.screening;

	return  this.http.post(uri_, this.screening );
}

	//********************************************************************
	// loadHelper - internal helper to load a Screening
	//********************************************************************	
	loadHelper( id ) {
		this.getScreening(id)
			.subscribe((res : Screening) => {
				this.screening = res;
			});
	}
}