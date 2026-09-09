import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

 #======================================================================
# 
# Encapsulates data for View LineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LineItem index.")

def get(request, lineItemId ):
	delegate = LineItemDelegate()
	responseData = delegate.get( lineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	lineItem = json.loads(request.body)
	delegate = LineItemDelegate()
	responseData = delegate.createFromJson( lineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	lineItem = json.loads(request.body)
	delegate = LineItemDelegate()
	responseData = delegate.save( lineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, lineItemId ):
	delegate = LineItemDelegate()
	responseData = delegate.delete( lineItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LineItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, lineItemId, CampaignId ):
	delegate = LineItemDelegate()
	responseData = delegate.saveCampaign( lineItemId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, lineItemId ):
	delegate = LineItemDelegate()
	responseData = delegate.deleteCampaign( lineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTargetingProfile( request, lineItemId, TargetingProfileId ):
	delegate = LineItemDelegate()
	responseData = delegate.saveTargetingProfile( lineItemId, TargetingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTargetingProfile( request, lineItemId ):
	delegate = LineItemDelegate()
	responseData = delegate.deleteTargetingProfile( lineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDeal( request, lineItemId, DealId ):
	delegate = LineItemDelegate()
	responseData = delegate.saveDeal( lineItemId, DealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDeal( request, lineItemId ):
	delegate = LineItemDelegate()
	responseData = delegate.deleteDeal( lineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlacements( request, lineItemId, PlacementsIds ):
	delegate = LineItemDelegate()
	responseData = delegate.addPlacements( lineItemId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlacements( request, lineItemId, PlacementsIds ):
	delegate = LineItemDelegate()
	responseData = delegate.removePlacements( lineItemId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCreatives( request, lineItemId, CreativesIds ):
	delegate = LineItemDelegate()
	responseData = delegate.addCreatives( lineItemId, CreativesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCreatives( request, lineItemId, CreativesIds ):
	delegate = LineItemDelegate()
	responseData = delegate.removeCreatives( lineItemId, CreativesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPerformanceMetrics( request, lineItemId, PerformanceMetricsIds ):
	delegate = LineItemDelegate()
	responseData = delegate.addPerformanceMetrics( lineItemId, PerformanceMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePerformanceMetrics( request, lineItemId, PerformanceMetricsIds ):
	delegate = LineItemDelegate()
	responseData = delegate.removePerformanceMetrics( lineItemId, PerformanceMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

