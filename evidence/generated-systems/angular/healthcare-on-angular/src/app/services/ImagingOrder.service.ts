import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ImagingOrder} from '../models/ImagingOrder';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {ImagingCenterService} from '../services/ImagingCenter.service';
import {ImagingReportService} from '../services/ImagingReport.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ImagingOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	imagingOrder : ImagingOrder;

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
	// add a ImagingOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality) : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingOrder/create';
		const obj = {
			      		bodySite: bodySite,
      		contrast: contrast,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		ImagingCenter: ImagingCenter != null && ImagingCenter.length > 0 ? ImagingCenter : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
			Modality: Modality
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ImagingOrder/update/' + id;
		const obj = {
				      		bodySite: bodySite,
      		contrast: contrast,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		ImagingCenter: ImagingCenter != null && ImagingCenter.length > 0 ? ImagingCenter : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
			Modality: Modality
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteImagingOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ImagingOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ImagingOrder
	// returns the results untouched as an Observable ImagingOrder
	// ImagingOrder model
	// delegates via URI
	//********************************************************************
	getImagingOrder(id) : Observable<ImagingOrder> {
		const uri_ = this.apiUrl + '/ImagingOrder/load/' + id;

		return this.http.get<ImagingOrder>(uri_);
	}
	
	//********************************************************************
	// gets all ImagingOrder
	// returns the results untouched as JSON representation of an
	// Observable array of ImagingOrder models
	// delegates via URI
	//********************************************************************
	getImagingOrders() : Observable<ImagingOrder[]> {
		const uri_ = this.apiUrl + '/ImagingOrder/';

		return this
			.http.get<ImagingOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( imagingOrderId, _orderId ): Observable<any> {

		// get the ImagingOrder from storage
		this.loadHelper( imagingOrderId );

	// get the ClinicalOrder from storage
	var tmp 	= new ClinicalOrderService(this.http).getClinicalOrder(_orderId);

	// assign the Order
	this.imagingOrder.order = tmp;

	// save the ImagingOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( imagingOrderId ): Observable<any> {

		// get the ImagingOrder from storage
		this.loadHelper( imagingOrderId );

	// assign Order to null
	this.imagingOrder.order = null;

	// save the ImagingOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ImagingCenter on a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignImagingCenter( imagingOrderId, _imagingCenterId ): Observable<any> {

		// get the ImagingOrder from storage
		this.loadHelper( imagingOrderId );

	// get the ImagingCenter from storage
	var tmp 	= new ImagingCenterService(this.http).getImagingCenter(_imagingCenterId);

	// assign the ImagingCenter
	this.imagingOrder.imagingCenter = tmp;

	// save the ImagingOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ImagingCenter on a ImagingOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignImagingCenter( imagingOrderId ): Observable<any> {

		// get the ImagingOrder from storage
		this.loadHelper( imagingOrderId );

	// assign ImagingCenter to null
	this.imagingOrder.imagingCenter = null;

	// save the ImagingOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a ImagingOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( imagingOrderId, reportsIds ): Observable<any> {

		// get the ImagingOrder
		this.loadHelper( imagingOrderId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the ImagingReport
		var imagingReport = new ImagingReportService(this.http).getImagingReport(id);
		// add the ImagingReport if not already assigned
		if ( this.imagingOrder.reports.indexOf(imagingReport) == -1 )
		this.imagingOrder.reports.push(imagingReport);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a ImagingOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( imagingOrderId, reportsIds ): Observable<any> {

		// get the ImagingOrder
		this.loadHelper( imagingOrderId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.imagingOrder.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingReport
				this.imagingOrder.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ImagingOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ImagingOrder/update/' + this.imagingOrder;

	return  this.http.post(uri_, this.imagingOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a ImagingOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getImagingOrder(id)
			.subscribe((res : ImagingOrder) => {
				this.imagingOrder = res;
			});
	}
}