import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {KPI} from '../models/KPI';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class KPIService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	kPI : KPI;

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
	// add a KPI
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addKPI(targetValue, Campaign, MetricType) : Observable<any> {
		const uri_ = this.apiUrl + '/KPI/create';
		const obj = {
			      		targetValue: targetValue,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
			MetricType: MetricType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a KPI
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateKPI(targetValue, Campaign, MetricType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/KPI/update/' + id;
		const obj = {
				      		targetValue: targetValue,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
			MetricType: MetricType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a KPI
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteKPI(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/KPI/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a KPI
	// returns the results untouched as an Observable KPI
	// KPI model
	// delegates via URI
	//********************************************************************
	getKPI(id) : Observable<KPI> {
		const uri_ = this.apiUrl + '/KPI/load/' + id;

		return this.http.get<KPI>(uri_);
	}
	
	//********************************************************************
	// gets all KPI
	// returns the results untouched as JSON representation of an
	// Observable array of KPI models
	// delegates via URI
	//********************************************************************
	getKPIs() : Observable<KPI[]> {
		const uri_ = this.apiUrl + '/KPI/';

		return this
			.http.get<KPI[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a KPI
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( kPIId, _campaignId ): Observable<any> {

		// get the KPI from storage
		this.loadHelper( kPIId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.kPI.campaign = tmp;

	// save the KPI
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a KPI
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( kPIId ): Observable<any> {

		// get the KPI from storage
		this.loadHelper( kPIId );

	// assign Campaign to null
	this.kPI.campaign = null;

	// save the KPI
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a KPI
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/KPI/update/' + this.kPI;

	return  this.http.post(uri_, this.kPI );
}

	//********************************************************************
	// loadHelper - internal helper to load a KPI
	//********************************************************************	
	loadHelper( id ) {
		this.getKPI(id)
			.subscribe((res : KPI) => {
				this.kPI = res;
			});
	}
}