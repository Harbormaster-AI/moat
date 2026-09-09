import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.RunMetricDelegate import RunMetricDelegate

 #======================================================================
# 
# Encapsulates data for View RunMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunMetricView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RunMetric index.")

def get(request, runMetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.get( runMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	runMetric = json.loads(request.body)
	delegate = RunMetricDelegate()
	responseData = delegate.createFromJson( runMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	runMetric = json.loads(request.body)
	delegate = RunMetricDelegate()
	responseData = delegate.save( runMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, runMetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.delete( runMetricId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RunMetricDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTrainingRun( request, runMetricId, TrainingRunId ):
	delegate = RunMetricDelegate()
	responseData = delegate.saveTrainingRun( runMetricId, TrainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTrainingRun( request, runMetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.deleteTrainingRun( runMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMetric( request, runMetricId, MetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.saveMetric( runMetricId, MetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMetric( request, runMetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.deleteMetric( runMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, runMetricId, DatasetId ):
	delegate = RunMetricDelegate()
	responseData = delegate.saveDataset( runMetricId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, runMetricId ):
	delegate = RunMetricDelegate()
	responseData = delegate.deleteDataset( runMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

