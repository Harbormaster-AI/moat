import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

 #======================================================================
# 
# Encapsulates data for View FeatureSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureSetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FeatureSet index.")

def get(request, featureSetId ):
	delegate = FeatureSetDelegate()
	responseData = delegate.get( featureSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	featureSet = json.loads(request.body)
	delegate = FeatureSetDelegate()
	responseData = delegate.createFromJson( featureSet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	featureSet = json.loads(request.body)
	delegate = FeatureSetDelegate()
	responseData = delegate.save( featureSet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, featureSetId ):
	delegate = FeatureSetDelegate()
	responseData = delegate.delete( featureSetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FeatureSetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, featureSetId, WorkspaceId ):
	delegate = FeatureSetDelegate()
	responseData = delegate.saveWorkspace( featureSetId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, featureSetId ):
	delegate = FeatureSetDelegate()
	responseData = delegate.deleteWorkspace( featureSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatures( request, featureSetId, FeaturesIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.addFeatures( featureSetId, FeaturesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatures( request, featureSetId, FeaturesIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.removeFeatures( featureSetId, FeaturesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, featureSetId, DatasetsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.addDatasets( featureSetId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, featureSetId, DatasetsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.removeDatasets( featureSetId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, featureSetId, ModelsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.addModels( featureSetId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, featureSetId, ModelsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.removeModels( featureSetId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModelVersions( request, featureSetId, ModelVersionsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.addModelVersions( featureSetId, ModelVersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModelVersions( request, featureSetId, ModelVersionsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.removeModelVersions( featureSetId, ModelVersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTags( request, featureSetId, TagsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.addTags( featureSetId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTags( request, featureSetId, TagsIds ):
	delegate = FeatureSetDelegate()
	responseData = delegate.removeTags( featureSetId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

