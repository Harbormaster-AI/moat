import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.TagDelegate import TagDelegate

 #======================================================================
# 
# Encapsulates data for View Tag
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TagView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Tag index.")

def get(request, tagId ):
	delegate = TagDelegate()
	responseData = delegate.get( tagId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	tag = json.loads(request.body)
	delegate = TagDelegate()
	responseData = delegate.createFromJson( tag )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	tag = json.loads(request.body)
	delegate = TagDelegate()
	responseData = delegate.save( tag )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, tagId ):
	delegate = TagDelegate()
	responseData = delegate.delete( tagId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TagDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, tagId, DatasetsIds ):
	delegate = TagDelegate()
	responseData = delegate.addDatasets( tagId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, tagId, DatasetsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeDatasets( tagId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, tagId, ModelsIds ):
	delegate = TagDelegate()
	responseData = delegate.addModels( tagId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, tagId, ModelsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeModels( tagId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModelVersions( request, tagId, ModelVersionsIds ):
	delegate = TagDelegate()
	responseData = delegate.addModelVersions( tagId, ModelVersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModelVersions( request, tagId, ModelVersionsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeModelVersions( tagId, ModelVersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDashboards( request, tagId, DashboardsIds ):
	delegate = TagDelegate()
	responseData = delegate.addDashboards( tagId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDashboards( request, tagId, DashboardsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeDashboards( tagId, DashboardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, tagId, ReportsIds ):
	delegate = TagDelegate()
	responseData = delegate.addReports( tagId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, tagId, ReportsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeReports( tagId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatureSets( request, tagId, FeatureSetsIds ):
	delegate = TagDelegate()
	responseData = delegate.addFeatureSets( tagId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatureSets( request, tagId, FeatureSetsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeFeatureSets( tagId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMetrics( request, tagId, MetricsIds ):
	delegate = TagDelegate()
	responseData = delegate.addMetrics( tagId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMetrics( request, tagId, MetricsIds ):
	delegate = TagDelegate()
	responseData = delegate.removeMetrics( tagId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

