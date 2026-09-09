import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

 #======================================================================
# 
# Encapsulates data for View SerialNumber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerialNumberView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SerialNumber index.")

def get(request, serialNumberId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.get( serialNumberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	serialNumber = json.loads(request.body)
	delegate = SerialNumberDelegate()
	responseData = delegate.createFromJson( serialNumber )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	serialNumber = json.loads(request.body)
	delegate = SerialNumberDelegate()
	responseData = delegate.save( serialNumber )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, serialNumberId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.delete( serialNumberId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SerialNumberDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, serialNumberId, SkuId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.saveSku( serialNumberId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, serialNumberId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.deleteSku( serialNumberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCurrentInventoryItem( request, serialNumberId, CurrentInventoryItemId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.saveCurrentInventoryItem( serialNumberId, CurrentInventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCurrentInventoryItem( request, serialNumberId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.deleteCurrentInventoryItem( serialNumberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, serialNumberId, LotId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.saveLot( serialNumberId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, serialNumberId ):
	delegate = SerialNumberDelegate()
	responseData = delegate.deleteLot( serialNumberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

