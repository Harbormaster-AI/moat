import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Tag} from '../models/Tag';
import {DataSetService} from '../services/DataSet.service';
import {Model_Service} from '../services/Model_.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {DashboardService} from '../services/Dashboard.service';
import {ReportService} from '../services/Report.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import {MetricService} from '../services/Metric.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TagService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	tag : Tag;

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
	// add a Tag
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category) : Observable<any> {
		const uri_ = this.apiUrl + '/Tag/create';
		const obj = {
			      		name: name,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		ModelVersions: ModelVersions != null && ModelVersions.length > 0 ? ModelVersions : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
			Category: Category
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Tag
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Tag/update/' + id;
		const obj = {
				      		name: name,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		ModelVersions: ModelVersions != null && ModelVersions.length > 0 ? ModelVersions : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
			Category: Category
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Tag
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTag(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Tag/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Tag
	// returns the results untouched as an Observable Tag
	// Tag model
	// delegates via URI
	//********************************************************************
	getTag(id) : Observable<Tag> {
		const uri_ = this.apiUrl + '/Tag/load/' + id;

		return this.http.get<Tag>(uri_);
	}
	
	//********************************************************************
	// gets all Tag
	// returns the results untouched as JSON representation of an
	// Observable array of Tag models
	// delegates via URI
	//********************************************************************
	getTags() : Observable<Tag[]> {
		const uri_ = this.apiUrl + '/Tag/';

		return this
			.http.get<Tag[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( tagId, datasetsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.tag.datasets.indexOf(dataSet) == -1 )
		this.tag.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( tagId, datasetsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.tag.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.tag.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( tagId, modelsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.tag.models.indexOf(model_) == -1 )
		this.tag.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( tagId, modelsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.tag.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.tag.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelVersionsIds as a ModelVersions
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModelVersions( tagId, modelVersionsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = modelVersionsIds.split(',')

	// iterate over array of modelVersions ids
	idList.forEach(function (id) {
		// read the ModelVersion
		var modelVersion = new ModelVersionService(this.http).getModelVersion(id);
		// add the ModelVersion if not already assigned
		if ( this.tag.modelVersions.indexOf(modelVersion) == -1 )
		this.tag.modelVersions.push(modelVersion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelVersionsIds as a ModelVersions
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModelVersions( tagId, modelVersionsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= modelVersionsIds.split(',');
	var modelVersions 	= this.tag.modelVersions;

	if ( modelVersions != null && modelVersionsIds != null ) {

		// iterate over array of modelVersions ids
		modelVersions.forEach(function (obj) {
			if ( modelVersionsIds.indexOf(obj._id) > -1 ) {
				// remove the ModelVersion
				this.tag.modelVersions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dashboardsIds as a Dashboards
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDashboards( tagId, dashboardsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = dashboardsIds.split(',')

	// iterate over array of dashboards ids
	idList.forEach(function (id) {
		// read the Dashboard
		var dashboard = new DashboardService(this.http).getDashboard(id);
		// add the Dashboard if not already assigned
		if ( this.tag.dashboards.indexOf(dashboard) == -1 )
		this.tag.dashboards.push(dashboard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dashboardsIds as a Dashboards
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDashboards( tagId, dashboardsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= dashboardsIds.split(',');
	var dashboards 	= this.tag.dashboards;

	if ( dashboards != null && dashboardsIds != null ) {

		// iterate over array of dashboards ids
		dashboards.forEach(function (obj) {
			if ( dashboardsIds.indexOf(obj._id) > -1 ) {
				// remove the Dashboard
				this.tag.dashboards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( tagId, reportsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.tag.reports.indexOf(report) == -1 )
		this.tag.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( tagId, reportsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.tag.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.tag.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featureSetsIds as a FeatureSets
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatureSets( tagId, featureSetsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = featureSetsIds.split(',')

	// iterate over array of featureSets ids
	idList.forEach(function (id) {
		// read the FeatureSet
		var featureSet = new FeatureSetService(this.http).getFeatureSet(id);
		// add the FeatureSet if not already assigned
		if ( this.tag.featureSets.indexOf(featureSet) == -1 )
		this.tag.featureSets.push(featureSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featureSetsIds as a FeatureSets
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatureSets( tagId, featureSetsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= featureSetsIds.split(',');
	var featureSets 	= this.tag.featureSets;

	if ( featureSets != null && featureSetsIds != null ) {

		// iterate over array of featureSets ids
		featureSets.forEach(function (obj) {
			if ( featureSetsIds.indexOf(obj._id) > -1 ) {
				// remove the FeatureSet
				this.tag.featureSets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more metricsIds as a Metrics
	// to a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMetrics( tagId, metricsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );

	// split on a comma with no spaces
	var idList = metricsIds.split(',')

	// iterate over array of metrics ids
	idList.forEach(function (id) {
		// read the Metric
		var metric = new MetricService(this.http).getMetric(id);
		// add the Metric if not already assigned
		if ( this.tag.metrics.indexOf(metric) == -1 )
		this.tag.metrics.push(metric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more metricsIds as a Metrics
	// from a Tag
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMetrics( tagId, metricsIds ): Observable<any> {

		// get the Tag
		this.loadHelper( tagId );


	// split on a comma with no spaces
	var idList 					= metricsIds.split(',');
	var metrics 	= this.tag.metrics;

	if ( metrics != null && metricsIds != null ) {

		// iterate over array of metrics ids
		metrics.forEach(function (obj) {
			if ( metricsIds.indexOf(obj._id) > -1 ) {
				// remove the Metric
				this.tag.metrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Tag
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Tag/update/' + this.tag;

	return  this.http.post(uri_, this.tag );
}

	//********************************************************************
	// loadHelper - internal helper to load a Tag
	//********************************************************************	
	loadHelper( id ) {
		this.getTag(id)
			.subscribe((res : Tag) => {
				this.tag = res;
			});
	}
}