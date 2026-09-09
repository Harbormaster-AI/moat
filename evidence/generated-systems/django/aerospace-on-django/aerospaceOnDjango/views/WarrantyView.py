import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.WarrantyDelegate import WarrantyDelegate

 #======================================================================
# 
# Encapsulates data for View Warranty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarrantyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Warranty index.")

def get(request, warrantyId ):
	delegate = WarrantyDelegate()
	responseData = delegate.get( warrantyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	warranty = json.loads(request.body)
	delegate = WarrantyDelegate()
	responseData = delegate.createFromJson( warranty )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	warranty = json.loads(request.body)
	delegate = WarrantyDelegate()
	responseData = delegate.save( warranty )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, warrantyId ):
	delegate = WarrantyDelegate()
	responseData = delegate.delete( warrantyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WarrantyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraft( request, warrantyId, AircraftId ):
	delegate = WarrantyDelegate()
	responseData = delegate.saveAircraft( warrantyId, AircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraft( request, warrantyId ):
	delegate = WarrantyDelegate()
	responseData = delegate.deleteAircraft( warrantyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

