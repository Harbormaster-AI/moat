import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.EvaluationMetricDelegate import EvaluationMetricDelegate

 #======================================================================
# 
# Encapsulates data for View EvaluationMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvaluationMetricView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EvaluationMetric index.")

def get(request, evaluationMetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.get( evaluationMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	evaluationMetric = json.loads(request.body)
	delegate = EvaluationMetricDelegate()
	responseData = delegate.createFromJson( evaluationMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	evaluationMetric = json.loads(request.body)
	delegate = EvaluationMetricDelegate()
	responseData = delegate.save( evaluationMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, evaluationMetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.delete( evaluationMetricId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, evaluationMetricId, ModelVersionId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.saveModelVersion( evaluationMetricId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, evaluationMetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.deleteModelVersion( evaluationMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMetric( request, evaluationMetricId, MetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.saveMetric( evaluationMetricId, MetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMetric( request, evaluationMetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.deleteMetric( evaluationMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, evaluationMetricId, DatasetId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.saveDataset( evaluationMetricId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, evaluationMetricId ):
	delegate = EvaluationMetricDelegate()
	responseData = delegate.deleteDataset( evaluationMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

