import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

 #======================================================================
# 
# Encapsulates data for View Pharmacy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PharmacyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Pharmacy index.")

def get(request, pharmacyId ):
	delegate = PharmacyDelegate()
	responseData = delegate.get( pharmacyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	pharmacy = json.loads(request.body)
	delegate = PharmacyDelegate()
	responseData = delegate.createFromJson( pharmacy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	pharmacy = json.loads(request.body)
	delegate = PharmacyDelegate()
	responseData = delegate.save( pharmacy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, pharmacyId ):
	delegate = PharmacyDelegate()
	responseData = delegate.delete( pharmacyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PharmacyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, pharmacyId, FacilityId ):
	delegate = PharmacyDelegate()
	responseData = delegate.saveFacility( pharmacyId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, pharmacyId ):
	delegate = PharmacyDelegate()
	responseData = delegate.deleteFacility( pharmacyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMedicationDispenses( request, pharmacyId, MedicationDispensesIds ):
	delegate = PharmacyDelegate()
	responseData = delegate.addMedicationDispenses( pharmacyId, MedicationDispensesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMedicationDispenses( request, pharmacyId, MedicationDispensesIds ):
	delegate = PharmacyDelegate()
	responseData = delegate.removeMedicationDispenses( pharmacyId, MedicationDispensesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMedicationOrders( request, pharmacyId, MedicationOrdersIds ):
	delegate = PharmacyDelegate()
	responseData = delegate.addMedicationOrders( pharmacyId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMedicationOrders( request, pharmacyId, MedicationOrdersIds ):
	delegate = PharmacyDelegate()
	responseData = delegate.removeMedicationOrders( pharmacyId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

