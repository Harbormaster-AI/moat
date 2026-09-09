import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

 #======================================================================
# 
# Encapsulates data for View Routing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoutingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Routing index.")

def get(request, routingId ):
	delegate = RoutingDelegate()
	responseData = delegate.get( routingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	routing = json.loads(request.body)
	delegate = RoutingDelegate()
	responseData = delegate.createFromJson( routing )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	routing = json.loads(request.body)
	delegate = RoutingDelegate()
	responseData = delegate.save( routing )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, routingId ):
	delegate = RoutingDelegate()
	responseData = delegate.delete( routingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RoutingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, routingId, ItemId ):
	delegate = RoutingDelegate()
	responseData = delegate.saveItem( routingId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, routingId ):
	delegate = RoutingDelegate()
	responseData = delegate.deleteItem( routingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOperations( request, routingId, OperationsIds ):
	delegate = RoutingDelegate()
	responseData = delegate.addOperations( routingId, OperationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOperations( request, routingId, OperationsIds ):
	delegate = RoutingDelegate()
	responseData = delegate.removeOperations( routingId, OperationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

