import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ImagingReport} from '../models/ImagingReport';
import {ImagingOrderService} from '../services/ImagingOrder.service';
import {ClinicianService} from '../services/Clinician.service';
import {EncounterService} from '../services/Encounter.service';
import {ImagingCenterService} from '../services/ImagingCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ImagingReportService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	imagingReport : ImagingReport;

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
	// add a ImagingReport
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingReport/create';
		const obj = {
			      		reportNumber: reportNumber,
      		impression: impression,
      		reportedDate: reportedDate,
      		ImagingOrder: ImagingOrder != null && ImagingOrder.length > 0 ? ImagingOrder : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		ImagingCenter: ImagingCenter != null && ImagingCenter.length > 0 ? ImagingCenter : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ImagingReport/update/' + id;
		const obj = {
				      		reportNumber: reportNumber,
      		impression: impression,
      		reportedDate: reportedDate,
      		ImagingOrder: ImagingOrder != null && ImagingOrder.length > 0 ? ImagingOrder : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		ImagingCenter: ImagingCenter != null && ImagingCenter.length > 0 ? ImagingCenter : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteImagingReport(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingReport/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ImagingReport
	// returns the results untouched as an Observable ImagingReport
	// ImagingReport model
	// delegates via URI
	//********************************************************************
	getImagingReport(id) : Observable<ImagingReport> {
		const uri_ = this.apiUrl + '/ImagingReport/load/' + id;

		return this.http.get<ImagingReport>(uri_);
	}
	
	//********************************************************************
	// gets all ImagingReport
	// returns the results untouched as JSON representation of an
	// Observable array of ImagingReport models
	// delegates via URI
	//********************************************************************
	getImagingReports() : Observable<ImagingReport[]> {
		const uri_ = this.apiUrl + '/ImagingReport/';

		return this
			.http.get<ImagingReport[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ImagingOrder on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignImagingOrder( imagingReportId, _imagingOrderId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// get the ImagingOrder from storage
	var tmp 	= new ImagingOrderService(this.http).getImagingOrder(_imagingOrderId);

	// assign the ImagingOrder
	this.imagingReport.imagingOrder = tmp;

	// save the ImagingReport
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ImagingOrder on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignImagingOrder( imagingReportId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// assign ImagingOrder to null
	this.imagingReport.imagingOrder = null;

	// save the ImagingReport
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Clinician on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClinician( imagingReportId, _clinicianId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_clinicianId);

	// assign the Clinician
	this.imagingReport.clinician = tmp;

	// save the ImagingReport
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Clinician on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClinician( imagingReportId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// assign Clinician to null
	this.imagingReport.clinician = null;

	// save the ImagingReport
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Encounter on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( imagingReportId, _encounterId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.imagingReport.encounter = tmp;

	// save the ImagingReport
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( imagingReportId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// assign Encounter to null
	this.imagingReport.encounter = null;

	// save the ImagingReport
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ImagingCenter on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignImagingCenter( imagingReportId, _imagingCenterId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// get the ImagingCenter from storage
	var tmp 	= new ImagingCenterService(this.http).getImagingCenter(_imagingCenterId);

	// assign the ImagingCenter
	this.imagingReport.imagingCenter = tmp;

	// save the ImagingReport
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ImagingCenter on a ImagingReport
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignImagingCenter( imagingReportId ): Observable<any> {

		// get the ImagingReport from storage
		this.loadHelper( imagingReportId );

	// assign ImagingCenter to null
	this.imagingReport.imagingCenter = null;

	// save the ImagingReport
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ImagingReport
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ImagingReport/update/' + this.imagingReport;

	return  this.http.post(uri_, this.imagingReport );
}

	//********************************************************************
	// loadHelper - internal helper to load a ImagingReport
	//********************************************************************	
	loadHelper( id ) {
		this.getImagingReport(id)
			.subscribe((res : ImagingReport) => {
				this.imagingReport = res;
			});
	}
}