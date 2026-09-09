import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

 #======================================================================
# 
# Encapsulates data for View WorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkOrder index.")

def get(request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.get( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workOrder = json.loads(request.body)
	delegate = WorkOrderDelegate()
	responseData = delegate.createFromJson( workOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workOrder = json.loads(request.body)
	delegate = WorkOrderDelegate()
	responseData = delegate.save( workOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.delete( workOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, workOrderId, ItemId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.saveItem( workOrderId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deleteItem( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, workOrderId, PlantId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.savePlant( workOrderId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deletePlant( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRouting( request, workOrderId, RoutingId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.saveRouting( workOrderId, RoutingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRouting( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deleteRouting( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBom( request, workOrderId, BomId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.saveBom( workOrderId, BomId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBom( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deleteBom( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProductionSchedule( request, workOrderId, ProductionScheduleId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.saveProductionSchedule( workOrderId, ProductionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProductionSchedule( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deleteProductionSchedule( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSalesOrder( request, workOrderId, SalesOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.saveSalesOrder( workOrderId, SalesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSalesOrder( request, workOrderId ):
	delegate = WorkOrderDelegate()
	responseData = delegate.deleteSalesOrder( workOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

