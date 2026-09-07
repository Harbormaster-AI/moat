import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataSet} from '../models/DataSet';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSourceService} from '../services/DataSource.service';
import {DataPipelineService} from '../services/DataPipeline.service';
import {SemanticModelService} from '../services/SemanticModel.service';
import {DimensionService} from '../services/Dimension.service';
import {MeasureService} from '../services/Measure.service';
import {MetricService} from '../services/Metric.service';
import {QualityRuleService} from '../services/QualityRule.service';
import {LineageNodeService} from '../services/LineageNode.service';
import {TagService} from '../services/Tag.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataSetService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataSet : DataSet;

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
	// add a DataSet
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat) : Observable<any> {
		const uri_ = this.apiUrl + '/DataSet/create';
		const obj = {
			      		name: name,
      		schemaVersion: schemaVersion,
      		refreshSchedule: refreshSchedule,
      		Sensitive: Sensitive,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Sources: Sources != null && Sources.length > 0 ? Sources : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		SemanticModels: SemanticModels != null && SemanticModels.length > 0 ? SemanticModels : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Measures: Measures != null && Measures.length > 0 ? Measures : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		QualityRules: QualityRules != null && QualityRules.length > 0 ? QualityRules : null,
      		LineageNode: LineageNode != null && LineageNode.length > 0 ? LineageNode : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			DataFormat: DataFormat
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataSet/update/' + id;
		const obj = {
				      		name: name,
      		schemaVersion: schemaVersion,
      		refreshSchedule: refreshSchedule,
      		Sensitive: Sensitive,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Sources: Sources != null && Sources.length > 0 ? Sources : null,
      		Pipelines: Pipelines != null && Pipelines.length > 0 ? Pipelines : null,
      		SemanticModels: SemanticModels != null && SemanticModels.length > 0 ? SemanticModels : null,
      		Dimensions: Dimensions != null && Dimensions.length > 0 ? Dimensions : null,
      		Measures: Measures != null && Measures.length > 0 ? Measures : null,
      		Metrics: Metrics != null && Metrics.length > 0 ? Metrics : null,
      		QualityRules: QualityRules != null && QualityRules.length > 0 ? QualityRules : null,
      		LineageNode: LineageNode != null && LineageNode.length > 0 ? LineageNode : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			DataFormat: DataFormat
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataSet(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataSet/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataSet
	// returns the results untouched as an Observable DataSet
	// DataSet model
	// delegates via URI
	//********************************************************************
	getDataSet(id) : Observable<DataSet> {
		const uri_ = this.apiUrl + '/DataSet/load/' + id;

		return this.http.get<DataSet>(uri_);
	}
	
	//********************************************************************
	// gets all DataSet
	// returns the results untouched as JSON representation of an
	// Observable array of DataSet models
	// delegates via URI
	//********************************************************************
	getDataSets() : Observable<DataSet[]> {
		const uri_ = this.apiUrl + '/DataSet/';

		return this
			.http.get<DataSet[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( dataSetId, _workspaceId ): Observable<any> {

		// get the DataSet from storage
		this.loadHelper( dataSetId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.dataSet.workspace = tmp;

	// save the DataSet
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( dataSetId ): Observable<any> {

		// get the DataSet from storage
		this.loadHelper( dataSetId );

	// assign Workspace to null
	this.dataSet.workspace = null;

	// save the DataSet
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LineageNode on a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineageNode( dataSetId, _lineageNodeId ): Observable<any> {

		// get the DataSet from storage
		this.loadHelper( dataSetId );

	// get the LineageNode from storage
	var tmp 	= new LineageNodeService(this.http).getLineageNode(_lineageNodeId);

	// assign the LineageNode
	this.dataSet.lineageNode = tmp;

	// save the DataSet
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineageNode on a DataSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineageNode( dataSetId ): Observable<any> {

		// get the DataSet from storage
		this.loadHelper( dataSetId );

	// assign LineageNode to null
	this.dataSet.lineageNode = null;

	// save the DataSet
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more sourcesIds as a Sources
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSources( dataSetId, sourcesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = sourcesIds.split(',')

	// iterate over array of sources ids
	idList.forEach(function (id) {
		// read the DataSource
		var dataSource = new DataSourceService(this.http).getDataSource(id);
		// add the DataSource if not already assigned
		if ( this.dataSet.sources.indexOf(dataSource) == -1 )
		this.dataSet.sources.push(dataSource);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more sourcesIds as a Sources
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSources( dataSetId, sourcesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= sourcesIds.split(',');
	var sources 	= this.dataSet.sources;

	if ( sources != null && sourcesIds != null ) {

		// iterate over array of sources ids
		sources.forEach(function (obj) {
			if ( sourcesIds.indexOf(obj._id) > -1 ) {
				// remove the DataSource
				this.dataSet.sources.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more pipelinesIds as a Pipelines
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPipelines( dataSetId, pipelinesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = pipelinesIds.split(',')

	// iterate over array of pipelines ids
	idList.forEach(function (id) {
		// read the DataPipeline
		var dataPipeline = new DataPipelineService(this.http).getDataPipeline(id);
		// add the DataPipeline if not already assigned
		if ( this.dataSet.pipelines.indexOf(dataPipeline) == -1 )
		this.dataSet.pipelines.push(dataPipeline);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pipelinesIds as a Pipelines
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePipelines( dataSetId, pipelinesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= pipelinesIds.split(',');
	var pipelines 	= this.dataSet.pipelines;

	if ( pipelines != null && pipelinesIds != null ) {

		// iterate over array of pipelines ids
		pipelines.forEach(function (obj) {
			if ( pipelinesIds.indexOf(obj._id) > -1 ) {
				// remove the DataPipeline
				this.dataSet.pipelines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more semanticModelsIds as a SemanticModels
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSemanticModels( dataSetId, semanticModelsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = semanticModelsIds.split(',')

	// iterate over array of semanticModels ids
	idList.forEach(function (id) {
		// read the SemanticModel
		var semanticModel = new SemanticModelService(this.http).getSemanticModel(id);
		// add the SemanticModel if not already assigned
		if ( this.dataSet.semanticModels.indexOf(semanticModel) == -1 )
		this.dataSet.semanticModels.push(semanticModel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more semanticModelsIds as a SemanticModels
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSemanticModels( dataSetId, semanticModelsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= semanticModelsIds.split(',');
	var semanticModels 	= this.dataSet.semanticModels;

	if ( semanticModels != null && semanticModelsIds != null ) {

		// iterate over array of semanticModels ids
		semanticModels.forEach(function (obj) {
			if ( semanticModelsIds.indexOf(obj._id) > -1 ) {
				// remove the SemanticModel
				this.dataSet.semanticModels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dimensionsIds as a Dimensions
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDimensions( dataSetId, dimensionsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = dimensionsIds.split(',')

	// iterate over array of dimensions ids
	idList.forEach(function (id) {
		// read the Dimension
		var dimension = new DimensionService(this.http).getDimension(id);
		// add the Dimension if not already assigned
		if ( this.dataSet.dimensions.indexOf(dimension) == -1 )
		this.dataSet.dimensions.push(dimension);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dimensionsIds as a Dimensions
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDimensions( dataSetId, dimensionsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= dimensionsIds.split(',');
	var dimensions 	= this.dataSet.dimensions;

	if ( dimensions != null && dimensionsIds != null ) {

		// iterate over array of dimensions ids
		dimensions.forEach(function (obj) {
			if ( dimensionsIds.indexOf(obj._id) > -1 ) {
				// remove the Dimension
				this.dataSet.dimensions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more measuresIds as a Measures
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMeasures( dataSetId, measuresIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = measuresIds.split(',')

	// iterate over array of measures ids
	idList.forEach(function (id) {
		// read the Measure
		var measure = new MeasureService(this.http).getMeasure(id);
		// add the Measure if not already assigned
		if ( this.dataSet.measures.indexOf(measure) == -1 )
		this.dataSet.measures.push(measure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more measuresIds as a Measures
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMeasures( dataSetId, measuresIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= measuresIds.split(',');
	var measures 	= this.dataSet.measures;

	if ( measures != null && measuresIds != null ) {

		// iterate over array of measures ids
		measures.forEach(function (obj) {
			if ( measuresIds.indexOf(obj._id) > -1 ) {
				// remove the Measure
				this.dataSet.measures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more metricsIds as a Metrics
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMetrics( dataSetId, metricsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = metricsIds.split(',')

	// iterate over array of metrics ids
	idList.forEach(function (id) {
		// read the Metric
		var metric = new MetricService(this.http).getMetric(id);
		// add the Metric if not already assigned
		if ( this.dataSet.metrics.indexOf(metric) == -1 )
		this.dataSet.metrics.push(metric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more metricsIds as a Metrics
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMetrics( dataSetId, metricsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= metricsIds.split(',');
	var metrics 	= this.dataSet.metrics;

	if ( metrics != null && metricsIds != null ) {

		// iterate over array of metrics ids
		metrics.forEach(function (obj) {
			if ( metricsIds.indexOf(obj._id) > -1 ) {
				// remove the Metric
				this.dataSet.metrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more qualityRulesIds as a QualityRules
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQualityRules( dataSetId, qualityRulesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = qualityRulesIds.split(',')

	// iterate over array of qualityRules ids
	idList.forEach(function (id) {
		// read the QualityRule
		var qualityRule = new QualityRuleService(this.http).getQualityRule(id);
		// add the QualityRule if not already assigned
		if ( this.dataSet.qualityRules.indexOf(qualityRule) == -1 )
		this.dataSet.qualityRules.push(qualityRule);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more qualityRulesIds as a QualityRules
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQualityRules( dataSetId, qualityRulesIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= qualityRulesIds.split(',');
	var qualityRules 	= this.dataSet.qualityRules;

	if ( qualityRules != null && qualityRulesIds != null ) {

		// iterate over array of qualityRules ids
		qualityRules.forEach(function (obj) {
			if ( qualityRulesIds.indexOf(obj._id) > -1 ) {
				// remove the QualityRule
				this.dataSet.qualityRules.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tagsIds as a Tags
	// to a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTags( dataSetId, tagsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );

	// split on a comma with no spaces
	var idList = tagsIds.split(',')

	// iterate over array of tags ids
	idList.forEach(function (id) {
		// read the Tag
		var tag = new TagService(this.http).getTag(id);
		// add the Tag if not already assigned
		if ( this.dataSet.tags.indexOf(tag) == -1 )
		this.dataSet.tags.push(tag);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tagsIds as a Tags
	// from a DataSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTags( dataSetId, tagsIds ): Observable<any> {

		// get the DataSet
		this.loadHelper( dataSetId );


	// split on a comma with no spaces
	var idList 					= tagsIds.split(',');
	var tags 	= this.dataSet.tags;

	if ( tags != null && tagsIds != null ) {

		// iterate over array of tags ids
		tags.forEach(function (obj) {
			if ( tagsIds.indexOf(obj._id) > -1 ) {
				// remove the Tag
				this.dataSet.tags.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataSet
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataSet/update/' + this.dataSet;

	return  this.http.post(uri_, this.dataSet );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataSet
	//********************************************************************	
	loadHelper( id ) {
		this.getDataSet(id)
			.subscribe((res : DataSet) => {
				this.dataSet = res;
			});
	}
}