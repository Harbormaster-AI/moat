import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

 #======================================================================
# 
# Encapsulates data for View AerospaceManufacturer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AerospaceManufacturerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AerospaceManufacturer index.")

def get(request, aerospaceManufacturerId ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.get( aerospaceManufacturerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aerospaceManufacturer = json.loads(request.body)
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.createFromJson( aerospaceManufacturer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aerospaceManufacturer = json.loads(request.body)
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.save( aerospaceManufacturer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aerospaceManufacturerId ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.delete( aerospaceManufacturerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPrograms( request, aerospaceManufacturerId, ProgramsIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.addPrograms( aerospaceManufacturerId, ProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePrograms( request, aerospaceManufacturerId, ProgramsIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.removePrograms( aerospaceManufacturerId, ProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlants( request, aerospaceManufacturerId, PlantsIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.addPlants( aerospaceManufacturerId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlants( request, aerospaceManufacturerId, PlantsIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.removePlants( aerospaceManufacturerId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSuppliers( request, aerospaceManufacturerId, SuppliersIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.addSuppliers( aerospaceManufacturerId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSuppliers( request, aerospaceManufacturerId, SuppliersIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.removeSuppliers( aerospaceManufacturerId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProductionCertificates( request, aerospaceManufacturerId, ProductionCertificatesIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.addProductionCertificates( aerospaceManufacturerId, ProductionCertificatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProductionCertificates( request, aerospaceManufacturerId, ProductionCertificatesIds ):
	delegate = AerospaceManufacturerDelegate()
	responseData = delegate.removeProductionCertificates( aerospaceManufacturerId, ProductionCertificatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

