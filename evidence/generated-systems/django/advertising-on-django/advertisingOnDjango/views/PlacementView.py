import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

 #======================================================================
# 
# Encapsulates data for View Placement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlacementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Placement index.")

def get(request, placementId ):
	delegate = PlacementDelegate()
	responseData = delegate.get( placementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	placement = json.loads(request.body)
	delegate = PlacementDelegate()
	responseData = delegate.createFromJson( placement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	placement = json.loads(request.body)
	delegate = PlacementDelegate()
	responseData = delegate.save( placement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, placementId ):
	delegate = PlacementDelegate()
	responseData = delegate.delete( placementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PlacementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineItem( request, placementId, LineItemId ):
	delegate = PlacementDelegate()
	responseData = delegate.saveLineItem( placementId, LineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineItem( request, placementId ):
	delegate = PlacementDelegate()
	responseData = delegate.deleteLineItem( placementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdSlot( request, placementId, AdSlotId ):
	delegate = PlacementDelegate()
	responseData = delegate.saveAdSlot( placementId, AdSlotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdSlot( request, placementId ):
	delegate = PlacementDelegate()
	responseData = delegate.deleteAdSlot( placementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDeal( request, placementId, DealId ):
	delegate = PlacementDelegate()
	responseData = delegate.saveDeal( placementId, DealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDeal( request, placementId ):
	delegate = PlacementDelegate()
	responseData = delegate.deleteDeal( placementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

