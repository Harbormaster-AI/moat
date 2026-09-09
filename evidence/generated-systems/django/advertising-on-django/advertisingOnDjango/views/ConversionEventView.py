import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.ConversionEventDelegate import ConversionEventDelegate

 #======================================================================
# 
# Encapsulates data for View ConversionEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConversionEventView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ConversionEvent index.")

def get(request, conversionEventId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.get( conversionEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	conversionEvent = json.loads(request.body)
	delegate = ConversionEventDelegate()
	responseData = delegate.createFromJson( conversionEvent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	conversionEvent = json.loads(request.body)
	delegate = ConversionEventDelegate()
	responseData = delegate.save( conversionEvent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, conversionEventId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.delete( conversionEventId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ConversionEventDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, conversionEventId, CampaignId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.saveCampaign( conversionEventId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, conversionEventId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.deleteCampaign( conversionEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineItem( request, conversionEventId, LineItemId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.saveLineItem( conversionEventId, LineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineItem( request, conversionEventId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.deleteLineItem( conversionEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTrackingPixel( request, conversionEventId, TrackingPixelId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.saveTrackingPixel( conversionEventId, TrackingPixelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTrackingPixel( request, conversionEventId ):
	delegate = ConversionEventDelegate()
	responseData = delegate.deleteTrackingPixel( conversionEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

