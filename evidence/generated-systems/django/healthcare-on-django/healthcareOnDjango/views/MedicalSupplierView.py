import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.MedicalSupplierDelegate import MedicalSupplierDelegate

 #======================================================================
# 
# Encapsulates data for View MedicalSupplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalSupplierView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MedicalSupplier index.")

def get(request, medicalSupplierId ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.get( medicalSupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	medicalSupplier = json.loads(request.body)
	delegate = MedicalSupplierDelegate()
	responseData = delegate.createFromJson( medicalSupplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	medicalSupplier = json.loads(request.body)
	delegate = MedicalSupplierDelegate()
	responseData = delegate.save( medicalSupplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, medicalSupplierId ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.delete( medicalSupplierId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFacilities( request, medicalSupplierId, FacilitiesIds ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.addFacilities( medicalSupplierId, FacilitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFacilities( request, medicalSupplierId, FacilitiesIds ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.removeFacilities( medicalSupplierId, FacilitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, medicalSupplierId, InventoryItemsIds ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.addInventoryItems( medicalSupplierId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, medicalSupplierId, InventoryItemsIds ):
	delegate = MedicalSupplierDelegate()
	responseData = delegate.removeInventoryItems( medicalSupplierId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

