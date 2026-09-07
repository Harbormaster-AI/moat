import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Measure} from '../models/Measure';
import {SemanticModelService} from '../services/SemanticModel.service';
import {DataSetService} from '../services/DataSet.service';
import {BusinessGlossaryTermService} from '../services/BusinessGlossaryTerm.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MeasureService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	measure : Measure;

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
	// add a Measure
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation) : Observable<any> {
		const uri_ = this.apiUrl + '/Measure/create';
		const obj = {
			      		name: name,
      		format: format,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
			Aggregation: Aggregation
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Measure
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMeasure(name, format, SemanticModel, Datasets, GlossaryTerms, Aggregation, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Measure/update/' + id;
		const obj = {
				      		name: name,
      		format: format,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
			Aggregation: Aggregation
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Measure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMeasure(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Measure/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Measure
	// returns the results untouched as an Observable Measure
	// Measure model
	// delegates via URI
	//********************************************************************
	getMeasure(id) : Observable<Measure> {
		const uri_ = this.apiUrl + '/Measure/load/' + id;

		return this.http.get<Measure>(uri_);
	}
	
	//********************************************************************
	// gets all Measure
	// returns the results untouched as JSON representation of an
	// Observable array of Measure models
	// delegates via URI
	//********************************************************************
	getMeasures() : Observable<Measure[]> {
		const uri_ = this.apiUrl + '/Measure/';

		return this
			.http.get<Measure[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SemanticModel on a Measure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSemanticModel( measureId, _semanticModelId ): Observable<any> {

		// get the Measure from storage
		this.loadHelper( measureId );

	// get the SemanticModel from storage
	var tmp 	= new SemanticModelService(this.http).getSemanticModel(_semanticModelId);

	// assign the SemanticModel
	this.measure.semanticModel = tmp;

	// save the Measure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SemanticModel on a Measure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSemanticModel( measureId ): Observable<any> {

		// get the Measure from storage
		this.loadHelper( measureId );

	// assign SemanticModel to null
	this.measure.semanticModel = null;

	// save the Measure
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Measure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( measureId, datasetsIds ): Observable<any> {

		// get the Measure
		this.loadHelper( measureId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.measure.datasets.indexOf(dataSet) == -1 )
		this.measure.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Measure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( measureId, datasetsIds ): Observable<any> {

		// get the Measure
		this.loadHelper( measureId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.measure.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.measure.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more glossaryTermsIds as a GlossaryTerms
	// to a Measure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGlossaryTerms( measureId, glossaryTermsIds ): Observable<any> {

		// get the Measure
		this.loadHelper( measureId );

	// split on a comma with no spaces
	var idList = glossaryTermsIds.split(',')

	// iterate over array of glossaryTerms ids
	idList.forEach(function (id) {
		// read the BusinessGlossaryTerm
		var businessGlossaryTerm = new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerm(id);
		// add the BusinessGlossaryTerm if not already assigned
		if ( this.measure.glossaryTerms.indexOf(businessGlossaryTerm) == -1 )
		this.measure.glossaryTerms.push(businessGlossaryTerm);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more glossaryTermsIds as a GlossaryTerms
	// from a Measure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGlossaryTerms( measureId, glossaryTermsIds ): Observable<any> {

		// get the Measure
		this.loadHelper( measureId );


	// split on a comma with no spaces
	var idList 					= glossaryTermsIds.split(',');
	var glossaryTerms 	= this.measure.glossaryTerms;

	if ( glossaryTerms != null && glossaryTermsIds != null ) {

		// iterate over array of glossaryTerms ids
		glossaryTerms.forEach(function (obj) {
			if ( glossaryTermsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessGlossaryTerm
				this.measure.glossaryTerms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Measure
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Measure/update/' + this.measure;

	return  this.http.post(uri_, this.measure );
}

	//********************************************************************
	// loadHelper - internal helper to load a Measure
	//********************************************************************	
	loadHelper( id ) {
		this.getMeasure(id)
			.subscribe((res : Measure) => {
				this.measure = res;
			});
	}
}