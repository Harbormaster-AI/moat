import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.StockAdjustmentLineDelegate import StockAdjustmentLineDelegate

 #======================================================================
# 
# Encapsulates data for View StockAdjustmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the StockAdjustmentLine index.")

def get(request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.get( stockAdjustmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	stockAdjustmentLine = json.loads(request.body)
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.createFromJson( stockAdjustmentLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	stockAdjustmentLine = json.loads(request.body)
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.save( stockAdjustmentLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.delete( stockAdjustmentLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdjustment( request, stockAdjustmentLineId, AdjustmentId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.saveAdjustment( stockAdjustmentLineId, AdjustmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdjustment( request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.deleteAdjustment( stockAdjustmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, stockAdjustmentLineId, SkuId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.saveSku( stockAdjustmentLineId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.deleteSku( stockAdjustmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, stockAdjustmentLineId, LotId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.saveLot( stockAdjustmentLineId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.deleteLot( stockAdjustmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, stockAdjustmentLineId, LocationId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.saveLocation( stockAdjustmentLineId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, stockAdjustmentLineId ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.deleteLocation( stockAdjustmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, stockAdjustmentLineId, SerialNumbersIds ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.addSerialNumbers( stockAdjustmentLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, stockAdjustmentLineId, SerialNumbersIds ):
	delegate = StockAdjustmentLineDelegate()
	responseData = delegate.removeSerialNumbers( stockAdjustmentLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

