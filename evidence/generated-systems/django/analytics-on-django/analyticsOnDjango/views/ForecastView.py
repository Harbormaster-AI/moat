import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.ForecastDelegate import ForecastDelegate

 #======================================================================
# 
# Encapsulates data for View Forecast
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Forecast index.")

def get(request, forecastId ):
	delegate = ForecastDelegate()
	responseData = delegate.get( forecastId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	forecast = json.loads(request.body)
	delegate = ForecastDelegate()
	responseData = delegate.createFromJson( forecast )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	forecast = json.loads(request.body)
	delegate = ForecastDelegate()
	responseData = delegate.save( forecast )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, forecastId ):
	delegate = ForecastDelegate()
	responseData = delegate.delete( forecastId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ForecastDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, forecastId, ModelVersionId ):
	delegate = ForecastDelegate()
	responseData = delegate.saveModelVersion( forecastId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, forecastId ):
	delegate = ForecastDelegate()
	responseData = delegate.deleteModelVersion( forecastId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTimeSeries( request, forecastId, TimeSeriesId ):
	delegate = ForecastDelegate()
	responseData = delegate.saveTimeSeries( forecastId, TimeSeriesId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTimeSeries( request, forecastId ):
	delegate = ForecastDelegate()
	responseData = delegate.deleteTimeSeries( forecastId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, forecastId, DatasetsIds ):
	delegate = ForecastDelegate()
	responseData = delegate.addDatasets( forecastId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, forecastId, DatasetsIds ):
	delegate = ForecastDelegate()
	responseData = delegate.removeDatasets( forecastId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

