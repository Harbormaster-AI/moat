import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

 #======================================================================
# 
# Encapsulates data for View Feature
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Feature index.")

def get(request, featureId ):
	delegate = FeatureDelegate()
	responseData = delegate.get( featureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	feature = json.loads(request.body)
	delegate = FeatureDelegate()
	responseData = delegate.createFromJson( feature )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	feature = json.loads(request.body)
	delegate = FeatureDelegate()
	responseData = delegate.save( feature )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, featureId ):
	delegate = FeatureDelegate()
	responseData = delegate.delete( featureId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FeatureDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFeatureSet( request, featureId, FeatureSetId ):
	delegate = FeatureDelegate()
	responseData = delegate.saveFeatureSet( featureId, FeatureSetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFeatureSet( request, featureId ):
	delegate = FeatureDelegate()
	responseData = delegate.deleteFeatureSet( featureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSourceDatasets( request, featureId, SourceDatasetsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.addSourceDatasets( featureId, SourceDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSourceDatasets( request, featureId, SourceDatasetsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.removeSourceDatasets( featureId, SourceDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, featureId, ModelsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.addModels( featureId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, featureId, ModelsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.removeModels( featureId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrainingRuns( request, featureId, TrainingRunsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.addTrainingRuns( featureId, TrainingRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrainingRuns( request, featureId, TrainingRunsIds ):
	delegate = FeatureDelegate()
	responseData = delegate.removeTrainingRuns( featureId, TrainingRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

