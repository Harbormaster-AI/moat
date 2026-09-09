import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.OperationDelegate import OperationDelegate

 #======================================================================
# 
# Encapsulates data for View Operation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Operation index.")

def get(request, operationId ):
	delegate = OperationDelegate()
	responseData = delegate.get( operationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	operation = json.loads(request.body)
	delegate = OperationDelegate()
	responseData = delegate.createFromJson( operation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	operation = json.loads(request.body)
	delegate = OperationDelegate()
	responseData = delegate.save( operation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, operationId ):
	delegate = OperationDelegate()
	responseData = delegate.delete( operationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OperationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRouting( request, operationId, RoutingId ):
	delegate = OperationDelegate()
	responseData = delegate.saveRouting( operationId, RoutingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRouting( request, operationId ):
	delegate = OperationDelegate()
	responseData = delegate.deleteRouting( operationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkCenter( request, operationId, WorkCenterId ):
	delegate = OperationDelegate()
	responseData = delegate.saveWorkCenter( operationId, WorkCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkCenter( request, operationId ):
	delegate = OperationDelegate()
	responseData = delegate.deleteWorkCenter( operationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInspectionPlan( request, operationId, InspectionPlanId ):
	delegate = OperationDelegate()
	responseData = delegate.saveInspectionPlan( operationId, InspectionPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInspectionPlan( request, operationId ):
	delegate = OperationDelegate()
	responseData = delegate.deleteInspectionPlan( operationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

