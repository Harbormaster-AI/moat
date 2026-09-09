import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

 #======================================================================
# 
# Encapsulates data for View TrackingPixel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrackingPixelView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TrackingPixel index.")

def get(request, trackingPixelId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.get( trackingPixelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	trackingPixel = json.loads(request.body)
	delegate = TrackingPixelDelegate()
	responseData = delegate.createFromJson( trackingPixel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	trackingPixel = json.loads(request.body)
	delegate = TrackingPixelDelegate()
	responseData = delegate.save( trackingPixel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, trackingPixelId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.delete( trackingPixelId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TrackingPixelDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, trackingPixelId, CampaignId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.saveCampaign( trackingPixelId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, trackingPixelId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.deleteCampaign( trackingPixelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdvertiser( request, trackingPixelId, AdvertiserId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.saveAdvertiser( trackingPixelId, AdvertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdvertiser( request, trackingPixelId ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.deleteAdvertiser( trackingPixelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConversionEvents( request, trackingPixelId, ConversionEventsIds ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.addConversionEvents( trackingPixelId, ConversionEventsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConversionEvents( request, trackingPixelId, ConversionEventsIds ):
	delegate = TrackingPixelDelegate()
	responseData = delegate.removeConversionEvents( trackingPixelId, ConversionEventsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

