import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.MRPRunDelegate import MRPRunDelegate

 #======================================================================
# 
# Encapsulates data for View MRPRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MRPRunView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MRPRun index.")

def get(request, mRPRunId ):
	delegate = MRPRunDelegate()
	responseData = delegate.get( mRPRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	mRPRun = json.loads(request.body)
	delegate = MRPRunDelegate()
	responseData = delegate.createFromJson( mRPRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	mRPRun = json.loads(request.body)
	delegate = MRPRunDelegate()
	responseData = delegate.save( mRPRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, mRPRunId ):
	delegate = MRPRunDelegate()
	responseData = delegate.delete( mRPRunId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MRPRunDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, mRPRunId, PlantId ):
	delegate = MRPRunDelegate()
	responseData = delegate.savePlant( mRPRunId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, mRPRunId ):
	delegate = MRPRunDelegate()
	responseData = delegate.deletePlant( mRPRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlannedOrders( request, mRPRunId, PlannedOrdersIds ):
	delegate = MRPRunDelegate()
	responseData = delegate.addPlannedOrders( mRPRunId, PlannedOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlannedOrders( request, mRPRunId, PlannedOrdersIds ):
	delegate = MRPRunDelegate()
	responseData = delegate.removePlannedOrders( mRPRunId, PlannedOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

