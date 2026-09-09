import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

 #======================================================================
# 
# Encapsulates data for View LineageNode
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineageNodeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LineageNode index.")

def get(request, lineageNodeId ):
	delegate = LineageNodeDelegate()
	responseData = delegate.get( lineageNodeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	lineageNode = json.loads(request.body)
	delegate = LineageNodeDelegate()
	responseData = delegate.createFromJson( lineageNode )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	lineageNode = json.loads(request.body)
	delegate = LineageNodeDelegate()
	responseData = delegate.save( lineageNode )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, lineageNodeId ):
	delegate = LineageNodeDelegate()
	responseData = delegate.delete( lineageNodeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LineageNodeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, lineageNodeId, WorkspaceId ):
	delegate = LineageNodeDelegate()
	responseData = delegate.saveWorkspace( lineageNodeId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, lineageNodeId ):
	delegate = LineageNodeDelegate()
	responseData = delegate.deleteWorkspace( lineageNodeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInputs( request, lineageNodeId, InputsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addInputs( lineageNodeId, InputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInputs( request, lineageNodeId, InputsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeInputs( lineageNodeId, InputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOutputs( request, lineageNodeId, OutputsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addOutputs( lineageNodeId, OutputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOutputs( request, lineageNodeId, OutputsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeOutputs( lineageNodeId, OutputsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, lineageNodeId, DatasetsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addDatasets( lineageNodeId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, lineageNodeId, DatasetsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeDatasets( lineageNodeId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, lineageNodeId, ModelsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addModels( lineageNodeId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, lineageNodeId, ModelsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeModels( lineageNodeId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPipelines( request, lineageNodeId, PipelinesIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addPipelines( lineageNodeId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePipelines( request, lineageNodeId, PipelinesIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removePipelines( lineageNodeId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDashboards( request, lineageNodeId, DashboardsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addDashboards( lineageNodeId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDashboards( request, lineageNodeId, DashboardsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeDashboards( lineageNodeId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, lineageNodeId, ReportsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.addReports( lineageNodeId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, lineageNodeId, ReportsIds ):
	delegate = LineageNodeDelegate()
	responseData = delegate.removeReports( lineageNodeId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

