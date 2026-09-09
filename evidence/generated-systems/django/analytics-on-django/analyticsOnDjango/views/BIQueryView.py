import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

 #======================================================================
# 
# Encapsulates data for View BIQuery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BIQueryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BIQuery index.")

def get(request, bIQueryId ):
	delegate = BIQueryDelegate()
	responseData = delegate.get( bIQueryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	bIQuery = json.loads(request.body)
	delegate = BIQueryDelegate()
	responseData = delegate.createFromJson( bIQuery )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	bIQuery = json.loads(request.body)
	delegate = BIQueryDelegate()
	responseData = delegate.save( bIQuery )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, bIQueryId ):
	delegate = BIQueryDelegate()
	responseData = delegate.delete( bIQueryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BIQueryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, bIQueryId, WorkspaceId ):
	delegate = BIQueryDelegate()
	responseData = delegate.saveWorkspace( bIQueryId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, bIQueryId ):
	delegate = BIQueryDelegate()
	responseData = delegate.deleteWorkspace( bIQueryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, bIQueryId, DatasetsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.addDatasets( bIQueryId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, bIQueryId, DatasetsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.removeDatasets( bIQueryId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, bIQueryId, ReportsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.addReports( bIQueryId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, bIQueryId, ReportsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.removeReports( bIQueryId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDashboards( request, bIQueryId, DashboardsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.addDashboards( bIQueryId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDashboards( request, bIQueryId, DashboardsIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.removeDashboards( bIQueryId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotebooks( request, bIQueryId, NotebooksIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.addNotebooks( bIQueryId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotebooks( request, bIQueryId, NotebooksIds ):
	delegate = BIQueryDelegate()
	responseData = delegate.removeNotebooks( bIQueryId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

