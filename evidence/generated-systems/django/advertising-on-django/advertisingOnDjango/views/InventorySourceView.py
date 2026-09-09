import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

 #======================================================================
# 
# Encapsulates data for View InventorySource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventorySourceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InventorySource index.")

def get(request, inventorySourceId ):
	delegate = InventorySourceDelegate()
	responseData = delegate.get( inventorySourceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inventorySource = json.loads(request.body)
	delegate = InventorySourceDelegate()
	responseData = delegate.createFromJson( inventorySource )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inventorySource = json.loads(request.body)
	delegate = InventorySourceDelegate()
	responseData = delegate.save( inventorySource )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inventorySourceId ):
	delegate = InventorySourceDelegate()
	responseData = delegate.delete( inventorySourceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InventorySourceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPublisher( request, inventorySourceId, PublisherId ):
	delegate = InventorySourceDelegate()
	responseData = delegate.savePublisher( inventorySourceId, PublisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPublisher( request, inventorySourceId ):
	delegate = InventorySourceDelegate()
	responseData = delegate.deletePublisher( inventorySourceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdSlots( request, inventorySourceId, AdSlotsIds ):
	delegate = InventorySourceDelegate()
	responseData = delegate.addAdSlots( inventorySourceId, AdSlotsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdSlots( request, inventorySourceId, AdSlotsIds ):
	delegate = InventorySourceDelegate()
	responseData = delegate.removeAdSlots( inventorySourceId, AdSlotsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDeals( request, inventorySourceId, DealsIds ):
	delegate = InventorySourceDelegate()
	responseData = delegate.addDeals( inventorySourceId, DealsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDeals( request, inventorySourceId, DealsIds ):
	delegate = InventorySourceDelegate()
	responseData = delegate.removeDeals( inventorySourceId, DealsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

