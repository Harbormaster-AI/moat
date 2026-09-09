import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

 #======================================================================
# 
# Encapsulates data for View Dashboard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DashboardView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Dashboard index.")

def get(request, dashboardId ):
	delegate = DashboardDelegate()
	responseData = delegate.get( dashboardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dashboard = json.loads(request.body)
	delegate = DashboardDelegate()
	responseData = delegate.createFromJson( dashboard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dashboard = json.loads(request.body)
	delegate = DashboardDelegate()
	responseData = delegate.save( dashboard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dashboardId ):
	delegate = DashboardDelegate()
	responseData = delegate.delete( dashboardId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DashboardDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, dashboardId, WorkspaceId ):
	delegate = DashboardDelegate()
	responseData = delegate.saveWorkspace( dashboardId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, dashboardId ):
	delegate = DashboardDelegate()
	responseData = delegate.deleteWorkspace( dashboardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVisualizations( request, dashboardId, VisualizationsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addVisualizations( dashboardId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVisualizations( request, dashboardId, VisualizationsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeVisualizations( dashboardId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, dashboardId, ReportsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addReports( dashboardId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, dashboardId, ReportsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeReports( dashboardId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, dashboardId, DatasetsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addDatasets( dashboardId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, dashboardId, DatasetsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeDatasets( dashboardId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, dashboardId, AlertsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addAlerts( dashboardId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, dashboardId, AlertsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeAlerts( dashboardId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQueries( request, dashboardId, QueriesIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addQueries( dashboardId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQueries( request, dashboardId, QueriesIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeQueries( dashboardId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTags( request, dashboardId, TagsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.addTags( dashboardId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTags( request, dashboardId, TagsIds ):
	delegate = DashboardDelegate()
	responseData = delegate.removeTags( dashboardId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

