import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataTask} from '../models/DataTask';
import {DataPipelineService} from '../services/DataPipeline.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataTaskService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataTask : DataTask;

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
	// add a DataTask
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType) : Observable<any> {
		const uri_ = this.apiUrl + '/DataTask/create';
		const obj = {
			      		name: name,
      		command: command,
      		retries: retries,
      		Pipeline: Pipeline != null && Pipeline.length > 0 ? Pipeline : null,
      		InputDatasets: InputDatasets != null && InputDatasets.length > 0 ? InputDatasets : null,
      		OutputDatasets: OutputDatasets != null && OutputDatasets.length > 0 ? OutputDatasets : null,
			TaskType: TaskType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataTask/update/' + id;
		const obj = {
				      		name: name,
      		command: command,
      		retries: retries,
      		Pipeline: Pipeline != null && Pipeline.length > 0 ? Pipeline : null,
      		InputDatasets: InputDatasets != null && InputDatasets.length > 0 ? InputDatasets : null,
      		OutputDatasets: OutputDatasets != null && OutputDatasets.length > 0 ? OutputDatasets : null,
			TaskType: TaskType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataTask(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataTask/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataTask
	// returns the results untouched as an Observable DataTask
	// DataTask model
	// delegates via URI
	//********************************************************************
	getDataTask(id) : Observable<DataTask> {
		const uri_ = this.apiUrl + '/DataTask/load/' + id;

		return this.http.get<DataTask>(uri_);
	}
	
	//********************************************************************
	// gets all DataTask
	// returns the results untouched as JSON representation of an
	// Observable array of DataTask models
	// delegates via URI
	//********************************************************************
	getDataTasks() : Observable<DataTask[]> {
		const uri_ = this.apiUrl + '/DataTask/';

		return this
			.http.get<DataTask[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Pipeline on a DataTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPipeline( dataTaskId, _pipelineId ): Observable<any> {

		// get the DataTask from storage
		this.loadHelper( dataTaskId );

	// get the DataPipeline from storage
	var tmp 	= new DataPipelineService(this.http).getDataPipeline(_pipelineId);

	// assign the Pipeline
	this.dataTask.pipeline = tmp;

	// save the DataTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Pipeline on a DataTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPipeline( dataTaskId ): Observable<any> {

		// get the DataTask from storage
		this.loadHelper( dataTaskId );

	// assign Pipeline to null
	this.dataTask.pipeline = null;

	// save the DataTask
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inputDatasetsIds as a InputDatasets
	// to a DataTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInputDatasets( dataTaskId, inputDatasetsIds ): Observable<any> {

		// get the DataTask
		this.loadHelper( dataTaskId );

	// split on a comma with no spaces
	var idList = inputDatasetsIds.split(',')

	// iterate over array of inputDatasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dataTask.inputDatasets.indexOf(dataSet) == -1 )
		this.dataTask.inputDatasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inputDatasetsIds as a InputDatasets
	// from a DataTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInputDatasets( dataTaskId, inputDatasetsIds ): Observable<any> {

		// get the DataTask
		this.loadHelper( dataTaskId );


	// split on a comma with no spaces
	var idList 					= inputDatasetsIds.split(',');
	var inputDatasets 	= this.dataTask.inputDatasets;

	if ( inputDatasets != null && inputDatasetsIds != null ) {

		// iterate over array of inputDatasets ids
		inputDatasets.forEach(function (obj) {
			if ( inputDatasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dataTask.inputDatasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more outputDatasetsIds as a OutputDatasets
	// to a DataTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOutputDatasets( dataTaskId, outputDatasetsIds ): Observable<any> {

		// get the DataTask
		this.loadHelper( dataTaskId );

	// split on a comma with no spaces
	var idList = outputDatasetsIds.split(',')

	// iterate over array of outputDatasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dataTask.outputDatasets.indexOf(dataSet) == -1 )
		this.dataTask.outputDatasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more outputDatasetsIds as a OutputDatasets
	// from a DataTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOutputDatasets( dataTaskId, outputDatasetsIds ): Observable<any> {

		// get the DataTask
		this.loadHelper( dataTaskId );


	// split on a comma with no spaces
	var idList 					= outputDatasetsIds.split(',');
	var outputDatasets 	= this.dataTask.outputDatasets;

	if ( outputDatasets != null && outputDatasetsIds != null ) {

		// iterate over array of outputDatasets ids
		outputDatasets.forEach(function (obj) {
			if ( outputDatasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dataTask.outputDatasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataTask
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataTask/update/' + this.dataTask;

	return  this.http.post(uri_, this.dataTask );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataTask
	//********************************************************************	
	loadHelper( id ) {
		this.getDataTask(id)
			.subscribe((res : DataTask) => {
				this.dataTask = res;
			});
	}
}