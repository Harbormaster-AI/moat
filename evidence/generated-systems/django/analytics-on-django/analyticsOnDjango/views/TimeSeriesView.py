import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.TimeSeriesDelegate import TimeSeriesDelegate

 #======================================================================
# 
# Encapsulates data for View TimeSeries
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeSeriesView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TimeSeries index.")

def get(request, timeSeriesId ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.get( timeSeriesId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	timeSeries = json.loads(request.body)
	delegate = TimeSeriesDelegate()
	responseData = delegate.createFromJson( timeSeries )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	timeSeries = json.loads(request.body)
	delegate = TimeSeriesDelegate()
	responseData = delegate.save( timeSeries )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, timeSeriesId ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.delete( timeSeriesId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TimeSeriesDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, timeSeriesId, DatasetsIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.addDatasets( timeSeriesId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, timeSeriesId, DatasetsIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.removeDatasets( timeSeriesId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addForecasts( request, timeSeriesId, ForecastsIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.addForecasts( timeSeriesId, ForecastsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeForecasts( request, timeSeriesId, ForecastsIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.removeForecasts( timeSeriesId, ForecastsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAnomalies( request, timeSeriesId, AnomaliesIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.addAnomalies( timeSeriesId, AnomaliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAnomalies( request, timeSeriesId, AnomaliesIds ):
	delegate = TimeSeriesDelegate()
	responseData = delegate.removeAnomalies( timeSeriesId, AnomaliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

