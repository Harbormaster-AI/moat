import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

 #======================================================================
# 
# Encapsulates data for View Model
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Model index.")

def get(request, modelId ):
	delegate = ModelDelegate()
	responseData = delegate.get( modelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	model = json.loads(request.body)
	delegate = ModelDelegate()
	responseData = delegate.createFromJson( model )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	model = json.loads(request.body)
	delegate = ModelDelegate()
	responseData = delegate.save( model )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, modelId ):
	delegate = ModelDelegate()
	responseData = delegate.delete( modelId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ModelDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, modelId, WorkspaceId ):
	delegate = ModelDelegate()
	responseData = delegate.saveWorkspace( modelId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, modelId ):
	delegate = ModelDelegate()
	responseData = delegate.deleteWorkspace( modelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVersions( request, modelId, VersionsIds ):
	delegate = ModelDelegate()
	responseData = delegate.addVersions( modelId, VersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVersions( request, modelId, VersionsIds ):
	delegate = ModelDelegate()
	responseData = delegate.removeVersions( modelId, VersionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatureSets( request, modelId, FeatureSetsIds ):
	delegate = ModelDelegate()
	responseData = delegate.addFeatureSets( modelId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatureSets( request, modelId, FeatureSetsIds ):
	delegate = ModelDelegate()
	responseData = delegate.removeFeatureSets( modelId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExperiments( request, modelId, ExperimentsIds ):
	delegate = ModelDelegate()
	responseData = delegate.addExperiments( modelId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExperiments( request, modelId, ExperimentsIds ):
	delegate = ModelDelegate()
	responseData = delegate.removeExperiments( modelId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTags( request, modelId, TagsIds ):
	delegate = ModelDelegate()
	responseData = delegate.addTags( modelId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTags( request, modelId, TagsIds ):
	delegate = ModelDelegate()
	responseData = delegate.removeTags( modelId, TagsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

