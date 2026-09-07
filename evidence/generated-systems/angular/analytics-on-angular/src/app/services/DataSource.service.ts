import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataSource} from '../models/DataSource';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSetService} from '../services/DataSet.service';
import {DataPipelineService} from '../services/DataPipeline.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataSourceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataSource : DataSource;

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
	// add a DataSource
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format) : Observable<any> {
		const uri_ = this.apiUrl + '/DataSource/create';
		const obj = {
			      		name: name,
      		connection: connection,
      		Streaming: Streaming,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		ProducedDatasets: ProducedDatasets != null && ProducedDatasets.length > 0 ? ProducedDatasets : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		SourceType: SourceType,
			Format: Format
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataSource
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataSource/update/' + id;
		const obj = {
				      		name: name,
      		connection: connection,
      		Streaming: Streaming,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		ProducedDatasets: ProducedDatasets != null && ProducedDatasets.length > 0 ? ProducedDatasets : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		SourceType: SourceType,
			Format: Format
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataSource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataSource(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataSource/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataSource
	// returns the results untouched as an Observable DataSource
	// DataSource model
	// delegates via URI
	//********************************************************************
	getDataSource(id) : Observable<DataSource> {
		const uri_ = this.apiUrl + '/DataSource/load/' + id;

		return this.http.get<DataSource>(uri_);
	}
	
	//********************************************************************
	// gets all DataSource
	// returns the results untouched as JSON representation of an
	// Observable array of DataSource models
	// delegates via URI
	//********************************************************************
	getDataSources() : Observable<DataSource[]> {
		const uri_ = this.apiUrl + '/DataSource/';

		return this
			.http.get<DataSource[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a DataSource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( dataSourceId, _workspaceId ): Observable<any> {

		// get the DataSource from storage
		this.loadHelper( dataSourceId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.dataSource.workspace = tmp;

	// save the DataSource
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a DataSource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( dataSourceId ): Observable<any> {

		// get the DataSource from storage
		this.loadHelper( dataSourceId );

	// assign Workspace to null
	this.dataSource.workspace = null;

	// save the DataSource
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more producedDatasetsIds as a ProducedDatasets
	// to a DataSource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProducedDatasets( dataSourceId, producedDatasetsIds ): Observable<any> {

		// get the DataSource
		this.loadHelper( dataSourceId );

	// split on a comma with no spaces
	var idList = producedDatasetsIds.split(',')

	// iterate over array of producedDatasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dataSource.producedDatasets.indexOf(dataSet) == -1 )
		this.dataSource.producedDatasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more producedDatasetsIds as a ProducedDatasets
	// from a DataSource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProducedDatasets( dataSourceId, producedDatasetsIds ): Observable<any> {

		// get the DataSource
		this.loadHelper( dataSourceId );


	// split on a comma with no spaces
	var idList 					= producedDatasetsIds.split(',');
	var producedDatasets 	= this.dataSource.producedDatasets;

	if ( producedDatasets != null && producedDatasetsIds != null ) {

		// iterate over array of producedDatasets ids
		producedDatasets.forEach(function (obj) {
			if ( producedDatasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dataSource.producedDatasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more pipelinesIds as a Pipelines
	// to a DataSource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPipelines( dataSourceId, pipelinesIds ): Observable<any> {

		// get the DataSource
		this.loadHelper( dataSourceId );

	// split on a comma with no spaces
	var idList = pipelinesIds.split(',')

	// iterate over array of pipelines ids
	idList.forEach(function (id) {
		// read the DataPipeline
		var dataPipeline = new DataPipelineService(this.http).getDataPipeline(id);
		// add the DataPipeline if not already assigned
		if ( this.dataSource.pipelines.indexOf(dataPipeline) == -1 )
		this.dataSource.pipelines.push(dataPipeline);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pipelinesIds as a Pipelines
	// from a DataSource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePipelines( dataSourceId, pipelinesIds ): Observable<any> {

		// get the DataSource
		this.loadHelper( dataSourceId );


	// split on a comma with no spaces
	var idList 					= pipelinesIds.split(',');
	var pipelines 	= this.dataSource.pipelines;

	if ( pipelines != null && pipelinesIds != null ) {

		// iterate over array of pipelines ids
		pipelines.forEach(function (obj) {
			if ( pipelinesIds.indexOf(obj._id) > -1 ) {
				// remove the DataPipeline
				this.dataSource.pipelines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataSource
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataSource/update/' + this.dataSource;

	return  this.http.post(uri_, this.dataSource );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataSource
	//********************************************************************	
	loadHelper( id ) {
		this.getDataSource(id)
			.subscribe((res : DataSource) => {
				this.dataSource = res;
			});
	}
}