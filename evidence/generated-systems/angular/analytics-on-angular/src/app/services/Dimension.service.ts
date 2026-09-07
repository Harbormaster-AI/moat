import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Dimension} from '../models/Dimension';
import {SemanticModelService} from '../services/SemanticModel.service';
import {DataSetService} from '../services/DataSet.service';
import {BusinessGlossaryTermService} from '../services/BusinessGlossaryTerm.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DimensionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dimension : Dimension;

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
	// add a Dimension
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType) : Observable<any> {
		const uri_ = this.apiUrl + '/Dimension/create';
		const obj = {
			      		name: name,
      		typeTime: typeTime,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
			DimensionType: DimensionType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Dimension
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Dimension/update/' + id;
		const obj = {
				      		name: name,
      		typeTime: typeTime,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
			DimensionType: DimensionType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Dimension
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDimension(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Dimension/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Dimension
	// returns the results untouched as an Observable Dimension
	// Dimension model
	// delegates via URI
	//********************************************************************
	getDimension(id) : Observable<Dimension> {
		const uri_ = this.apiUrl + '/Dimension/load/' + id;

		return this.http.get<Dimension>(uri_);
	}
	
	//********************************************************************
	// gets all Dimension
	// returns the results untouched as JSON representation of an
	// Observable array of Dimension models
	// delegates via URI
	//********************************************************************
	getDimensions() : Observable<Dimension[]> {
		const uri_ = this.apiUrl + '/Dimension/';

		return this
			.http.get<Dimension[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SemanticModel on a Dimension
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSemanticModel( dimensionId, _semanticModelId ): Observable<any> {

		// get the Dimension from storage
		this.loadHelper( dimensionId );

	// get the SemanticModel from storage
	var tmp 	= new SemanticModelService(this.http).getSemanticModel(_semanticModelId);

	// assign the SemanticModel
	this.dimension.semanticModel = tmp;

	// save the Dimension
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SemanticModel on a Dimension
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSemanticModel( dimensionId ): Observable<any> {

		// get the Dimension from storage
		this.loadHelper( dimensionId );

	// assign SemanticModel to null
	this.dimension.semanticModel = null;

	// save the Dimension
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Dimension
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( dimensionId, datasetsIds ): Observable<any> {

		// get the Dimension
		this.loadHelper( dimensionId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dimension.datasets.indexOf(dataSet) == -1 )
		this.dimension.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Dimension
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( dimensionId, datasetsIds ): Observable<any> {

		// get the Dimension
		this.loadHelper( dimensionId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.dimension.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dimension.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more glossaryTermsIds as a GlossaryTerms
	// to a Dimension
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGlossaryTerms( dimensionId, glossaryTermsIds ): Observable<any> {

		// get the Dimension
		this.loadHelper( dimensionId );

	// split on a comma with no spaces
	var idList = glossaryTermsIds.split(',')

	// iterate over array of glossaryTerms ids
	idList.forEach(function (id) {
		// read the BusinessGlossaryTerm
		var businessGlossaryTerm = new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerm(id);
		// add the BusinessGlossaryTerm if not already assigned
		if ( this.dimension.glossaryTerms.indexOf(businessGlossaryTerm) == -1 )
		this.dimension.glossaryTerms.push(businessGlossaryTerm);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more glossaryTermsIds as a GlossaryTerms
	// from a Dimension
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGlossaryTerms( dimensionId, glossaryTermsIds ): Observable<any> {

		// get the Dimension
		this.loadHelper( dimensionId );


	// split on a comma with no spaces
	var idList 					= glossaryTermsIds.split(',');
	var glossaryTerms 	= this.dimension.glossaryTerms;

	if ( glossaryTerms != null && glossaryTermsIds != null ) {

		// iterate over array of glossaryTerms ids
		glossaryTerms.forEach(function (obj) {
			if ( glossaryTermsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessGlossaryTerm
				this.dimension.glossaryTerms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Dimension
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Dimension/update/' + this.dimension;

	return  this.http.post(uri_, this.dimension );
}

	//********************************************************************
	// loadHelper - internal helper to load a Dimension
	//********************************************************************	
	loadHelper( id ) {
		this.getDimension(id)
			.subscribe((res : Dimension) => {
				this.dimension = res;
			});
	}
}