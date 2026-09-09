import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.HealthSystemDelegate import HealthSystemDelegate

 #======================================================================
# 
# Encapsulates data for View HealthSystem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class HealthSystemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the HealthSystem index.")

def get(request, healthSystemId ):
	delegate = HealthSystemDelegate()
	responseData = delegate.get( healthSystemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	healthSystem = json.loads(request.body)
	delegate = HealthSystemDelegate()
	responseData = delegate.createFromJson( healthSystem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	healthSystem = json.loads(request.body)
	delegate = HealthSystemDelegate()
	responseData = delegate.save( healthSystem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, healthSystemId ):
	delegate = HealthSystemDelegate()
	responseData = delegate.delete( healthSystemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = HealthSystemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFacilities( request, healthSystemId, FacilitiesIds ):
	delegate = HealthSystemDelegate()
	responseData = delegate.addFacilities( healthSystemId, FacilitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFacilities( request, healthSystemId, FacilitiesIds ):
	delegate = HealthSystemDelegate()
	responseData = delegate.removeFacilities( healthSystemId, FacilitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSuppliers( request, healthSystemId, SuppliersIds ):
	delegate = HealthSystemDelegate()
	responseData = delegate.addSuppliers( healthSystemId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSuppliers( request, healthSystemId, SuppliersIds ):
	delegate = HealthSystemDelegate()
	responseData = delegate.removeSuppliers( healthSystemId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

