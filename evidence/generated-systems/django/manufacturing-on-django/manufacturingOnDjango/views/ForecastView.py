import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ForecastDelegate import ForecastDelegate

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

def addLines( request, forecastId, LinesIds ):
	delegate = ForecastDelegate()
	responseData = delegate.addLines( forecastId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, forecastId, LinesIds ):
	delegate = ForecastDelegate()
	responseData = delegate.removeLines( forecastId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

