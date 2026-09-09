import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.LotDelegate import LotDelegate

 #======================================================================
# 
# Encapsulates data for View Lot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LotView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Lot index.")

def get(request, lotId ):
	delegate = LotDelegate()
	responseData = delegate.get( lotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	lot = json.loads(request.body)
	delegate = LotDelegate()
	responseData = delegate.createFromJson( lot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	lot = json.loads(request.body)
	delegate = LotDelegate()
	responseData = delegate.save( lot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, lotId ):
	delegate = LotDelegate()
	responseData = delegate.delete( lotId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LotDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, lotId, SkuId ):
	delegate = LotDelegate()
	responseData = delegate.saveSku( lotId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, lotId ):
	delegate = LotDelegate()
	responseData = delegate.deleteSku( lotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, lotId, InventoryItemsIds ):
	delegate = LotDelegate()
	responseData = delegate.addInventoryItems( lotId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, lotId, InventoryItemsIds ):
	delegate = LotDelegate()
	responseData = delegate.removeInventoryItems( lotId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

