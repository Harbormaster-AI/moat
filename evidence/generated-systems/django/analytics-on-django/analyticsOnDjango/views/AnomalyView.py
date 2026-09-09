import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

 #======================================================================
# 
# Encapsulates data for View Anomaly
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnomalyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Anomaly index.")

def get(request, anomalyId ):
	delegate = AnomalyDelegate()
	responseData = delegate.get( anomalyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	anomaly = json.loads(request.body)
	delegate = AnomalyDelegate()
	responseData = delegate.createFromJson( anomaly )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	anomaly = json.loads(request.body)
	delegate = AnomalyDelegate()
	responseData = delegate.save( anomaly )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, anomalyId ):
	delegate = AnomalyDelegate()
	responseData = delegate.delete( anomalyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AnomalyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTimeSeries( request, anomalyId, TimeSeriesId ):
	delegate = AnomalyDelegate()
	responseData = delegate.saveTimeSeries( anomalyId, TimeSeriesId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTimeSeries( request, anomalyId ):
	delegate = AnomalyDelegate()
	responseData = delegate.deleteTimeSeries( anomalyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAlert( request, anomalyId, AlertId ):
	delegate = AnomalyDelegate()
	responseData = delegate.saveAlert( anomalyId, AlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAlert( request, anomalyId ):
	delegate = AnomalyDelegate()
	responseData = delegate.deleteAlert( anomalyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, anomalyId, DatasetId ):
	delegate = AnomalyDelegate()
	responseData = delegate.saveDataset( anomalyId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, anomalyId ):
	delegate = AnomalyDelegate()
	responseData = delegate.deleteDataset( anomalyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

