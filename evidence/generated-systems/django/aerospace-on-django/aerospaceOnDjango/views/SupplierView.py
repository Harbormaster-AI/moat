import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

 #======================================================================
# 
# Encapsulates data for View Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Supplier index.")

def get(request, supplierId ):
	delegate = SupplierDelegate()
	responseData = delegate.get( supplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	supplier = json.loads(request.body)
	delegate = SupplierDelegate()
	responseData = delegate.createFromJson( supplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	supplier = json.loads(request.body)
	delegate = SupplierDelegate()
	responseData = delegate.save( supplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, supplierId ):
	delegate = SupplierDelegate()
	responseData = delegate.delete( supplierId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SupplierDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addManufacturers( request, supplierId, ManufacturersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addManufacturers( supplierId, ManufacturersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeManufacturers( request, supplierId, ManufacturersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeManufacturers( supplierId, ManufacturersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addComponents( request, supplierId, ComponentsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addComponents( supplierId, ComponentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeComponents( request, supplierId, ComponentsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeComponents( supplierId, ComponentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEngineTypes( request, supplierId, EngineTypesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addEngineTypes( supplierId, EngineTypesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEngineTypes( request, supplierId, EngineTypesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeEngineTypes( supplierId, EngineTypesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAvionicsSuites( request, supplierId, AvionicsSuitesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addAvionicsSuites( supplierId, AvionicsSuitesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAvionicsSuites( request, supplierId, AvionicsSuitesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeAvionicsSuites( supplierId, AvionicsSuitesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApus( request, supplierId, ApusIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addApus( supplierId, ApusIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApus( request, supplierId, ApusIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeApus( supplierId, ApusIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLandingGears( request, supplierId, LandingGearsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addLandingGears( supplierId, LandingGearsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLandingGears( request, supplierId, LandingGearsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeLandingGears( supplierId, LandingGearsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

