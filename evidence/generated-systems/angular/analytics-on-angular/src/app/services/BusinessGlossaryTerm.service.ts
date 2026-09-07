import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BusinessGlossaryTerm} from '../models/BusinessGlossaryTerm';
import {MetricService} from '../services/Metric.service';
import {DataSetService} from '../services/DataSet.service';
import {DimensionService} from '../services/Dimension.service';
import {MeasureService} from '../services/Measure.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BusinessGlossaryTermService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	businessGlossaryTerm : BusinessGlossaryTerm;

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
	// add a BusinessGlossaryTerm
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures) : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessGlossaryTerm/create';
		const obj = {
			      		term: term,
      		definition: definition,
      		steward: steward,
      		RelatedTerms: RelatedTerms != null && RelatedTerms.length > 0 ? RelatedTerms : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
			Measures: Measures != null && Measures.length > 0 ? Measures : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BusinessGlossaryTerm
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBusinessGlossaryTerm(term, definition, steward, RelatedTerms, Metrics, Datasets, Dimensions, Measures, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BusinessGlossaryTerm/update/' + id;
		const obj = {
				      		term: term,
      		definition: definition,
      		steward: steward,
      		RelatedTerms: RelatedTerms != null && RelatedTerms.length > 0 ? RelatedTerms : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
			Measures: Measures != null && Measures.length > 0 ? Measures : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BusinessGlossaryTerm
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBusinessGlossaryTerm(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessGlossaryTerm/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BusinessGlossaryTerm
	// returns the results untouched as an Observable BusinessGlossaryTerm
	// BusinessGlossaryTerm model
	// delegates via URI
	//********************************************************************
	getBusinessGlossaryTerm(id) : Observable<BusinessGlossaryTerm> {
		const uri_ = this.apiUrl + '/BusinessGlossaryTerm/load/' + id;

		return this.http.get<BusinessGlossaryTerm>(uri_);
	}
	
	//********************************************************************
	// gets all BusinessGlossaryTerm
	// returns the results untouched as JSON representation of an
	// Observable array of BusinessGlossaryTerm models
	// delegates via URI
	//********************************************************************
	getBusinessGlossaryTerms() : Observable<BusinessGlossaryTerm[]> {
		const uri_ = this.apiUrl + '/BusinessGlossaryTerm/';

		return this
			.http.get<BusinessGlossaryTerm[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more relatedTermsIds as a RelatedTerms
	// to a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedTerms( businessGlossaryTermId, relatedTermsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );

	// split on a comma with no spaces
	var idList = relatedTermsIds.split(',')

	// iterate over array of relatedTerms ids
	idList.forEach(function (id) {
		// read the BusinessGlossaryTerm
		var businessGlossaryTerm = new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerm(id);
		// add the BusinessGlossaryTerm if not already assigned
		if ( this.businessGlossaryTerm.relatedTerms.indexOf(businessGlossaryTerm) == -1 )
		this.businessGlossaryTerm.relatedTerms.push(businessGlossaryTerm);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedTermsIds as a RelatedTerms
	// from a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedTerms( businessGlossaryTermId, relatedTermsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );


	// split on a comma with no spaces
	var idList 					= relatedTermsIds.split(',');
	var relatedTerms 	= this.businessGlossaryTerm.relatedTerms;

	if ( relatedTerms != null && relatedTermsIds != null ) {

		// iterate over array of relatedTerms ids
		relatedTerms.forEach(function (obj) {
			if ( relatedTermsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessGlossaryTerm
				this.businessGlossaryTerm.relatedTerms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more metricsIds as a Metrics
	// to a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMetrics( businessGlossaryTermId, metricsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );

	// split on a comma with no spaces
	var idList = metricsIds.split(',')

	// iterate over array of metrics ids
	idList.forEach(function (id) {
		// read the Metric
		var metric = new MetricService(this.http).getMetric(id);
		// add the Metric if not already assigned
		if ( this.businessGlossaryTerm.metrics.indexOf(metric) == -1 )
		this.businessGlossaryTerm.metrics.push(metric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more metricsIds as a Metrics
	// from a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMetrics( businessGlossaryTermId, metricsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );


	// split on a comma with no spaces
	var idList 					= metricsIds.split(',');
	var metrics 	= this.businessGlossaryTerm.metrics;

	if ( metrics != null && metricsIds != null ) {

		// iterate over array of metrics ids
		metrics.forEach(function (obj) {
			if ( metricsIds.indexOf(obj._id) > -1 ) {
				// remove the Metric
				this.businessGlossaryTerm.metrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( businessGlossaryTermId, datasetsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.businessGlossaryTerm.datasets.indexOf(dataSet) == -1 )
		this.businessGlossaryTerm.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( businessGlossaryTermId, datasetsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.businessGlossaryTerm.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.businessGlossaryTerm.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dimensionsIds as a Dimensions
	// to a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDimensions( businessGlossaryTermId, dimensionsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );

	// split on a comma with no spaces
	var idList = dimensionsIds.split(',')

	// iterate over array of dimensions ids
	idList.forEach(function (id) {
		// read the Dimension
		var dimension = new DimensionService(this.http).getDimension(id);
		// add the Dimension if not already assigned
		if ( this.businessGlossaryTerm.dimensions.indexOf(dimension) == -1 )
		this.businessGlossaryTerm.dimensions.push(dimension);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dimensionsIds as a Dimensions
	// from a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDimensions( businessGlossaryTermId, dimensionsIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );


	// split on a comma with no spaces
	var idList 					= dimensionsIds.split(',');
	var dimensions 	= this.businessGlossaryTerm.dimensions;

	if ( dimensions != null && dimensionsIds != null ) {

		// iterate over array of dimensions ids
		dimensions.forEach(function (obj) {
			if ( dimensionsIds.indexOf(obj._id) > -1 ) {
				// remove the Dimension
				this.businessGlossaryTerm.dimensions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more measuresIds as a Measures
	// to a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMeasures( businessGlossaryTermId, measuresIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );

	// split on a comma with no spaces
	var idList = measuresIds.split(',')

	// iterate over array of measures ids
	idList.forEach(function (id) {
		// read the Measure
		var measure = new MeasureService(this.http).getMeasure(id);
		// add the Measure if not already assigned
		if ( this.businessGlossaryTerm.measures.indexOf(measure) == -1 )
		this.businessGlossaryTerm.measures.push(measure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more measuresIds as a Measures
	// from a BusinessGlossaryTerm
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMeasures( businessGlossaryTermId, measuresIds ): Observable<any> {

		// get the BusinessGlossaryTerm
		this.loadHelper( businessGlossaryTermId );


	// split on a comma with no spaces
	var idList 					= measuresIds.split(',');
	var measures 	= this.businessGlossaryTerm.measures;

	if ( measures != null && measuresIds != null ) {

		// iterate over array of measures ids
		measures.forEach(function (obj) {
			if ( measuresIds.indexOf(obj._id) > -1 ) {
				// remove the Measure
				this.businessGlossaryTerm.measures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BusinessGlossaryTerm
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BusinessGlossaryTerm/update/' + this.businessGlossaryTerm;

	return  this.http.post(uri_, this.businessGlossaryTerm );
}

	//********************************************************************
	// loadHelper - internal helper to load a BusinessGlossaryTerm
	//********************************************************************	
	loadHelper( id ) {
		this.getBusinessGlossaryTerm(id)
			.subscribe((res : BusinessGlossaryTerm) => {
				this.businessGlossaryTerm = res;
			});
	}
}