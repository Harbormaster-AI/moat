import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

 #======================================================================
# 
# Encapsulates data for View DataSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataSet index.")

def get(request, dataSetId ):
	delegate = DataSetDelegate()
	responseData = delegate.get( dataSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataSet = json.loads(request.body)
	delegate = DataSetDelegate()
	responseData = delegate.createFromJson( dataSet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataSet = json.loads(request.body)
	delegate = DataSetDelegate()
	responseData = delegate.save( dataSet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataSetId ):
	delegate = DataSetDelegate()
	responseData = delegate.delete( dataSetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataSetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, dataSetId, WorkspaceId ):
	delegate = DataSetDelegate()
	responseData = delegate.saveWorkspace( dataSetId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, dataSetId ):
	delegate = DataSetDelegate()
	responseData = delegate.deleteWorkspace( dataSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineageNode( request, dataSetId, LineageNodeId ):
	delegate = DataSetDelegate()
	responseData = delegate.saveLineageNode( dataSetId, LineageNodeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineageNode( request, dataSetId ):
	delegate = DataSetDelegate()
	responseData = delegate.deleteLineageNode( dataSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSources( request, dataSetId, SourcesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addSources( dataSetId, SourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSources( request, dataSetId, SourcesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeSources( dataSetId, SourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPipelines( request, dataSetId, PipelinesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addPipelines( dataSetId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePipelines( request, dataSetId, PipelinesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removePipelines( dataSetId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSemanticModels( request, dataSetId, SemanticModelsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addSemanticModels( dataSetId, SemanticModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSemanticModels( request, dataSetId, SemanticModelsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeSemanticModels( dataSetId, SemanticModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDimensions( request, dataSetId, DimensionsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addDimensions( dataSetId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDimensions( request, dataSetId, DimensionsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeDimensions( dataSetId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMeasures( request, dataSetId, MeasuresIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addMeasures( dataSetId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMeasures( request, dataSetId, MeasuresIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeMeasures( dataSetId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMetrics( request, dataSetId, MetricsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addMetrics( dataSetId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMetrics( request, dataSetId, MetricsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeMetrics( dataSetId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQualityRules( request, dataSetId, QualityRulesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addQualityRules( dataSetId, QualityRulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQualityRules( request, dataSetId, QualityRulesIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeQualityRules( dataSetId, QualityRulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTags( request, dataSetId, TagsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.addTags( dataSetId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTags( request, dataSetId, TagsIds ):
	delegate = DataSetDelegate()
	responseData = delegate.removeTags( dataSetId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

