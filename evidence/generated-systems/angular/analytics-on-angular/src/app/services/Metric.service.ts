import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Metric} from '../models/Metric';
import {SemanticModelService} from '../services/SemanticModel.service';
import {DataSetService} from '../services/DataSet.service';
import {BusinessGlossaryTermService} from '../services/BusinessGlossaryTerm.service';
import {AlertService} from '../services/Alert.service';
import {VisualizationService} from '../services/Visualization.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MetricService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	metric : Metric;

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
	// add a Metric
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType) : Observable<any> {
		const uri_ = this.apiUrl + '/Metric/create';
		const obj = {
			      		name: name,
      		expression: expression,
      		unit: unit,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
			MetricType: MetricType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Metric
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMetric(name, expression, unit, SemanticModel, Datasets, GlossaryTerms, Alerts, Visualizations, MetricType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Metric/update/' + id;
		const obj = {
				      		name: name,
      		expression: expression,
      		unit: unit,
      		SemanticModel: SemanticModel != null && SemanticModel.length > 0 ? SemanticModel : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		GlossaryTerms: GlossaryTerms != null && GlossaryTerms.length > 0 ? GlossaryTerms : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
			MetricType: MetricType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Metric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMetric(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Metric/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Metric
	// returns the results untouched as an Observable Metric
	// Metric model
	// delegates via URI
	//********************************************************************
	getMetric(id) : Observable<Metric> {
		const uri_ = this.apiUrl + '/Metric/load/' + id;

		return this.http.get<Metric>(uri_);
	}
	
	//********************************************************************
	// gets all Metric
	// returns the results untouched as JSON representation of an
	// Observable array of Metric models
	// delegates via URI
	//********************************************************************
	getMetrics() : Observable<Metric[]> {
		const uri_ = this.apiUrl + '/Metric/';

		return this
			.http.get<Metric[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SemanticModel on a Metric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSemanticModel( metricId, _semanticModelId ): Observable<any> {

		// get the Metric from storage
		this.loadHelper( metricId );

	// get the SemanticModel from storage
	var tmp 	= new SemanticModelService(this.http).getSemanticModel(_semanticModelId);

	// assign the SemanticModel
	this.metric.semanticModel = tmp;

	// save the Metric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SemanticModel on a Metric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSemanticModel( metricId ): Observable<any> {

		// get the Metric from storage
		this.loadHelper( metricId );

	// assign SemanticModel to null
	this.metric.semanticModel = null;

	// save the Metric
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( metricId, datasetsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.metric.datasets.indexOf(dataSet) == -1 )
		this.metric.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( metricId, datasetsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.metric.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.metric.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more glossaryTermsIds as a GlossaryTerms
	// to a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGlossaryTerms( metricId, glossaryTermsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );

	// split on a comma with no spaces
	var idList = glossaryTermsIds.split(',')

	// iterate over array of glossaryTerms ids
	idList.forEach(function (id) {
		// read the BusinessGlossaryTerm
		var businessGlossaryTerm = new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerm(id);
		// add the BusinessGlossaryTerm if not already assigned
		if ( this.metric.glossaryTerms.indexOf(businessGlossaryTerm) == -1 )
		this.metric.glossaryTerms.push(businessGlossaryTerm);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more glossaryTermsIds as a GlossaryTerms
	// from a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGlossaryTerms( metricId, glossaryTermsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );


	// split on a comma with no spaces
	var idList 					= glossaryTermsIds.split(',');
	var glossaryTerms 	= this.metric.glossaryTerms;

	if ( glossaryTerms != null && glossaryTermsIds != null ) {

		// iterate over array of glossaryTerms ids
		glossaryTerms.forEach(function (obj) {
			if ( glossaryTermsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessGlossaryTerm
				this.metric.glossaryTerms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( metricId, alertsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the Alert
		var alert = new AlertService(this.http).getAlert(id);
		// add the Alert if not already assigned
		if ( this.metric.alerts.indexOf(alert) == -1 )
		this.metric.alerts.push(alert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( metricId, alertsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.metric.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the Alert
				this.metric.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more visualizationsIds as a Visualizations
	// to a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVisualizations( metricId, visualizationsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );

	// split on a comma with no spaces
	var idList = visualizationsIds.split(',')

	// iterate over array of visualizations ids
	idList.forEach(function (id) {
		// read the Visualization
		var visualization = new VisualizationService(this.http).getVisualization(id);
		// add the Visualization if not already assigned
		if ( this.metric.visualizations.indexOf(visualization) == -1 )
		this.metric.visualizations.push(visualization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more visualizationsIds as a Visualizations
	// from a Metric
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVisualizations( metricId, visualizationsIds ): Observable<any> {

		// get the Metric
		this.loadHelper( metricId );


	// split on a comma with no spaces
	var idList 					= visualizationsIds.split(',');
	var visualizations 	= this.metric.visualizations;

	if ( visualizations != null && visualizationsIds != null ) {

		// iterate over array of visualizations ids
		visualizations.forEach(function (obj) {
			if ( visualizationsIds.indexOf(obj._id) > -1 ) {
				// remove the Visualization
				this.metric.visualizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Metric
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Metric/update/' + this.metric;

	return  this.http.post(uri_, this.metric );
}

	//********************************************************************
	// loadHelper - internal helper to load a Metric
	//********************************************************************	
	loadHelper( id ) {
		this.getMetric(id)
			.subscribe((res : Metric) => {
				this.metric = res;
			});
	}
}