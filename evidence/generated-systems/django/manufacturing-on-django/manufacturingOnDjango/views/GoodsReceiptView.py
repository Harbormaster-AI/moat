import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

 #======================================================================
# 
# Encapsulates data for View GoodsReceipt
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GoodsReceipt index.")

def get(request, goodsReceiptId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.get( goodsReceiptId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	goodsReceipt = json.loads(request.body)
	delegate = GoodsReceiptDelegate()
	responseData = delegate.createFromJson( goodsReceipt )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	goodsReceipt = json.loads(request.body)
	delegate = GoodsReceiptDelegate()
	responseData = delegate.save( goodsReceipt )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, goodsReceiptId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.delete( goodsReceiptId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPurchaseOrder( request, goodsReceiptId, PurchaseOrderId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.savePurchaseOrder( goodsReceiptId, PurchaseOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPurchaseOrder( request, goodsReceiptId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.deletePurchaseOrder( goodsReceiptId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, goodsReceiptId, WarehouseId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.saveWarehouse( goodsReceiptId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, goodsReceiptId ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.deleteWarehouse( goodsReceiptId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, goodsReceiptId, LinesIds ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.addLines( goodsReceiptId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, goodsReceiptId, LinesIds ):
	delegate = GoodsReceiptDelegate()
	responseData = delegate.removeLines( goodsReceiptId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

