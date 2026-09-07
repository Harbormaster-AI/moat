import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AnalyticsWorkspace} from '../models/AnalyticsWorkspace';
import {DataSetService} from '../services/DataSet.service';
import {DataSourceService} from '../services/DataSource.service';
import {DataPipelineService} from '../services/DataPipeline.service';
import {DashboardService} from '../services/Dashboard.service';
import {ReportService} from '../services/Report.service';
import {NotebookService} from '../services/Notebook.service';
import {Model_Service} from '../services/Model_.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import {AccessPolicyService} from '../services/AccessPolicy.service';
import {LineageNodeService} from '../services/LineageNode.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AnalyticsWorkspaceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	analyticsWorkspace : AnalyticsWorkspace;

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
	// add a AnalyticsWorkspace
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier) : Observable<any> {
		const uri_ = this.apiUrl + '/AnalyticsWorkspace/create';
		const obj = {
			      		name: name,
      		businessDomain: businessDomain,
      		ownerTeam: ownerTeam,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		DataSources: DataSources != null && DataSources.length > 0 ? DataSources : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		LineageNodes: LineageNodes != null && LineageNodes.length > 0 ? LineageNodes : null,
			GovernanceTier: GovernanceTier
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AnalyticsWorkspace
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AnalyticsWorkspace/update/' + id;
		const obj = {
				      		name: name,
      		businessDomain: businessDomain,
      		ownerTeam: ownerTeam,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		DataSources: DataSources != null && DataSources.length > 0 ? DataSources : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		LineageNodes: LineageNodes != null && LineageNodes.length > 0 ? LineageNodes : null,
			GovernanceTier: GovernanceTier
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AnalyticsWorkspace
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAnalyticsWorkspace(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AnalyticsWorkspace/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AnalyticsWorkspace
	// returns the results untouched as an Observable AnalyticsWorkspace
	// AnalyticsWorkspace model
	// delegates via URI
	//********************************************************************
	getAnalyticsWorkspace(id) : Observable<AnalyticsWorkspace> {
		const uri_ = this.apiUrl + '/AnalyticsWorkspace/load/' + id;

		return this.http.get<AnalyticsWorkspace>(uri_);
	}
	
	//********************************************************************
	// gets all AnalyticsWorkspace
	// returns the results untouched as JSON representation of an
	// Observable array of AnalyticsWorkspace models
	// delegates via URI
	//********************************************************************
	getAnalyticsWorkspaces() : Observable<AnalyticsWorkspace[]> {
		const uri_ = this.apiUrl + '/AnalyticsWorkspace/';

		return this
			.http.get<AnalyticsWorkspace[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( analyticsWorkspaceId, datasetsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.analyticsWorkspace.datasets.indexOf(dataSet) == -1 )
		this.analyticsWorkspace.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( analyticsWorkspaceId, datasetsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.analyticsWorkspace.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.analyticsWorkspace.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataSourcesIds as a DataSources
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataSources( analyticsWorkspaceId, dataSourcesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = dataSourcesIds.split(',')

	// iterate over array of dataSources ids
	idList.forEach(function (id) {
		// read the DataSource
		var dataSource = new DataSourceService(this.http).getDataSource(id);
		// add the DataSource if not already assigned
		if ( this.analyticsWorkspace.dataSources.indexOf(dataSource) == -1 )
		this.analyticsWorkspace.dataSources.push(dataSource);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataSourcesIds as a DataSources
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataSources( analyticsWorkspaceId, dataSourcesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= dataSourcesIds.split(',');
	var dataSources 	= this.analyticsWorkspace.dataSources;

	if ( dataSources != null && dataSourcesIds != null ) {

		// iterate over array of dataSources ids
		dataSources.forEach(function (obj) {
			if ( dataSourcesIds.indexOf(obj._id) > -1 ) {
				// remove the DataSource
				this.analyticsWorkspace.dataSources.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more pipelinesIds as a Pipelines
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPipelines( analyticsWorkspaceId, pipelinesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = pipelinesIds.split(',')

	// iterate over array of pipelines ids
	idList.forEach(function (id) {
		// read the DataPipeline
		var dataPipeline = new DataPipelineService(this.http).getDataPipeline(id);
		// add the DataPipeline if not already assigned
		if ( this.analyticsWorkspace.pipelines.indexOf(dataPipeline) == -1 )
		this.analyticsWorkspace.pipelines.push(dataPipeline);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pipelinesIds as a Pipelines
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePipelines( analyticsWorkspaceId, pipelinesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= pipelinesIds.split(',');
	var pipelines 	= this.analyticsWorkspace.pipelines;

	if ( pipelines != null && pipelinesIds != null ) {

		// iterate over array of pipelines ids
		pipelines.forEach(function (obj) {
			if ( pipelinesIds.indexOf(obj._id) > -1 ) {
				// remove the DataPipeline
				this.analyticsWorkspace.pipelines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dashboardsIds as a Dashboards
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDashboards( analyticsWorkspaceId, dashboardsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = dashboardsIds.split(',')

	// iterate over array of dashboards ids
	idList.forEach(function (id) {
		// read the Dashboard
		var dashboard = new DashboardService(this.http).getDashboard(id);
		// add the Dashboard if not already assigned
		if ( this.analyticsWorkspace.dashboards.indexOf(dashboard) == -1 )
		this.analyticsWorkspace.dashboards.push(dashboard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dashboardsIds as a Dashboards
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDashboards( analyticsWorkspaceId, dashboardsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= dashboardsIds.split(',');
	var dashboards 	= this.analyticsWorkspace.dashboards;

	if ( dashboards != null && dashboardsIds != null ) {

		// iterate over array of dashboards ids
		dashboards.forEach(function (obj) {
			if ( dashboardsIds.indexOf(obj._id) > -1 ) {
				// remove the Dashboard
				this.analyticsWorkspace.dashboards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( analyticsWorkspaceId, reportsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.analyticsWorkspace.reports.indexOf(report) == -1 )
		this.analyticsWorkspace.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( analyticsWorkspaceId, reportsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.analyticsWorkspace.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.analyticsWorkspace.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notebooksIds as a Notebooks
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotebooks( analyticsWorkspaceId, notebooksIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = notebooksIds.split(',')

	// iterate over array of notebooks ids
	idList.forEach(function (id) {
		// read the Notebook
		var notebook = new NotebookService(this.http).getNotebook(id);
		// add the Notebook if not already assigned
		if ( this.analyticsWorkspace.notebooks.indexOf(notebook) == -1 )
		this.analyticsWorkspace.notebooks.push(notebook);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notebooksIds as a Notebooks
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotebooks( analyticsWorkspaceId, notebooksIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= notebooksIds.split(',');
	var notebooks 	= this.analyticsWorkspace.notebooks;

	if ( notebooks != null && notebooksIds != null ) {

		// iterate over array of notebooks ids
		notebooks.forEach(function (obj) {
			if ( notebooksIds.indexOf(obj._id) > -1 ) {
				// remove the Notebook
				this.analyticsWorkspace.notebooks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( analyticsWorkspaceId, modelsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.analyticsWorkspace.models.indexOf(model_) == -1 )
		this.analyticsWorkspace.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( analyticsWorkspaceId, modelsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.analyticsWorkspace.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.analyticsWorkspace.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featureSetsIds as a FeatureSets
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatureSets( analyticsWorkspaceId, featureSetsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = featureSetsIds.split(',')

	// iterate over array of featureSets ids
	idList.forEach(function (id) {
		// read the FeatureSet
		var featureSet = new FeatureSetService(this.http).getFeatureSet(id);
		// add the FeatureSet if not already assigned
		if ( this.analyticsWorkspace.featureSets.indexOf(featureSet) == -1 )
		this.analyticsWorkspace.featureSets.push(featureSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featureSetsIds as a FeatureSets
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatureSets( analyticsWorkspaceId, featureSetsIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= featureSetsIds.split(',');
	var featureSets 	= this.analyticsWorkspace.featureSets;

	if ( featureSets != null && featureSetsIds != null ) {

		// iterate over array of featureSets ids
		featureSets.forEach(function (obj) {
			if ( featureSetsIds.indexOf(obj._id) > -1 ) {
				// remove the FeatureSet
				this.analyticsWorkspace.featureSets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( analyticsWorkspaceId, policiesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the AccessPolicy
		var accessPolicy = new AccessPolicyService(this.http).getAccessPolicy(id);
		// add the AccessPolicy if not already assigned
		if ( this.analyticsWorkspace.policies.indexOf(accessPolicy) == -1 )
		this.analyticsWorkspace.policies.push(accessPolicy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( analyticsWorkspaceId, policiesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.analyticsWorkspace.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the AccessPolicy
				this.analyticsWorkspace.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more lineageNodesIds as a LineageNodes
	// to a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLineageNodes( analyticsWorkspaceId, lineageNodesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );

	// split on a comma with no spaces
	var idList = lineageNodesIds.split(',')

	// iterate over array of lineageNodes ids
	idList.forEach(function (id) {
		// read the LineageNode
		var lineageNode = new LineageNodeService(this.http).getLineageNode(id);
		// add the LineageNode if not already assigned
		if ( this.analyticsWorkspace.lineageNodes.indexOf(lineageNode) == -1 )
		this.analyticsWorkspace.lineageNodes.push(lineageNode);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lineageNodesIds as a LineageNodes
	// from a AnalyticsWorkspace
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLineageNodes( analyticsWorkspaceId, lineageNodesIds ): Observable<any> {

		// get the AnalyticsWorkspace
		this.loadHelper( analyticsWorkspaceId );


	// split on a comma with no spaces
	var idList 					= lineageNodesIds.split(',');
	var lineageNodes 	= this.analyticsWorkspace.lineageNodes;

	if ( lineageNodes != null && lineageNodesIds != null ) {

		// iterate over array of lineageNodes ids
		lineageNodes.forEach(function (obj) {
			if ( lineageNodesIds.indexOf(obj._id) > -1 ) {
				// remove the LineageNode
				this.analyticsWorkspace.lineageNodes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AnalyticsWorkspace
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AnalyticsWorkspace/update/' + this.analyticsWorkspace;

	return  this.http.post(uri_, this.analyticsWorkspace );
}

	//********************************************************************
	// loadHelper - internal helper to load a AnalyticsWorkspace
	//********************************************************************	
	loadHelper( id ) {
		this.getAnalyticsWorkspace(id)
			.subscribe((res : AnalyticsWorkspace) => {
				this.analyticsWorkspace = res;
			});
	}
}