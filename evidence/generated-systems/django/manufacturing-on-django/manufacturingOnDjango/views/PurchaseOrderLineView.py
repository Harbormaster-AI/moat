import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.PurchaseOrderLineDelegate import PurchaseOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for View PurchaseOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PurchaseOrderLine index.")

def get(request, purchaseOrderLineId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.get( purchaseOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	purchaseOrderLine = json.loads(request.body)
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.createFromJson( purchaseOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	purchaseOrderLine = json.loads(request.body)
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.save( purchaseOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, purchaseOrderLineId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.delete( purchaseOrderLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPurchaseOrder( request, purchaseOrderLineId, PurchaseOrderId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.savePurchaseOrder( purchaseOrderLineId, PurchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPurchaseOrder( request, purchaseOrderLineId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.deletePurchaseOrder( purchaseOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, purchaseOrderLineId, ItemId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.saveItem( purchaseOrderLineId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, purchaseOrderLineId ):
	delegate = PurchaseOrderLineDelegate()
	responseData = delegate.deleteItem( purchaseOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

