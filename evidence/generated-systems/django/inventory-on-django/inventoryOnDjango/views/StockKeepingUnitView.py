import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

 #======================================================================
# 
# Encapsulates data for View StockKeepingUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockKeepingUnitView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the StockKeepingUnit index.")

def get(request, stockKeepingUnitId ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.get( stockKeepingUnitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	stockKeepingUnit = json.loads(request.body)
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.createFromJson( stockKeepingUnit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	stockKeepingUnit = json.loads(request.body)
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.save( stockKeepingUnit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, stockKeepingUnitId ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.delete( stockKeepingUnitId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, stockKeepingUnitId, InventoryItemsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.addInventoryItems( stockKeepingUnitId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, stockKeepingUnitId, InventoryItemsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.removeInventoryItems( stockKeepingUnitId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUomConversions( request, stockKeepingUnitId, UomConversionsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.addUomConversions( stockKeepingUnitId, UomConversionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUomConversions( request, stockKeepingUnitId, UomConversionsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.removeUomConversions( stockKeepingUnitId, UomConversionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReplenishmentPolicies( request, stockKeepingUnitId, ReplenishmentPoliciesIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.addReplenishmentPolicies( stockKeepingUnitId, ReplenishmentPoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReplenishmentPolicies( request, stockKeepingUnitId, ReplenishmentPoliciesIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.removeReplenishmentPolicies( stockKeepingUnitId, ReplenishmentPoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLots( request, stockKeepingUnitId, LotsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.addLots( stockKeepingUnitId, LotsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLots( request, stockKeepingUnitId, LotsIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.removeLots( stockKeepingUnitId, LotsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, stockKeepingUnitId, SerialNumbersIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.addSerialNumbers( stockKeepingUnitId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, stockKeepingUnitId, SerialNumbersIds ):
	delegate = StockKeepingUnitDelegate()
	responseData = delegate.removeSerialNumbers( stockKeepingUnitId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

