import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Experiment} from '../models/Experiment';
import {CampaignService} from '../services/Campaign.service';
import {ExperimentVariantService} from '../services/ExperimentVariant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExperimentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	experiment : Experiment;

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
	// add a Experiment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Experiment/create';
		const obj = {
			      		name: name,
      		hypothesis: hypothesis,
      		startDate: startDate,
      		endDate: endDate,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExperiment(name, hypothesis, startDate, endDate, Campaign, Variants, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Experiment/update/' + id;
		const obj = {
				      		name: name,
      		hypothesis: hypothesis,
      		startDate: startDate,
      		endDate: endDate,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExperiment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Experiment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Experiment
	// returns the results untouched as an Observable Experiment
	// Experiment model
	// delegates via URI
	//********************************************************************
	getExperiment(id) : Observable<Experiment> {
		const uri_ = this.apiUrl + '/Experiment/load/' + id;

		return this.http.get<Experiment>(uri_);
	}
	
	//********************************************************************
	// gets all Experiment
	// returns the results untouched as JSON representation of an
	// Observable array of Experiment models
	// delegates via URI
	//********************************************************************
	getExperiments() : Observable<Experiment[]> {
		const uri_ = this.apiUrl + '/Experiment/';

		return this
			.http.get<Experiment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( experimentId, _campaignId ): Observable<any> {

		// get the Experiment from storage
		this.loadHelper( experimentId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.experiment.campaign = tmp;

	// save the Experiment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( experimentId ): Observable<any> {

		// get the Experiment from storage
		this.loadHelper( experimentId );

	// assign Campaign to null
	this.experiment.campaign = null;

	// save the Experiment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( experimentId, variantsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the ExperimentVariant
		var experimentVariant = new ExperimentVariantService(this.http).getExperimentVariant(id);
		// add the ExperimentVariant if not already assigned
		if ( this.experiment.variants.indexOf(experimentVariant) == -1 )
		this.experiment.variants.push(experimentVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( experimentId, variantsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.experiment.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the ExperimentVariant
				this.experiment.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Experiment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Experiment/update/' + this.experiment;

	return  this.http.post(uri_, this.experiment );
}

	//********************************************************************
	// loadHelper - internal helper to load a Experiment
	//********************************************************************	
	loadHelper( id ) {
		this.getExperiment(id)
			.subscribe((res : Experiment) => {
				this.experiment = res;
			});
	}
}