import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

 #======================================================================
# 
# Encapsulates data for View Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Report index.")

def get(request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.get( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	report = json.loads(request.body)
	delegate = ReportDelegate()
	responseData = delegate.createFromJson( report )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	report = json.loads(request.body)
	delegate = ReportDelegate()
	responseData = delegate.save( report )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.delete( reportId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReportDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, reportId, WorkspaceId ):
	delegate = ReportDelegate()
	responseData = delegate.saveWorkspace( reportId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.deleteWorkspace( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVisualizations( request, reportId, VisualizationsIds ):
	delegate = ReportDelegate()
	responseData = delegate.addVisualizations( reportId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVisualizations( request, reportId, VisualizationsIds ):
	delegate = ReportDelegate()
	responseData = delegate.removeVisualizations( reportId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, reportId, DatasetsIds ):
	delegate = ReportDelegate()
	responseData = delegate.addDatasets( reportId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, reportId, DatasetsIds ):
	delegate = ReportDelegate()
	responseData = delegate.removeDatasets( reportId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSemanticModels( request, reportId, SemanticModelsIds ):
	delegate = ReportDelegate()
	responseData = delegate.addSemanticModels( reportId, SemanticModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSemanticModels( request, reportId, SemanticModelsIds ):
	delegate = ReportDelegate()
	responseData = delegate.removeSemanticModels( reportId, SemanticModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQueries( request, reportId, QueriesIds ):
	delegate = ReportDelegate()
	responseData = delegate.addQueries( reportId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQueries( request, reportId, QueriesIds ):
	delegate = ReportDelegate()
	responseData = delegate.removeQueries( reportId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTags( request, reportId, TagsIds ):
	delegate = ReportDelegate()
	responseData = delegate.addTags( reportId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTags( request, reportId, TagsIds ):
	delegate = ReportDelegate()
	responseData = delegate.removeTags( reportId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

