import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

 #======================================================================
# 
# Encapsulates data for View Item
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Item index.")

def get(request, itemId ):
	delegate = ItemDelegate()
	responseData = delegate.get( itemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	item = json.loads(request.body)
	delegate = ItemDelegate()
	responseData = delegate.createFromJson( item )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	item = json.loads(request.body)
	delegate = ItemDelegate()
	responseData = delegate.save( item )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, itemId ):
	delegate = ItemDelegate()
	responseData = delegate.delete( itemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBusinessUnit( request, itemId, BusinessUnitId ):
	delegate = ItemDelegate()
	responseData = delegate.saveBusinessUnit( itemId, BusinessUnitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBusinessUnit( request, itemId ):
	delegate = ItemDelegate()
	responseData = delegate.deleteBusinessUnit( itemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBoms( request, itemId, BomsIds ):
	delegate = ItemDelegate()
	responseData = delegate.addBoms( itemId, BomsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBoms( request, itemId, BomsIds ):
	delegate = ItemDelegate()
	responseData = delegate.removeBoms( itemId, BomsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRoutings( request, itemId, RoutingsIds ):
	delegate = ItemDelegate()
	responseData = delegate.addRoutings( itemId, RoutingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRoutings( request, itemId, RoutingsIds ):
	delegate = ItemDelegate()
	responseData = delegate.removeRoutings( itemId, RoutingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSuppliers( request, itemId, SuppliersIds ):
	delegate = ItemDelegate()
	responseData = delegate.addSuppliers( itemId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSuppliers( request, itemId, SuppliersIds ):
	delegate = ItemDelegate()
	responseData = delegate.removeSuppliers( itemId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQualitySpecifications( request, itemId, QualitySpecificationsIds ):
	delegate = ItemDelegate()
	responseData = delegate.addQualitySpecifications( itemId, QualitySpecificationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQualitySpecifications( request, itemId, QualitySpecificationsIds ):
	delegate = ItemDelegate()
	responseData = delegate.removeQualitySpecifications( itemId, QualitySpecificationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, itemId, InventoryItemsIds ):
	delegate = ItemDelegate()
	responseData = delegate.addInventoryItems( itemId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, itemId, InventoryItemsIds ):
	delegate = ItemDelegate()
	responseData = delegate.removeInventoryItems( itemId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

