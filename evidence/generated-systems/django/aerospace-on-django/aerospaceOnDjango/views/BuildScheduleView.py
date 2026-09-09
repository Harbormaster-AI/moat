import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.BuildScheduleDelegate import BuildScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View BuildSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BuildScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BuildSchedule index.")

def get(request, buildScheduleId ):
	delegate = BuildScheduleDelegate()
	responseData = delegate.get( buildScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	buildSchedule = json.loads(request.body)
	delegate = BuildScheduleDelegate()
	responseData = delegate.createFromJson( buildSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	buildSchedule = json.loads(request.body)
	delegate = BuildScheduleDelegate()
	responseData = delegate.save( buildSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, buildScheduleId ):
	delegate = BuildScheduleDelegate()
	responseData = delegate.delete( buildScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BuildScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProductionOrders( request, buildScheduleId, ProductionOrdersIds ):
	delegate = BuildScheduleDelegate()
	responseData = delegate.addProductionOrders( buildScheduleId, ProductionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProductionOrders( request, buildScheduleId, ProductionOrdersIds ):
	delegate = BuildScheduleDelegate()
	responseData = delegate.removeProductionOrders( buildScheduleId, ProductionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

