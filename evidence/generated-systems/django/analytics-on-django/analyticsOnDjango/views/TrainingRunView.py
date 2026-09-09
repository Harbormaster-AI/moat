import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

 #======================================================================
# 
# Encapsulates data for View TrainingRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingRunView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TrainingRun index.")

def get(request, trainingRunId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.get( trainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	trainingRun = json.loads(request.body)
	delegate = TrainingRunDelegate()
	responseData = delegate.createFromJson( trainingRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	trainingRun = json.loads(request.body)
	delegate = TrainingRunDelegate()
	responseData = delegate.save( trainingRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, trainingRunId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.delete( trainingRunId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TrainingRunDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignExperiment( request, trainingRunId, ExperimentId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.saveExperiment( trainingRunId, ExperimentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignExperiment( request, trainingRunId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.deleteExperiment( trainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, trainingRunId, ModelVersionId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.saveModelVersion( trainingRunId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, trainingRunId ):
	delegate = TrainingRunDelegate()
	responseData = delegate.deleteModelVersion( trainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInputDatasets( request, trainingRunId, InputDatasetsIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.addInputDatasets( trainingRunId, InputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInputDatasets( request, trainingRunId, InputDatasetsIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.removeInputDatasets( trainingRunId, InputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeatures( request, trainingRunId, FeaturesIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.addFeatures( trainingRunId, FeaturesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeatures( request, trainingRunId, FeaturesIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.removeFeatures( trainingRunId, FeaturesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRunMetrics( request, trainingRunId, RunMetricsIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.addRunMetrics( trainingRunId, RunMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRunMetrics( request, trainingRunId, RunMetricsIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.removeRunMetrics( trainingRunId, RunMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRunParameters( request, trainingRunId, RunParametersIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.addRunParameters( trainingRunId, RunParametersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRunParameters( request, trainingRunId, RunParametersIds ):
	delegate = TrainingRunDelegate()
	responseData = delegate.removeRunParameters( trainingRunId, RunParametersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

