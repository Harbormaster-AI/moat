import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.SecurityDelegate import SecurityDelegate

 #======================================================================
# 
# Encapsulates data for View Security
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SecurityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Security index.")

def get(request, securityId ):
	delegate = SecurityDelegate()
	responseData = delegate.get( securityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	security = json.loads(request.body)
	delegate = SecurityDelegate()
	responseData = delegate.createFromJson( security )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	security = json.loads(request.body)
	delegate = SecurityDelegate()
	responseData = delegate.save( security )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, securityId ):
	delegate = SecurityDelegate()
	responseData = delegate.delete( securityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SecurityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPositions( request, securityId, PositionsIds ):
	delegate = SecurityDelegate()
	responseData = delegate.addPositions( securityId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePositions( request, securityId, PositionsIds ):
	delegate = SecurityDelegate()
	responseData = delegate.removePositions( securityId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrades( request, securityId, TradesIds ):
	delegate = SecurityDelegate()
	responseData = delegate.addTrades( securityId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrades( request, securityId, TradesIds ):
	delegate = SecurityDelegate()
	responseData = delegate.removeTrades( securityId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, securityId, OrdersIds ):
	delegate = SecurityDelegate()
	responseData = delegate.addOrders( securityId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, securityId, OrdersIds ):
	delegate = SecurityDelegate()
	responseData = delegate.removeOrders( securityId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

