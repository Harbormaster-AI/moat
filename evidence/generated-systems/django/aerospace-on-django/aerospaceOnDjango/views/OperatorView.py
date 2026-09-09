import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

 #======================================================================
# 
# Encapsulates data for View Operator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperatorView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Operator index.")

def get(request, operatorId ):
	delegate = OperatorDelegate()
	responseData = delegate.get( operatorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	operator = json.loads(request.body)
	delegate = OperatorDelegate()
	responseData = delegate.createFromJson( operator )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	operator = json.loads(request.body)
	delegate = OperatorDelegate()
	responseData = delegate.save( operator )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, operatorId ):
	delegate = OperatorDelegate()
	responseData = delegate.delete( operatorId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OperatorDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSalesRegion( request, operatorId, SalesRegionId ):
	delegate = OperatorDelegate()
	responseData = delegate.saveSalesRegion( operatorId, SalesRegionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSalesRegion( request, operatorId ):
	delegate = OperatorDelegate()
	responseData = delegate.deleteSalesRegion( operatorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAircraftOrders( request, operatorId, AircraftOrdersIds ):
	delegate = OperatorDelegate()
	responseData = delegate.addAircraftOrders( operatorId, AircraftOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAircraftOrders( request, operatorId, AircraftOrdersIds ):
	delegate = OperatorDelegate()
	responseData = delegate.removeAircraftOrders( operatorId, AircraftOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOperatedAircraft( request, operatorId, OperatedAircraftIds ):
	delegate = OperatorDelegate()
	responseData = delegate.addOperatedAircraft( operatorId, OperatedAircraftIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOperatedAircraft( request, operatorId, OperatedAircraftIds ):
	delegate = OperatorDelegate()
	responseData = delegate.removeOperatedAircraft( operatorId, OperatedAircraftIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

