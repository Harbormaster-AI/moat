import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

 #======================================================================
# 
# Encapsulates data for View PurchaseOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PurchaseOrder index.")

def get(request, purchaseOrderId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.get( purchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	purchaseOrder = json.loads(request.body)
	delegate = PurchaseOrderDelegate()
	responseData = delegate.createFromJson( purchaseOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	purchaseOrder = json.loads(request.body)
	delegate = PurchaseOrderDelegate()
	responseData = delegate.save( purchaseOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, purchaseOrderId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.delete( purchaseOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, purchaseOrderId, SupplierId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.saveSupplier( purchaseOrderId, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, purchaseOrderId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.deleteSupplier( purchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, purchaseOrderId, PlantId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.savePlant( purchaseOrderId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, purchaseOrderId ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.deletePlant( purchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, purchaseOrderId, LinesIds ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.addLines( purchaseOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, purchaseOrderId, LinesIds ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.removeLines( purchaseOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGoodsReceipts( request, purchaseOrderId, GoodsReceiptsIds ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.addGoodsReceipts( purchaseOrderId, GoodsReceiptsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGoodsReceipts( request, purchaseOrderId, GoodsReceiptsIds ):
	delegate = PurchaseOrderDelegate()
	responseData = delegate.removeGoodsReceipts( purchaseOrderId, GoodsReceiptsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

