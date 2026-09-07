import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ImagingCenter} from '../models/ImagingCenter';
import {FacilityService} from '../services/Facility.service';
import {ImagingOrderService} from '../services/ImagingOrder.service';
import {ImagingReportService} from '../services/ImagingReport.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ImagingCenterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	imagingCenter : ImagingCenter;

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
	// add a ImagingCenter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addImagingCenter(name, Facility, ImagingOrders, ImagingReports) : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingCenter/create';
		const obj = {
			      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
			ImagingReports: ImagingReports != null && ImagingReports.length > 0 ? ImagingReports : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ImagingCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateImagingCenter(name, Facility, ImagingOrders, ImagingReports, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ImagingCenter/update/' + id;
		const obj = {
				      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
			ImagingReports: ImagingReports != null && ImagingReports.length > 0 ? ImagingReports : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ImagingCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteImagingCenter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingCenter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ImagingCenter
	// returns the results untouched as an Observable ImagingCenter
	// ImagingCenter model
	// delegates via URI
	//********************************************************************
	getImagingCenter(id) : Observable<ImagingCenter> {
		const uri_ = this.apiUrl + '/ImagingCenter/load/' + id;

		return this.http.get<ImagingCenter>(uri_);
	}
	
	//********************************************************************
	// gets all ImagingCenter
	// returns the results untouched as JSON representation of an
	// Observable array of ImagingCenter models
	// delegates via URI
	//********************************************************************
	getImagingCenters() : Observable<ImagingCenter[]> {
		const uri_ = this.apiUrl + '/ImagingCenter/';

		return this
			.http.get<ImagingCenter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Facility on a ImagingCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( imagingCenterId, _facilityId ): Observable<any> {

		// get the ImagingCenter from storage
		this.loadHelper( imagingCenterId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.imagingCenter.facility = tmp;

	// save the ImagingCenter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a ImagingCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( imagingCenterId ): Observable<any> {

		// get the ImagingCenter from storage
		this.loadHelper( imagingCenterId );

	// assign Facility to null
	this.imagingCenter.facility = null;

	// save the ImagingCenter
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more imagingOrdersIds as a ImagingOrders
	// to a ImagingCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingOrders( imagingCenterId, imagingOrdersIds ): Observable<any> {

		// get the ImagingCenter
		this.loadHelper( imagingCenterId );

	// split on a comma with no spaces
	var idList = imagingOrdersIds.split(',')

	// iterate over array of imagingOrders ids
	idList.forEach(function (id) {
		// read the ImagingOrder
		var imagingOrder = new ImagingOrderService(this.http).getImagingOrder(id);
		// add the ImagingOrder if not already assigned
		if ( this.imagingCenter.imagingOrders.indexOf(imagingOrder) == -1 )
		this.imagingCenter.imagingOrders.push(imagingOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingOrdersIds as a ImagingOrders
	// from a ImagingCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingOrders( imagingCenterId, imagingOrdersIds ): Observable<any> {

		// get the ImagingCenter
		this.loadHelper( imagingCenterId );


	// split on a comma with no spaces
	var idList 					= imagingOrdersIds.split(',');
	var imagingOrders 	= this.imagingCenter.imagingOrders;

	if ( imagingOrders != null && imagingOrdersIds != null ) {

		// iterate over array of imagingOrders ids
		imagingOrders.forEach(function (obj) {
			if ( imagingOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingOrder
				this.imagingCenter.imagingOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more imagingReportsIds as a ImagingReports
	// to a ImagingCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingReports( imagingCenterId, imagingReportsIds ): Observable<any> {

		// get the ImagingCenter
		this.loadHelper( imagingCenterId );

	// split on a comma with no spaces
	var idList = imagingReportsIds.split(',')

	// iterate over array of imagingReports ids
	idList.forEach(function (id) {
		// read the ImagingReport
		var imagingReport = new ImagingReportService(this.http).getImagingReport(id);
		// add the ImagingReport if not already assigned
		if ( this.imagingCenter.imagingReports.indexOf(imagingReport) == -1 )
		this.imagingCenter.imagingReports.push(imagingReport);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingReportsIds as a ImagingReports
	// from a ImagingCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingReports( imagingCenterId, imagingReportsIds ): Observable<any> {

		// get the ImagingCenter
		this.loadHelper( imagingCenterId );


	// split on a comma with no spaces
	var idList 					= imagingReportsIds.split(',');
	var imagingReports 	= this.imagingCenter.imagingReports;

	if ( imagingReports != null && imagingReportsIds != null ) {

		// iterate over array of imagingReports ids
		imagingReports.forEach(function (obj) {
			if ( imagingReportsIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingReport
				this.imagingCenter.imagingReports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ImagingCenter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ImagingCenter/update/' + this.imagingCenter;

	return  this.http.post(uri_, this.imagingCenter );
}

	//********************************************************************
	// loadHelper - internal helper to load a ImagingCenter
	//********************************************************************	
	loadHelper( id ) {
		this.getImagingCenter(id)
			.subscribe((res : ImagingCenter) => {
				this.imagingCenter = res;
			});
	}
}