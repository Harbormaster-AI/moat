import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

 #======================================================================
# 
# Encapsulates data for View Visualization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VisualizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Visualization index.")

def get(request, visualizationId ):
	delegate = VisualizationDelegate()
	responseData = delegate.get( visualizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	visualization = json.loads(request.body)
	delegate = VisualizationDelegate()
	responseData = delegate.createFromJson( visualization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	visualization = json.loads(request.body)
	delegate = VisualizationDelegate()
	responseData = delegate.save( visualization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, visualizationId ):
	delegate = VisualizationDelegate()
	responseData = delegate.delete( visualizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = VisualizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDashboard( request, visualizationId, DashboardId ):
	delegate = VisualizationDelegate()
	responseData = delegate.saveDashboard( visualizationId, DashboardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDashboard( request, visualizationId ):
	delegate = VisualizationDelegate()
	responseData = delegate.deleteDashboard( visualizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReport( request, visualizationId, ReportId ):
	delegate = VisualizationDelegate()
	responseData = delegate.saveReport( visualizationId, ReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReport( request, visualizationId ):
	delegate = VisualizationDelegate()
	responseData = delegate.deleteReport( visualizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMetrics( request, visualizationId, MetricsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.addMetrics( visualizationId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMetrics( request, visualizationId, MetricsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.removeMetrics( visualizationId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDimensions( request, visualizationId, DimensionsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.addDimensions( visualizationId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDimensions( request, visualizationId, DimensionsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.removeDimensions( visualizationId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, visualizationId, DatasetsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.addDatasets( visualizationId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, visualizationId, DatasetsIds ):
	delegate = VisualizationDelegate()
	responseData = delegate.removeDatasets( visualizationId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

