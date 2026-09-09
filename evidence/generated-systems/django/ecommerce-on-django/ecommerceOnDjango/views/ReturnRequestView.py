import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ReturnRequestDelegate import ReturnRequestDelegate

 #======================================================================
# 
# Encapsulates data for View ReturnRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnRequestView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ReturnRequest index.")

def get(request, returnRequestId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.get( returnRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	returnRequest = json.loads(request.body)
	delegate = ReturnRequestDelegate()
	responseData = delegate.createFromJson( returnRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	returnRequest = json.loads(request.body)
	delegate = ReturnRequestDelegate()
	responseData = delegate.save( returnRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, returnRequestId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.delete( returnRequestId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReturnRequestDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, returnRequestId, OrderId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.saveOrder( returnRequestId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, returnRequestId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.deleteOrder( returnRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRefund( request, returnRequestId, RefundId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.saveRefund( returnRequestId, RefundId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRefund( request, returnRequestId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.deleteRefund( returnRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignShipment( request, returnRequestId, ShipmentId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.saveShipment( returnRequestId, ShipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignShipment( request, returnRequestId ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.deleteShipment( returnRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, returnRequestId, ItemsIds ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.addItems( returnRequestId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, returnRequestId, ItemsIds ):
	delegate = ReturnRequestDelegate()
	responseData = delegate.removeItems( returnRequestId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

