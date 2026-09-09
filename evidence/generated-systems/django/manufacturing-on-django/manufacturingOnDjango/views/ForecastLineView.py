import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ForecastLineDelegate import ForecastLineDelegate

 #======================================================================
# 
# Encapsulates data for View ForecastLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ForecastLine index.")

def get(request, forecastLineId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.get( forecastLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	forecastLine = json.loads(request.body)
	delegate = ForecastLineDelegate()
	responseData = delegate.createFromJson( forecastLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	forecastLine = json.loads(request.body)
	delegate = ForecastLineDelegate()
	responseData = delegate.save( forecastLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, forecastLineId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.delete( forecastLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ForecastLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignForecast( request, forecastLineId, ForecastId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.saveForecast( forecastLineId, ForecastId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignForecast( request, forecastLineId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.deleteForecast( forecastLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, forecastLineId, ItemId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.saveItem( forecastLineId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, forecastLineId ):
	delegate = ForecastLineDelegate()
	responseData = delegate.deleteItem( forecastLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

