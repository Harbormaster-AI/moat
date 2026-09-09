import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.NonconformanceDelegate import NonconformanceDelegate

 #======================================================================
# 
# Encapsulates data for View Nonconformance
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NonconformanceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Nonconformance index.")

def get(request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.get( nonconformanceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	nonconformance = json.loads(request.body)
	delegate = NonconformanceDelegate()
	responseData = delegate.createFromJson( nonconformance )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	nonconformance = json.loads(request.body)
	delegate = NonconformanceDelegate()
	responseData = delegate.save( nonconformance )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.delete( nonconformanceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = NonconformanceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, nonconformanceId, ItemId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.saveItem( nonconformanceId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.deleteItem( nonconformanceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkOrder( request, nonconformanceId, WorkOrderId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.saveWorkOrder( nonconformanceId, WorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkOrder( request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.deleteWorkOrder( nonconformanceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInspectionLot( request, nonconformanceId, InspectionLotId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.saveInspectionLot( nonconformanceId, InspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInspectionLot( request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.deleteInspectionLot( nonconformanceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCorrectiveAction( request, nonconformanceId, CorrectiveActionId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.saveCorrectiveAction( nonconformanceId, CorrectiveActionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCorrectiveAction( request, nonconformanceId ):
	delegate = NonconformanceDelegate()
	responseData = delegate.deleteCorrectiveAction( nonconformanceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

