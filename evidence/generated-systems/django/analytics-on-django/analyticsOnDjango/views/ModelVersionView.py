import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

 #======================================================================
# 
# Encapsulates data for View ModelVersion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelVersionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ModelVersion index.")

def get(request, modelVersionId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.get( modelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	modelVersion = json.loads(request.body)
	delegate = ModelVersionDelegate()
	responseData = delegate.createFromJson( modelVersion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	modelVersion = json.loads(request.body)
	delegate = ModelVersionDelegate()
	responseData = delegate.save( modelVersion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, modelVersionId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.delete( modelVersionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ModelVersionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModel( request, modelVersionId, ModelId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.saveModel( modelVersionId, ModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModel( request, modelVersionId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.deleteModel( modelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTrainingRun( request, modelVersionId, TrainingRunId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.saveTrainingRun( modelVersionId, TrainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTrainingRun( request, modelVersionId ):
	delegate = ModelVersionDelegate()
	responseData = delegate.deleteTrainingRun( modelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEvaluationMetrics( request, modelVersionId, EvaluationMetricsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.addEvaluationMetrics( modelVersionId, EvaluationMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEvaluationMetrics( request, modelVersionId, EvaluationMetricsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.removeEvaluationMetrics( modelVersionId, EvaluationMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDeployments( request, modelVersionId, DeploymentsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.addDeployments( modelVersionId, DeploymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDeployments( request, modelVersionId, DeploymentsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.removeDeployments( modelVersionId, DeploymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatureSets( request, modelVersionId, FeatureSetsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.addFeatureSets( modelVersionId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatureSets( request, modelVersionId, FeatureSetsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.removeFeatureSets( modelVersionId, FeatureSetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, modelVersionId, DatasetsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.addDatasets( modelVersionId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, modelVersionId, DatasetsIds ):
	delegate = ModelVersionDelegate()
	responseData = delegate.removeDatasets( modelVersionId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

