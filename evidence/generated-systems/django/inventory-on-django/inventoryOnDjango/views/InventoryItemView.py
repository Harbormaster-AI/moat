import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

 #======================================================================
# 
# Encapsulates data for View InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InventoryItem index.")

def get(request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.get( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inventoryItem = json.loads(request.body)
	delegate = InventoryItemDelegate()
	responseData = delegate.createFromJson( inventoryItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inventoryItem = json.loads(request.body)
	delegate = InventoryItemDelegate()
	responseData = delegate.save( inventoryItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.delete( inventoryItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InventoryItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, inventoryItemId, SkuId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveSku( inventoryItemId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteSku( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, inventoryItemId, WarehouseId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveWarehouse( inventoryItemId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteWarehouse( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, inventoryItemId, LocationId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveLocation( inventoryItemId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteLocation( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, inventoryItemId, LotId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveLot( inventoryItemId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteLot( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, inventoryItemId, SerialNumbersIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.addSerialNumbers( inventoryItemId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, inventoryItemId, SerialNumbersIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.removeSerialNumbers( inventoryItemId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, inventoryItemId, TransactionsIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.addTransactions( inventoryItemId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, inventoryItemId, TransactionsIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.removeTransactions( inventoryItemId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReservations( request, inventoryItemId, ReservationsIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.addReservations( inventoryItemId, ReservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReservations( request, inventoryItemId, ReservationsIds ):
	delegate = InventoryItemDelegate()
	responseData = delegate.removeReservations( inventoryItemId, ReservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

