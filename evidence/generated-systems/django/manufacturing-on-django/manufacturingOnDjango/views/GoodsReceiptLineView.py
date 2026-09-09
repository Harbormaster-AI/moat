import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.GoodsReceiptLineDelegate import GoodsReceiptLineDelegate

 #======================================================================
# 
# Encapsulates data for View GoodsReceiptLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GoodsReceiptLine index.")

def get(request, goodsReceiptLineId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.get( goodsReceiptLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	goodsReceiptLine = json.loads(request.body)
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.createFromJson( goodsReceiptLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	goodsReceiptLine = json.loads(request.body)
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.save( goodsReceiptLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, goodsReceiptLineId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.delete( goodsReceiptLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignGoodsReceipt( request, goodsReceiptLineId, GoodsReceiptId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.saveGoodsReceipt( goodsReceiptLineId, GoodsReceiptId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignGoodsReceipt( request, goodsReceiptLineId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.deleteGoodsReceipt( goodsReceiptLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, goodsReceiptLineId, ItemId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.saveItem( goodsReceiptLineId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, goodsReceiptLineId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.deleteItem( goodsReceiptLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInventoryTransaction( request, goodsReceiptLineId, InventoryTransactionId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.saveInventoryTransaction( goodsReceiptLineId, InventoryTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInventoryTransaction( request, goodsReceiptLineId ):
	delegate = GoodsReceiptLineDelegate()
	responseData = delegate.deleteInventoryTransaction( goodsReceiptLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

