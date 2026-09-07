import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ExperimentVariant} from '../models/ExperimentVariant';
import {ExperimentService} from '../services/Experiment.service';
import {CreativeVariationService} from '../services/CreativeVariation.service';
import {LineItemService} from '../services/LineItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExperimentVariantService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	experimentVariant : ExperimentVariant;

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
	// add a ExperimentVariant
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem) : Observable<any> {
		const uri_ = this.apiUrl + '/ExperimentVariant/create';
		const obj = {
			      		name: name,
      		allocation: allocation,
      		Experiment: Experiment != null && Experiment.length > 0 ? Experiment : null,
      		CreativeVariation: CreativeVariation != null && CreativeVariation.length > 0 ? CreativeVariation : null,
			LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ExperimentVariant/update/' + id;
		const obj = {
				      		name: name,
      		allocation: allocation,
      		Experiment: Experiment != null && Experiment.length > 0 ? Experiment : null,
      		CreativeVariation: CreativeVariation != null && CreativeVariation.length > 0 ? CreativeVariation : null,
			LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExperimentVariant(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ExperimentVariant/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ExperimentVariant
	// returns the results untouched as an Observable ExperimentVariant
	// ExperimentVariant model
	// delegates via URI
	//********************************************************************
	getExperimentVariant(id) : Observable<ExperimentVariant> {
		const uri_ = this.apiUrl + '/ExperimentVariant/load/' + id;

		return this.http.get<ExperimentVariant>(uri_);
	}
	
	//********************************************************************
	// gets all ExperimentVariant
	// returns the results untouched as JSON representation of an
	// Observable array of ExperimentVariant models
	// delegates via URI
	//********************************************************************
	getExperimentVariants() : Observable<ExperimentVariant[]> {
		const uri_ = this.apiUrl + '/ExperimentVariant/';

		return this
			.http.get<ExperimentVariant[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Experiment on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignExperiment( experimentVariantId, _experimentId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// get the Experiment from storage
	var tmp 	= new ExperimentService(this.http).getExperiment(_experimentId);

	// assign the Experiment
	this.experimentVariant.experiment = tmp;

	// save the ExperimentVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Experiment on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignExperiment( experimentVariantId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// assign Experiment to null
	this.experimentVariant.experiment = null;

	// save the ExperimentVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CreativeVariation on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreativeVariation( experimentVariantId, _creativeVariationId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// get the CreativeVariation from storage
	var tmp 	= new CreativeVariationService(this.http).getCreativeVariation(_creativeVariationId);

	// assign the CreativeVariation
	this.experimentVariant.creativeVariation = tmp;

	// save the ExperimentVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CreativeVariation on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreativeVariation( experimentVariantId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// assign CreativeVariation to null
	this.experimentVariant.creativeVariation = null;

	// save the ExperimentVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LineItem on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineItem( experimentVariantId, _lineItemId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// get the LineItem from storage
	var tmp 	= new LineItemService(this.http).getLineItem(_lineItemId);

	// assign the LineItem
	this.experimentVariant.lineItem = tmp;

	// save the ExperimentVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineItem on a ExperimentVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineItem( experimentVariantId ): Observable<any> {

		// get the ExperimentVariant from storage
		this.loadHelper( experimentVariantId );

	// assign LineItem to null
	this.experimentVariant.lineItem = null;

	// save the ExperimentVariant
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ExperimentVariant
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ExperimentVariant/update/' + this.experimentVariant;

	return  this.http.post(uri_, this.experimentVariant );
}

	//********************************************************************
	// loadHelper - internal helper to load a ExperimentVariant
	//********************************************************************	
	loadHelper( id ) {
		this.getExperimentVariant(id)
			.subscribe((res : ExperimentVariant) => {
				this.experimentVariant = res;
			});
	}
}