import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

 #======================================================================
# 
# Encapsulates data for View DataPipeline
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataPipelineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataPipeline index.")

def get(request, dataPipelineId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.get( dataPipelineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataPipeline = json.loads(request.body)
	delegate = DataPipelineDelegate()
	responseData = delegate.createFromJson( dataPipeline )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataPipeline = json.loads(request.body)
	delegate = DataPipelineDelegate()
	responseData = delegate.save( dataPipeline )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataPipelineId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.delete( dataPipelineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataPipelineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, dataPipelineId, WorkspaceId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.saveWorkspace( dataPipelineId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, dataPipelineId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.deleteWorkspace( dataPipelineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineageNode( request, dataPipelineId, LineageNodeId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.saveLineageNode( dataPipelineId, LineageNodeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineageNode( request, dataPipelineId ):
	delegate = DataPipelineDelegate()
	responseData = delegate.deleteLineageNode( dataPipelineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTasks( request, dataPipelineId, TasksIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.addTasks( dataPipelineId, TasksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTasks( request, dataPipelineId, TasksIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.removeTasks( dataPipelineId, TasksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSources( request, dataPipelineId, SourcesIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.addSources( dataPipelineId, SourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSources( request, dataPipelineId, SourcesIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.removeSources( dataPipelineId, SourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOutputs( request, dataPipelineId, OutputsIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.addOutputs( dataPipelineId, OutputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOutputs( request, dataPipelineId, OutputsIds ):
	delegate = DataPipelineDelegate()
	responseData = delegate.removeOutputs( dataPipelineId, OutputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

