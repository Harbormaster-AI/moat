import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LineageNode} from '../models/LineageNode';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSetService} from '../services/DataSet.service';
import {Model_Service} from '../services/Model_.service';
import {DataPipelineService} from '../services/DataPipeline.service';
import {DashboardService} from '../services/Dashboard.service';
import {ReportService} from '../services/Report.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LineageNodeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	lineageNode : LineageNode;

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
	// add a LineageNode
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType) : Observable<any> {
		const uri_ = this.apiUrl + '/LineageNode/create';
		const obj = {
			      		name: name,
      		qualifiedName: qualifiedName,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Inputs: Inputs != null && Inputs.length > 0 ? Inputs : null,
      		Outputs: Outputs != null && Outputs.length > 0 ? Outputs : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
			NodeType: NodeType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LineageNode
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LineageNode/update/' + id;
		const obj = {
				      		name: name,
      		qualifiedName: qualifiedName,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Inputs: Inputs != null && Inputs.length > 0 ? Inputs : null,
      		Outputs: Outputs != null && Outputs.length > 0 ? Outputs : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
			NodeType: NodeType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LineageNode
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLineageNode(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LineageNode/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LineageNode
	// returns the results untouched as an Observable LineageNode
	// LineageNode model
	// delegates via URI
	//********************************************************************
	getLineageNode(id) : Observable<LineageNode> {
		const uri_ = this.apiUrl + '/LineageNode/load/' + id;

		return this.http.get<LineageNode>(uri_);
	}
	
	//********************************************************************
	// gets all LineageNode
	// returns the results untouched as JSON representation of an
	// Observable array of LineageNode models
	// delegates via URI
	//********************************************************************
	getLineageNodes() : Observable<LineageNode[]> {
		const uri_ = this.apiUrl + '/LineageNode/';

		return this
			.http.get<LineageNode[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a LineageNode
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( lineageNodeId, _workspaceId ): Observable<any> {

		// get the LineageNode from storage
		this.loadHelper( lineageNodeId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.lineageNode.workspace = tmp;

	// save the LineageNode
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a LineageNode
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( lineageNodeId ): Observable<any> {

		// get the LineageNode from storage
		this.loadHelper( lineageNodeId );

	// assign Workspace to null
	this.lineageNode.workspace = null;

	// save the LineageNode
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inputsIds as a Inputs
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInputs( lineageNodeId, inputsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = inputsIds.split(',')

	// iterate over array of inputs ids
	idList.forEach(function (id) {
		// read the LineageNode
		var lineageNode = new LineageNodeService(this.http).getLineageNode(id);
		// add the LineageNode if not already assigned
		if ( this.lineageNode.inputs.indexOf(lineageNode) == -1 )
		this.lineageNode.inputs.push(lineageNode);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inputsIds as a Inputs
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInputs( lineageNodeId, inputsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= inputsIds.split(',');
	var inputs 	= this.lineageNode.inputs;

	if ( inputs != null && inputsIds != null ) {

		// iterate over array of inputs ids
		inputs.forEach(function (obj) {
			if ( inputsIds.indexOf(obj._id) > -1 ) {
				// remove the LineageNode
				this.lineageNode.inputs.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more outputsIds as a Outputs
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOutputs( lineageNodeId, outputsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = outputsIds.split(',')

	// iterate over array of outputs ids
	idList.forEach(function (id) {
		// read the LineageNode
		var lineageNode = new LineageNodeService(this.http).getLineageNode(id);
		// add the LineageNode if not already assigned
		if ( this.lineageNode.outputs.indexOf(lineageNode) == -1 )
		this.lineageNode.outputs.push(lineageNode);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more outputsIds as a Outputs
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOutputs( lineageNodeId, outputsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= outputsIds.split(',');
	var outputs 	= this.lineageNode.outputs;

	if ( outputs != null && outputsIds != null ) {

		// iterate over array of outputs ids
		outputs.forEach(function (obj) {
			if ( outputsIds.indexOf(obj._id) > -1 ) {
				// remove the LineageNode
				this.lineageNode.outputs.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( lineageNodeId, datasetsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.lineageNode.datasets.indexOf(dataSet) == -1 )
		this.lineageNode.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( lineageNodeId, datasetsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.lineageNode.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.lineageNode.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( lineageNodeId, modelsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.lineageNode.models.indexOf(model_) == -1 )
		this.lineageNode.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( lineageNodeId, modelsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.lineageNode.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.lineageNode.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more pipelinesIds as a Pipelines
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPipelines( lineageNodeId, pipelinesIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = pipelinesIds.split(',')

	// iterate over array of pipelines ids
	idList.forEach(function (id) {
		// read the DataPipeline
		var dataPipeline = new DataPipelineService(this.http).getDataPipeline(id);
		// add the DataPipeline if not already assigned
		if ( this.lineageNode.pipelines.indexOf(dataPipeline) == -1 )
		this.lineageNode.pipelines.push(dataPipeline);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pipelinesIds as a Pipelines
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePipelines( lineageNodeId, pipelinesIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= pipelinesIds.split(',');
	var pipelines 	= this.lineageNode.pipelines;

	if ( pipelines != null && pipelinesIds != null ) {

		// iterate over array of pipelines ids
		pipelines.forEach(function (obj) {
			if ( pipelinesIds.indexOf(obj._id) > -1 ) {
				// remove the DataPipeline
				this.lineageNode.pipelines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dashboardsIds as a Dashboards
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDashboards( lineageNodeId, dashboardsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = dashboardsIds.split(',')

	// iterate over array of dashboards ids
	idList.forEach(function (id) {
		// read the Dashboard
		var dashboard = new DashboardService(this.http).getDashboard(id);
		// add the Dashboard if not already assigned
		if ( this.lineageNode.dashboards.indexOf(dashboard) == -1 )
		this.lineageNode.dashboards.push(dashboard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dashboardsIds as a Dashboards
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDashboards( lineageNodeId, dashboardsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= dashboardsIds.split(',');
	var dashboards 	= this.lineageNode.dashboards;

	if ( dashboards != null && dashboardsIds != null ) {

		// iterate over array of dashboards ids
		dashboards.forEach(function (obj) {
			if ( dashboardsIds.indexOf(obj._id) > -1 ) {
				// remove the Dashboard
				this.lineageNode.dashboards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( lineageNodeId, reportsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.lineageNode.reports.indexOf(report) == -1 )
		this.lineageNode.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a LineageNode
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( lineageNodeId, reportsIds ): Observable<any> {

		// get the LineageNode
		this.loadHelper( lineageNodeId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.lineageNode.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.lineageNode.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LineageNode
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LineageNode/update/' + this.lineageNode;

	return  this.http.post(uri_, this.lineageNode );
}

	//********************************************************************
	// loadHelper - internal helper to load a LineageNode
	//********************************************************************	
	loadHelper( id ) {
		this.getLineageNode(id)
			.subscribe((res : LineageNode) => {
				this.lineageNode = res;
			});
	}
}