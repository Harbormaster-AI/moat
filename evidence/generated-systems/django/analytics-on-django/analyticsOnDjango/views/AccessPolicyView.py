import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.AccessPolicyDelegate import AccessPolicyDelegate

 #======================================================================
# 
# Encapsulates data for View AccessPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccessPolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AccessPolicy index.")

def get(request, accessPolicyId ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.get( accessPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	accessPolicy = json.loads(request.body)
	delegate = AccessPolicyDelegate()
	responseData = delegate.createFromJson( accessPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	accessPolicy = json.loads(request.body)
	delegate = AccessPolicyDelegate()
	responseData = delegate.save( accessPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, accessPolicyId ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.delete( accessPolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AccessPolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, accessPolicyId, WorkspaceId ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.saveWorkspace( accessPolicyId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, accessPolicyId ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.deleteWorkspace( accessPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, accessPolicyId, DatasetsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.addDatasets( accessPolicyId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, accessPolicyId, DatasetsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.removeDatasets( accessPolicyId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDashboards( request, accessPolicyId, DashboardsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.addDashboards( accessPolicyId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDashboards( request, accessPolicyId, DashboardsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.removeDashboards( accessPolicyId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, accessPolicyId, ReportsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.addReports( accessPolicyId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, accessPolicyId, ReportsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.removeReports( accessPolicyId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, accessPolicyId, ModelsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.addModels( accessPolicyId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, accessPolicyId, ModelsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.removeModels( accessPolicyId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatureSets( request, accessPolicyId, FeatureSetsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.addFeatureSets( accessPolicyId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatureSets( request, accessPolicyId, FeatureSetsIds ):
	delegate = AccessPolicyDelegate()
	responseData = delegate.removeFeatureSets( accessPolicyId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

