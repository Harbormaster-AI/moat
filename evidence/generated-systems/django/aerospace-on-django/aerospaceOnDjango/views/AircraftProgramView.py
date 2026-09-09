import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftProgramView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftProgram index.")

def get(request, aircraftProgramId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.get( aircraftProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftProgram = json.loads(request.body)
	delegate = AircraftProgramDelegate()
	responseData = delegate.createFromJson( aircraftProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftProgram = json.loads(request.body)
	delegate = AircraftProgramDelegate()
	responseData = delegate.save( aircraftProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftProgramId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.delete( aircraftProgramId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftProgramDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManufacturer( request, aircraftProgramId, ManufacturerId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.saveManufacturer( aircraftProgramId, ManufacturerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManufacturer( request, aircraftProgramId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.deleteManufacturer( aircraftProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTypeCertificate( request, aircraftProgramId, TypeCertificateId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.saveTypeCertificate( aircraftProgramId, TypeCertificateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTypeCertificate( request, aircraftProgramId ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.deleteTypeCertificate( aircraftProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAircraftFamilies( request, aircraftProgramId, AircraftFamiliesIds ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.addAircraftFamilies( aircraftProgramId, AircraftFamiliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAircraftFamilies( request, aircraftProgramId, AircraftFamiliesIds ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.removeAircraftFamilies( aircraftProgramId, AircraftFamiliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addKeySuppliers( request, aircraftProgramId, KeySuppliersIds ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.addKeySuppliers( aircraftProgramId, KeySuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeKeySuppliers( request, aircraftProgramId, KeySuppliersIds ):
	delegate = AircraftProgramDelegate()
	responseData = delegate.removeKeySuppliers( aircraftProgramId, KeySuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

