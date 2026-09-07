import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Visualization} from '../models/Visualization';
import {DashboardService} from '../services/Dashboard.service';
import {ReportService} from '../services/Report.service';
import {MetricService} from '../services/Metric.service';
import {DimensionService} from '../services/Dimension.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class VisualizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	visualization : Visualization;

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
	// add a Visualization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType) : Observable<any> {
		const uri_ = this.apiUrl + '/Visualization/create';
		const obj = {
			      		title: title,
      		options: options,
      		Dashboard: Dashboard != null && Dashboard.length > 0 ? Dashboard : null,
      		Report: Report != null && Report.length > 0 ? Report : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
			ChartType: ChartType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Visualization/update/' + id;
		const obj = {
				      		title: title,
      		options: options,
      		Dashboard: Dashboard != null && Dashboard.length > 0 ? Dashboard : null,
      		Report: Report != null && Report.length > 0 ? Report : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
			ChartType: ChartType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteVisualization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Visualization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Visualization
	// returns the results untouched as an Observable Visualization
	// Visualization model
	// delegates via URI
	//********************************************************************
	getVisualization(id) : Observable<Visualization> {
		const uri_ = this.apiUrl + '/Visualization/load/' + id;

		return this.http.get<Visualization>(uri_);
	}
	
	//********************************************************************
	// gets all Visualization
	// returns the results untouched as JSON representation of an
	// Observable array of Visualization models
	// delegates via URI
	//********************************************************************
	getVisualizations() : Observable<Visualization[]> {
		const uri_ = this.apiUrl + '/Visualization/';

		return this
			.http.get<Visualization[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Dashboard on a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDashboard( visualizationId, _dashboardId ): Observable<any> {

		// get the Visualization from storage
		this.loadHelper( visualizationId );

	// get the Dashboard from storage
	var tmp 	= new DashboardService(this.http).getDashboard(_dashboardId);

	// assign the Dashboard
	this.visualization.dashboard = tmp;

	// save the Visualization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dashboard on a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDashboard( visualizationId ): Observable<any> {

		// get the Visualization from storage
		this.loadHelper( visualizationId );

	// assign Dashboard to null
	this.visualization.dashboard = null;

	// save the Visualization
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Report on a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignReport( visualizationId, _reportId ): Observable<any> {

		// get the Visualization from storage
		this.loadHelper( visualizationId );

	// get the Report from storage
	var tmp 	= new ReportService(this.http).getReport(_reportId);

	// assign the Report
	this.visualization.report = tmp;

	// save the Visualization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Report on a Visualization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignReport( visualizationId ): Observable<any> {

		// get the Visualization from storage
		this.loadHelper( visualizationId );

	// assign Report to null
	this.visualization.report = null;

	// save the Visualization
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more metricsIds as a Metrics
	// to a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMetrics( visualizationId, metricsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );

	// split on a comma with no spaces
	var idList = metricsIds.split(',')

	// iterate over array of metrics ids
	idList.forEach(function (id) {
		// read the Metric
		var metric = new MetricService(this.http).getMetric(id);
		// add the Metric if not already assigned
		if ( this.visualization.metrics.indexOf(metric) == -1 )
		this.visualization.metrics.push(metric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more metricsIds as a Metrics
	// from a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMetrics( visualizationId, metricsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );


	// split on a comma with no spaces
	var idList 					= metricsIds.split(',');
	var metrics 	= this.visualization.metrics;

	if ( metrics != null && metricsIds != null ) {

		// iterate over array of metrics ids
		metrics.forEach(function (obj) {
			if ( metricsIds.indexOf(obj._id) > -1 ) {
				// remove the Metric
				this.visualization.metrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dimensionsIds as a Dimensions
	// to a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDimensions( visualizationId, dimensionsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );

	// split on a comma with no spaces
	var idList = dimensionsIds.split(',')

	// iterate over array of dimensions ids
	idList.forEach(function (id) {
		// read the Dimension
		var dimension = new DimensionService(this.http).getDimension(id);
		// add the Dimension if not already assigned
		if ( this.visualization.dimensions.indexOf(dimension) == -1 )
		this.visualization.dimensions.push(dimension);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dimensionsIds as a Dimensions
	// from a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDimensions( visualizationId, dimensionsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );


	// split on a comma with no spaces
	var idList 					= dimensionsIds.split(',');
	var dimensions 	= this.visualization.dimensions;

	if ( dimensions != null && dimensionsIds != null ) {

		// iterate over array of dimensions ids
		dimensions.forEach(function (obj) {
			if ( dimensionsIds.indexOf(obj._id) > -1 ) {
				// remove the Dimension
				this.visualization.dimensions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( visualizationId, datasetsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.visualization.datasets.indexOf(dataSet) == -1 )
		this.visualization.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Visualization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( visualizationId, datasetsIds ): Observable<any> {

		// get the Visualization
		this.loadHelper( visualizationId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.visualization.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.visualization.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Visualization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Visualization/update/' + this.visualization;

	return  this.http.post(uri_, this.visualization );
}

	//********************************************************************
	// loadHelper - internal helper to load a Visualization
	//********************************************************************	
	loadHelper( id ) {
		this.getVisualization(id)
			.subscribe((res : Visualization) => {
				this.visualization = res;
			});
	}
}