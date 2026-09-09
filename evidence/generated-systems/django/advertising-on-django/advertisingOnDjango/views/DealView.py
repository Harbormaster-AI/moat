import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.DealDelegate import DealDelegate

 #======================================================================
# 
# Encapsulates data for View Deal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DealView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Deal index.")

def get(request, dealId ):
	delegate = DealDelegate()
	responseData = delegate.get( dealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	deal = json.loads(request.body)
	delegate = DealDelegate()
	responseData = delegate.createFromJson( deal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	deal = json.loads(request.body)
	delegate = DealDelegate()
	responseData = delegate.save( deal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dealId ):
	delegate = DealDelegate()
	responseData = delegate.delete( dealId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DealDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPublisher( request, dealId, PublisherId ):
	delegate = DealDelegate()
	responseData = delegate.savePublisher( dealId, PublisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPublisher( request, dealId ):
	delegate = DealDelegate()
	responseData = delegate.deletePublisher( dealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventorySources( request, dealId, InventorySourcesIds ):
	delegate = DealDelegate()
	responseData = delegate.addInventorySources( dealId, InventorySourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventorySources( request, dealId, InventorySourcesIds ):
	delegate = DealDelegate()
	responseData = delegate.removeInventorySources( dealId, InventorySourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlacements( request, dealId, PlacementsIds ):
	delegate = DealDelegate()
	responseData = delegate.addPlacements( dealId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlacements( request, dealId, PlacementsIds ):
	delegate = DealDelegate()
	responseData = delegate.removePlacements( dealId, PlacementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

