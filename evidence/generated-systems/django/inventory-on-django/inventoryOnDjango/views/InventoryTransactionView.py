import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

 #======================================================================
# 
# Encapsulates data for View InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InventoryTransaction index.")

def get(request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.get( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inventoryTransaction = json.loads(request.body)
	delegate = InventoryTransactionDelegate()
	responseData = delegate.createFromJson( inventoryTransaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inventoryTransaction = json.loads(request.body)
	delegate = InventoryTransactionDelegate()
	responseData = delegate.save( inventoryTransaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.delete( inventoryTransactionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, inventoryTransactionId, SkuId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveSku( inventoryTransactionId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteSku( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, inventoryTransactionId, WarehouseId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveWarehouse( inventoryTransactionId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteWarehouse( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, inventoryTransactionId, LocationId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveLocation( inventoryTransactionId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteLocation( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, inventoryTransactionId, LotId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveLot( inventoryTransactionId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteLot( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRelatedReservation( request, inventoryTransactionId, RelatedReservationId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveRelatedReservation( inventoryTransactionId, RelatedReservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRelatedReservation( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteRelatedReservation( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransferOrder( request, inventoryTransactionId, TransferOrderId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveTransferOrder( inventoryTransactionId, TransferOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransferOrder( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteTransferOrder( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdjustment( request, inventoryTransactionId, AdjustmentId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveAdjustment( inventoryTransactionId, AdjustmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdjustment( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteAdjustment( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCycleCount( request, inventoryTransactionId, CycleCountId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveCycleCount( inventoryTransactionId, CycleCountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCycleCount( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteCycleCount( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, inventoryTransactionId, SerialNumbersIds ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.addSerialNumbers( inventoryTransactionId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, inventoryTransactionId, SerialNumbersIds ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.removeSerialNumbers( inventoryTransactionId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

