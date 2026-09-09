import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

 #======================================================================
# 
# Encapsulates data for View MedicationDispense
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationDispenseView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MedicationDispense index.")

def get(request, medicationDispenseId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.get( medicationDispenseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	medicationDispense = json.loads(request.body)
	delegate = MedicationDispenseDelegate()
	responseData = delegate.createFromJson( medicationDispense )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	medicationDispense = json.loads(request.body)
	delegate = MedicationDispenseDelegate()
	responseData = delegate.save( medicationDispense )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, medicationDispenseId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.delete( medicationDispenseId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMedicationOrder( request, medicationDispenseId, MedicationOrderId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.saveMedicationOrder( medicationDispenseId, MedicationOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMedicationOrder( request, medicationDispenseId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.deleteMedicationOrder( medicationDispenseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPharmacy( request, medicationDispenseId, PharmacyId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.savePharmacy( medicationDispenseId, PharmacyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPharmacy( request, medicationDispenseId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.deletePharmacy( medicationDispenseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, medicationDispenseId, PatientId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.savePatient( medicationDispenseId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, medicationDispenseId ):
	delegate = MedicationDispenseDelegate()
	responseData = delegate.deletePatient( medicationDispenseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

