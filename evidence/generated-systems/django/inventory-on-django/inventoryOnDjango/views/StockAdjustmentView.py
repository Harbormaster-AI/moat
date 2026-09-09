import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.StockAdjustmentDelegate import StockAdjustmentDelegate

 #======================================================================
# 
# Encapsulates data for View StockAdjustment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the StockAdjustment index.")

def get(request, stockAdjustmentId ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.get( stockAdjustmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	stockAdjustment = json.loads(request.body)
	delegate = StockAdjustmentDelegate()
	responseData = delegate.createFromJson( stockAdjustment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	stockAdjustment = json.loads(request.body)
	delegate = StockAdjustmentDelegate()
	responseData = delegate.save( stockAdjustment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, stockAdjustmentId ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.delete( stockAdjustmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, stockAdjustmentId, WarehouseId ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.saveWarehouse( stockAdjustmentId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, stockAdjustmentId ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.deleteWarehouse( stockAdjustmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, stockAdjustmentId, LinesIds ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.addLines( stockAdjustmentId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, stockAdjustmentId, LinesIds ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.removeLines( stockAdjustmentId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, stockAdjustmentId, TransactionsIds ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.addTransactions( stockAdjustmentId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, stockAdjustmentId, TransactionsIds ):
	delegate = StockAdjustmentDelegate()
	responseData = delegate.removeTransactions( stockAdjustmentId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

