import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

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

def assignItem( request, inventoryTransactionId, ItemId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveItem( inventoryTransactionId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteItem( inventoryTransactionId )
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

def assignWorkOrder( request, inventoryTransactionId, WorkOrderId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveWorkOrder( inventoryTransactionId, WorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkOrder( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteWorkOrder( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPurchaseOrder( request, inventoryTransactionId, PurchaseOrderId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.savePurchaseOrder( inventoryTransactionId, PurchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPurchaseOrder( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deletePurchaseOrder( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSalesOrder( request, inventoryTransactionId, SalesOrderId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.saveSalesOrder( inventoryTransactionId, SalesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSalesOrder( request, inventoryTransactionId ):
	delegate = InventoryTransactionDelegate()
	responseData = delegate.deleteSalesOrder( inventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

