import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

 #======================================================================
# 
# Encapsulates data for View AnalyticsWorkspace
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnalyticsWorkspaceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AnalyticsWorkspace index.")

def get(request, analyticsWorkspaceId ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.get( analyticsWorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	analyticsWorkspace = json.loads(request.body)
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.createFromJson( analyticsWorkspace )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	analyticsWorkspace = json.loads(request.body)
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.save( analyticsWorkspace )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, analyticsWorkspaceId ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.delete( analyticsWorkspaceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, analyticsWorkspaceId, DatasetsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addDatasets( analyticsWorkspaceId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, analyticsWorkspaceId, DatasetsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeDatasets( analyticsWorkspaceId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataSources( request, analyticsWorkspaceId, DataSourcesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addDataSources( analyticsWorkspaceId, DataSourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataSources( request, analyticsWorkspaceId, DataSourcesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeDataSources( analyticsWorkspaceId, DataSourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPipelines( request, analyticsWorkspaceId, PipelinesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addPipelines( analyticsWorkspaceId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePipelines( request, analyticsWorkspaceId, PipelinesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removePipelines( analyticsWorkspaceId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDashboards( request, analyticsWorkspaceId, DashboardsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addDashboards( analyticsWorkspaceId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDashboards( request, analyticsWorkspaceId, DashboardsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeDashboards( analyticsWorkspaceId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, analyticsWorkspaceId, ReportsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addReports( analyticsWorkspaceId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, analyticsWorkspaceId, ReportsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeReports( analyticsWorkspaceId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotebooks( request, analyticsWorkspaceId, NotebooksIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addNotebooks( analyticsWorkspaceId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotebooks( request, analyticsWorkspaceId, NotebooksIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeNotebooks( analyticsWorkspaceId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, analyticsWorkspaceId, ModelsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addModels( analyticsWorkspaceId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, analyticsWorkspaceId, ModelsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeModels( analyticsWorkspaceId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatureSets( request, analyticsWorkspaceId, FeatureSetsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addFeatureSets( analyticsWorkspaceId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatureSets( request, analyticsWorkspaceId, FeatureSetsIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeFeatureSets( analyticsWorkspaceId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, analyticsWorkspaceId, PoliciesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addPolicies( analyticsWorkspaceId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, analyticsWorkspaceId, PoliciesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removePolicies( analyticsWorkspaceId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLineageNodes( request, analyticsWorkspaceId, LineageNodesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.addLineageNodes( analyticsWorkspaceId, LineageNodesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLineageNodes( request, analyticsWorkspaceId, LineageNodesIds ):
	delegate = AnalyticsWorkspaceDelegate()
	responseData = delegate.removeLineageNodes( analyticsWorkspaceId, LineageNodesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

