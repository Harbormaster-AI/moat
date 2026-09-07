import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataPipeline} from '../models/DataPipeline';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataTaskService} from '../services/DataTask.service';
import {DataSourceService} from '../services/DataSource.service';
import {DataSetService} from '../services/DataSet.service';
import {LineageNodeService} from '../services/LineageNode.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataPipelineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataPipeline : DataPipeline;

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
	// add a DataPipeline
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/DataPipeline/create';
		const obj = {
			      		name: name,
      		schedule: schedule,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Tasks: Tasks != null && Tasks.length > 0 ? Tasks : null,
      		Sources: Sources != null && Sources.length > 0 ? Sources : null,
      		Outputs: Outputs != null && Outputs.length > 0 ? Outputs : null,
      		LineageNode: LineageNode != null && LineageNode.length > 0 ? LineageNode : null,
      		TriggerType: TriggerType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataPipeline/update/' + id;
		const obj = {
				      		name: name,
      		schedule: schedule,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Tasks: Tasks != null && Tasks.length > 0 ? Tasks : null,
      		Sources: Sources != null && Sources.length > 0 ? Sources : null,
      		Outputs: Outputs != null && Outputs.length > 0 ? Outputs : null,
      		LineageNode: LineageNode != null && LineageNode.length > 0 ? LineageNode : null,
      		TriggerType: TriggerType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataPipeline(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataPipeline/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataPipeline
	// returns the results untouched as an Observable DataPipeline
	// DataPipeline model
	// delegates via URI
	//********************************************************************
	getDataPipeline(id) : Observable<DataPipeline> {
		const uri_ = this.apiUrl + '/DataPipeline/load/' + id;

		return this.http.get<DataPipeline>(uri_);
	}
	
	//********************************************************************
	// gets all DataPipeline
	// returns the results untouched as JSON representation of an
	// Observable array of DataPipeline models
	// delegates via URI
	//********************************************************************
	getDataPipelines() : Observable<DataPipeline[]> {
		const uri_ = this.apiUrl + '/DataPipeline/';

		return this
			.http.get<DataPipeline[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( dataPipelineId, _workspaceId ): Observable<any> {

		// get the DataPipeline from storage
		this.loadHelper( dataPipelineId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.dataPipeline.workspace = tmp;

	// save the DataPipeline
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( dataPipelineId ): Observable<any> {

		// get the DataPipeline from storage
		this.loadHelper( dataPipelineId );

	// assign Workspace to null
	this.dataPipeline.workspace = null;

	// save the DataPipeline
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LineageNode on a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineageNode( dataPipelineId, _lineageNodeId ): Observable<any> {

		// get the DataPipeline from storage
		this.loadHelper( dataPipelineId );

	// get the LineageNode from storage
	var tmp 	= new LineageNodeService(this.http).getLineageNode(_lineageNodeId);

	// assign the LineageNode
	this.dataPipeline.lineageNode = tmp;

	// save the DataPipeline
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineageNode on a DataPipeline
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineageNode( dataPipelineId ): Observable<any> {

		// get the DataPipeline from storage
		this.loadHelper( dataPipelineId );

	// assign LineageNode to null
	this.dataPipeline.lineageNode = null;

	// save the DataPipeline
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more tasksIds as a Tasks
	// to a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTasks( dataPipelineId, tasksIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );

	// split on a comma with no spaces
	var idList = tasksIds.split(',')

	// iterate over array of tasks ids
	idList.forEach(function (id) {
		// read the DataTask
		var dataTask = new DataTaskService(this.http).getDataTask(id);
		// add the DataTask if not already assigned
		if ( this.dataPipeline.tasks.indexOf(dataTask) == -1 )
		this.dataPipeline.tasks.push(dataTask);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tasksIds as a Tasks
	// from a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTasks( dataPipelineId, tasksIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );


	// split on a comma with no spaces
	var idList 					= tasksIds.split(',');
	var tasks 	= this.dataPipeline.tasks;

	if ( tasks != null && tasksIds != null ) {

		// iterate over array of tasks ids
		tasks.forEach(function (obj) {
			if ( tasksIds.indexOf(obj._id) > -1 ) {
				// remove the DataTask
				this.dataPipeline.tasks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more sourcesIds as a Sources
	// to a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSources( dataPipelineId, sourcesIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );

	// split on a comma with no spaces
	var idList = sourcesIds.split(',')

	// iterate over array of sources ids
	idList.forEach(function (id) {
		// read the DataSource
		var dataSource = new DataSourceService(this.http).getDataSource(id);
		// add the DataSource if not already assigned
		if ( this.dataPipeline.sources.indexOf(dataSource) == -1 )
		this.dataPipeline.sources.push(dataSource);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more sourcesIds as a Sources
	// from a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSources( dataPipelineId, sourcesIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );


	// split on a comma with no spaces
	var idList 					= sourcesIds.split(',');
	var sources 	= this.dataPipeline.sources;

	if ( sources != null && sourcesIds != null ) {

		// iterate over array of sources ids
		sources.forEach(function (obj) {
			if ( sourcesIds.indexOf(obj._id) > -1 ) {
				// remove the DataSource
				this.dataPipeline.sources.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more outputsIds as a Outputs
	// to a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOutputs( dataPipelineId, outputsIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );

	// split on a comma with no spaces
	var idList = outputsIds.split(',')

	// iterate over array of outputs ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dataPipeline.outputs.indexOf(dataSet) == -1 )
		this.dataPipeline.outputs.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more outputsIds as a Outputs
	// from a DataPipeline
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOutputs( dataPipelineId, outputsIds ): Observable<any> {

		// get the DataPipeline
		this.loadHelper( dataPipelineId );


	// split on a comma with no spaces
	var idList 					= outputsIds.split(',');
	var outputs 	= this.dataPipeline.outputs;

	if ( outputs != null && outputsIds != null ) {

		// iterate over array of outputs ids
		outputs.forEach(function (obj) {
			if ( outputsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dataPipeline.outputs.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataPipeline
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataPipeline/update/' + this.dataPipeline;

	return  this.http.post(uri_, this.dataPipeline );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataPipeline
	//********************************************************************	
	loadHelper( id ) {
		this.getDataPipeline(id)
			.subscribe((res : DataPipeline) => {
				this.dataPipeline = res;
			});
	}
}