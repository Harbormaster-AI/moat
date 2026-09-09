import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ReturnItemDelegate import ReturnItemDelegate

 #======================================================================
# 
# Encapsulates data for View ReturnItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ReturnItem index.")

def get(request, returnItemId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.get( returnItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	returnItem = json.loads(request.body)
	delegate = ReturnItemDelegate()
	responseData = delegate.createFromJson( returnItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	returnItem = json.loads(request.body)
	delegate = ReturnItemDelegate()
	responseData = delegate.save( returnItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, returnItemId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.delete( returnItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReturnItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReturnRequest( request, returnItemId, ReturnRequestId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.saveReturnRequest( returnItemId, ReturnRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReturnRequest( request, returnItemId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.deleteReturnRequest( returnItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrderLine( request, returnItemId, OrderLineId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.saveOrderLine( returnItemId, OrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrderLine( request, returnItemId ):
	delegate = ReturnItemDelegate()
	responseData = delegate.deleteOrderLine( returnItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

