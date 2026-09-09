import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

 #======================================================================
# 
# Encapsulates data for View EngineType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EngineTypeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EngineType index.")

def get(request, engineTypeId ):
	delegate = EngineTypeDelegate()
	responseData = delegate.get( engineTypeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	engineType = json.loads(request.body)
	delegate = EngineTypeDelegate()
	responseData = delegate.createFromJson( engineType )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	engineType = json.loads(request.body)
	delegate = EngineTypeDelegate()
	responseData = delegate.save( engineType )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, engineTypeId ):
	delegate = EngineTypeDelegate()
	responseData = delegate.delete( engineTypeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EngineTypeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, engineTypeId, SupplierId ):
	delegate = EngineTypeDelegate()
	responseData = delegate.saveSupplier( engineTypeId, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, engineTypeId ):
	delegate = EngineTypeDelegate()
	responseData = delegate.deleteSupplier( engineTypeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompatibleModels( request, engineTypeId, CompatibleModelsIds ):
	delegate = EngineTypeDelegate()
	responseData = delegate.addCompatibleModels( engineTypeId, CompatibleModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompatibleModels( request, engineTypeId, CompatibleModelsIds ):
	delegate = EngineTypeDelegate()
	responseData = delegate.removeCompatibleModels( engineTypeId, CompatibleModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

