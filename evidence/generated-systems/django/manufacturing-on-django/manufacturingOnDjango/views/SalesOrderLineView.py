import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.SalesOrderLineDelegate import SalesOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for View SalesOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SalesOrderLine index.")

def get(request, salesOrderLineId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.get( salesOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	salesOrderLine = json.loads(request.body)
	delegate = SalesOrderLineDelegate()
	responseData = delegate.createFromJson( salesOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	salesOrderLine = json.loads(request.body)
	delegate = SalesOrderLineDelegate()
	responseData = delegate.save( salesOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, salesOrderLineId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.delete( salesOrderLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSalesOrder( request, salesOrderLineId, SalesOrderId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.saveSalesOrder( salesOrderLineId, SalesOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSalesOrder( request, salesOrderLineId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.deleteSalesOrder( salesOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, salesOrderLineId, ItemId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.saveItem( salesOrderLineId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, salesOrderLineId ):
	delegate = SalesOrderLineDelegate()
	responseData = delegate.deleteItem( salesOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

