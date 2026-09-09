import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

 #======================================================================
# 
# Encapsulates data for View SalesOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SalesOrder index.")

def get(request, salesOrderId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.get( salesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	salesOrder = json.loads(request.body)
	delegate = SalesOrderDelegate()
	responseData = delegate.createFromJson( salesOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	salesOrder = json.loads(request.body)
	delegate = SalesOrderDelegate()
	responseData = delegate.save( salesOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, salesOrderId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.delete( salesOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SalesOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, salesOrderId, CustomerId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.saveCustomer( salesOrderId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, salesOrderId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.deleteCustomer( salesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, salesOrderId, PlantId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.savePlant( salesOrderId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, salesOrderId ):
	delegate = SalesOrderDelegate()
	responseData = delegate.deletePlant( salesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, salesOrderId, LinesIds ):
	delegate = SalesOrderDelegate()
	responseData = delegate.addLines( salesOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, salesOrderId, LinesIds ):
	delegate = SalesOrderDelegate()
	responseData = delegate.removeLines( salesOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkOrders( request, salesOrderId, WorkOrdersIds ):
	delegate = SalesOrderDelegate()
	responseData = delegate.addWorkOrders( salesOrderId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkOrders( request, salesOrderId, WorkOrdersIds ):
	delegate = SalesOrderDelegate()
	responseData = delegate.removeWorkOrders( salesOrderId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

