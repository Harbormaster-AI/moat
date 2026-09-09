import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AirworthinessDirectiveDelegate import AirworthinessDirectiveDelegate

 #======================================================================
# 
# Encapsulates data for View AirworthinessDirective
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AirworthinessDirectiveView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AirworthinessDirective index.")

def get(request, airworthinessDirectiveId ):
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.get( airworthinessDirectiveId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	airworthinessDirective = json.loads(request.body)
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.createFromJson( airworthinessDirective )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	airworthinessDirective = json.loads(request.body)
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.save( airworthinessDirective )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, airworthinessDirectiveId ):
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.delete( airworthinessDirectiveId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkOrders( request, airworthinessDirectiveId, WorkOrdersIds ):
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.addWorkOrders( airworthinessDirectiveId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkOrders( request, airworthinessDirectiveId, WorkOrdersIds ):
	delegate = AirworthinessDirectiveDelegate()
	responseData = delegate.removeWorkOrders( airworthinessDirectiveId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

