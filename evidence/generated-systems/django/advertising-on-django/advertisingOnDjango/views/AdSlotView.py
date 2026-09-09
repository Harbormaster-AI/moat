import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

 #======================================================================
# 
# Encapsulates data for View AdSlot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdSlotView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AdSlot index.")

def get(request, adSlotId ):
	delegate = AdSlotDelegate()
	responseData = delegate.get( adSlotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	adSlot = json.loads(request.body)
	delegate = AdSlotDelegate()
	responseData = delegate.createFromJson( adSlot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	adSlot = json.loads(request.body)
	delegate = AdSlotDelegate()
	responseData = delegate.save( adSlot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, adSlotId ):
	delegate = AdSlotDelegate()
	responseData = delegate.delete( adSlotId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AdSlotDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInventorySource( request, adSlotId, InventorySourceId ):
	delegate = AdSlotDelegate()
	responseData = delegate.saveInventorySource( adSlotId, InventorySourceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInventorySource( request, adSlotId ):
	delegate = AdSlotDelegate()
	responseData = delegate.deleteInventorySource( adSlotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlacements( request, adSlotId, PlacementsIds ):
	delegate = AdSlotDelegate()
	responseData = delegate.addPlacements( adSlotId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlacements( request, adSlotId, PlacementsIds ):
	delegate = AdSlotDelegate()
	responseData = delegate.removePlacements( adSlotId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRates( request, adSlotId, RatesIds ):
	delegate = AdSlotDelegate()
	responseData = delegate.addRates( adSlotId, RatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRates( request, adSlotId, RatesIds ):
	delegate = AdSlotDelegate()
	responseData = delegate.removeRates( adSlotId, RatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

