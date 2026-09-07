import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SemanticModel} from '../models/SemanticModel';
import {DataSetService} from '../services/DataSet.service';
import {MetricService} from '../services/Metric.service';
import {DimensionService} from '../services/Dimension.service';
import {MeasureService} from '../services/Measure.service';
import {BusinessGlossaryTermService} from '../services/BusinessGlossaryTerm.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SemanticModelService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	semanticModel : SemanticModel;

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
	// add a SemanticModel
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms) : Observable<any> {
		const uri_ = this.apiUrl + '/SemanticModel/create';
		const obj = {
			      		name: name,
      		version: version,
      		grain: grain,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Measures: Measures != null && Measures.length > 0 ? Measures : null,
			GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SemanticModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SemanticModel/update/' + id;
		const obj = {
				      		name: name,
      		version: version,
      		grain: grain,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Measures: Measures != null && Measures.length > 0 ? Measures : null,
			GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SemanticModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSemanticModel(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SemanticModel/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SemanticModel
	// returns the results untouched as an Observable SemanticModel
	// SemanticModel model
	// delegates via URI
	//********************************************************************
	getSemanticModel(id) : Observable<SemanticModel> {
		const uri_ = this.apiUrl + '/SemanticModel/load/' + id;

		return this.http.get<SemanticModel>(uri_);
	}
	
	//********************************************************************
	// gets all SemanticModel
	// returns the results untouched as JSON representation of an
	// Observable array of SemanticModel models
	// delegates via URI
	//********************************************************************
	getSemanticModels() : Observable<SemanticModel[]> {
		const uri_ = this.apiUrl + '/SemanticModel/';

		return this
			.http.get<SemanticModel[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( semanticModelId, datasetsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.semanticModel.datasets.indexOf(dataSet) == -1 )
		this.semanticModel.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( semanticModelId, datasetsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.semanticModel.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.semanticModel.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more metricsIds as a Metrics
	// to a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMetrics( semanticModelId, metricsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );

	// split on a comma with no spaces
	var idList = metricsIds.split(',')

	// iterate over array of metrics ids
	idList.forEach(function (id) {
		// read the Metric
		var metric = new MetricService(this.http).getMetric(id);
		// add the Metric if not already assigned
		if ( this.semanticModel.metrics.indexOf(metric) == -1 )
		this.semanticModel.metrics.push(metric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more metricsIds as a Metrics
	// from a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMetrics( semanticModelId, metricsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );


	// split on a comma with no spaces
	var idList 					= metricsIds.split(',');
	var metrics 	= this.semanticModel.metrics;

	if ( metrics != null && metricsIds != null ) {

		// iterate over array of metrics ids
		metrics.forEach(function (obj) {
			if ( metricsIds.indexOf(obj._id) > -1 ) {
				// remove the Metric
				this.semanticModel.metrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dimensionsIds as a Dimensions
	// to a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDimensions( semanticModelId, dimensionsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );

	// split on a comma with no spaces
	var idList = dimensionsIds.split(',')

	// iterate over array of dimensions ids
	idList.forEach(function (id) {
		// read the Dimension
		var dimension = new DimensionService(this.http).getDimension(id);
		// add the Dimension if not already assigned
		if ( this.semanticModel.dimensions.indexOf(dimension) == -1 )
		this.semanticModel.dimensions.push(dimension);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dimensionsIds as a Dimensions
	// from a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDimensions( semanticModelId, dimensionsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );


	// split on a comma with no spaces
	var idList 					= dimensionsIds.split(',');
	var dimensions 	= this.semanticModel.dimensions;

	if ( dimensions != null && dimensionsIds != null ) {

		// iterate over array of dimensions ids
		dimensions.forEach(function (obj) {
			if ( dimensionsIds.indexOf(obj._id) > -1 ) {
				// remove the Dimension
				this.semanticModel.dimensions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more measuresIds as a Measures
	// to a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMeasures( semanticModelId, measuresIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );

	// split on a comma with no spaces
	var idList = measuresIds.split(',')

	// iterate over array of measures ids
	idList.forEach(function (id) {
		// read the Measure
		var measure = new MeasureService(this.http).getMeasure(id);
		// add the Measure if not already assigned
		if ( this.semanticModel.measures.indexOf(measure) == -1 )
		this.semanticModel.measures.push(measure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more measuresIds as a Measures
	// from a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMeasures( semanticModelId, measuresIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );


	// split on a comma with no spaces
	var idList 					= measuresIds.split(',');
	var measures 	= this.semanticModel.measures;

	if ( measures != null && measuresIds != null ) {

		// iterate over array of measures ids
		measures.forEach(function (obj) {
			if ( measuresIds.indexOf(obj._id) > -1 ) {
				// remove the Measure
				this.semanticModel.measures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more glossaryTermsIds as a GlossaryTerms
	// to a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGlossaryTerms( semanticModelId, glossaryTermsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );

	// split on a comma with no spaces
	var idList = glossaryTermsIds.split(',')

	// iterate over array of glossaryTerms ids
	idList.forEach(function (id) {
		// read the BusinessGlossaryTerm
		var businessGlossaryTerm = new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerm(id);
		// add the BusinessGlossaryTerm if not already assigned
		if ( this.semanticModel.glossaryTerms.indexOf(businessGlossaryTerm) == -1 )
		this.semanticModel.glossaryTerms.push(businessGlossaryTerm);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more glossaryTermsIds as a GlossaryTerms
	// from a SemanticModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGlossaryTerms( semanticModelId, glossaryTermsIds ): Observable<any> {

		// get the SemanticModel
		this.loadHelper( semanticModelId );


	// split on a comma with no spaces
	var idList 					= glossaryTermsIds.split(',');
	var glossaryTerms 	= this.semanticModel.glossaryTerms;

	if ( glossaryTerms != null && glossaryTermsIds != null ) {

		// iterate over array of glossaryTerms ids
		glossaryTerms.forEach(function (obj) {
			if ( glossaryTermsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessGlossaryTerm
				this.semanticModel.glossaryTerms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a SemanticModel
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SemanticModel/update/' + this.semanticModel;

	return  this.http.post(uri_, this.semanticModel );
}

	//********************************************************************
	// loadHelper - internal helper to load a SemanticModel
	//********************************************************************	
	loadHelper( id ) {
		this.getSemanticModel(id)
			.subscribe((res : SemanticModel) => {
				this.semanticModel = res;
			});
	}
}