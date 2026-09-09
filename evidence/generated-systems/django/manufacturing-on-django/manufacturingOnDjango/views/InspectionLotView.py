import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.InspectionLotDelegate import InspectionLotDelegate

 #======================================================================
# 
# Encapsulates data for View InspectionLot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionLotView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InspectionLot index.")

def get(request, inspectionLotId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.get( inspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inspectionLot = json.loads(request.body)
	delegate = InspectionLotDelegate()
	responseData = delegate.createFromJson( inspectionLot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inspectionLot = json.loads(request.body)
	delegate = InspectionLotDelegate()
	responseData = delegate.save( inspectionLot )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inspectionLotId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.delete( inspectionLotId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InspectionLotDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, inspectionLotId, ItemId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.saveItem( inspectionLotId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, inspectionLotId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.deleteItem( inspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkOrder( request, inspectionLotId, WorkOrderId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.saveWorkOrder( inspectionLotId, WorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkOrder( request, inspectionLotId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.deleteWorkOrder( inspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignGoodsReceipt( request, inspectionLotId, GoodsReceiptId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.saveGoodsReceipt( inspectionLotId, GoodsReceiptId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignGoodsReceipt( request, inspectionLotId ):
	delegate = InspectionLotDelegate()
	responseData = delegate.deleteGoodsReceipt( inspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addResults( request, inspectionLotId, ResultsIds ):
	delegate = InspectionLotDelegate()
	responseData = delegate.addResults( inspectionLotId, ResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeResults( request, inspectionLotId, ResultsIds ):
	delegate = InspectionLotDelegate()
	responseData = delegate.removeResults( inspectionLotId, ResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

