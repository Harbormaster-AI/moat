import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

 #======================================================================
# 
# Encapsulates data for View WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkCenter index.")

def get(request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.get( workCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workCenter = json.loads(request.body)
	delegate = WorkCenterDelegate()
	responseData = delegate.createFromJson( workCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workCenter = json.loads(request.body)
	delegate = WorkCenterDelegate()
	responseData = delegate.save( workCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.delete( workCenterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkCenterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProductionLine( request, workCenterId, ProductionLineId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.saveProductionLine( workCenterId, ProductionLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProductionLine( request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.deleteProductionLine( workCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

