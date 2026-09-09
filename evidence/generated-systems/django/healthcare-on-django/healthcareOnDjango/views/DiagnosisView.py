import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.DiagnosisDelegate import DiagnosisDelegate

 #======================================================================
# 
# Encapsulates data for View Diagnosis
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DiagnosisView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Diagnosis index.")

def get(request, diagnosisId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.get( diagnosisId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	diagnosis = json.loads(request.body)
	delegate = DiagnosisDelegate()
	responseData = delegate.createFromJson( diagnosis )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	diagnosis = json.loads(request.body)
	delegate = DiagnosisDelegate()
	responseData = delegate.save( diagnosis )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, diagnosisId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.delete( diagnosisId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DiagnosisDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, diagnosisId, EncounterId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.saveEncounter( diagnosisId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, diagnosisId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.deleteEncounter( diagnosisId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, diagnosisId, PatientId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.savePatient( diagnosisId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, diagnosisId ):
	delegate = DiagnosisDelegate()
	responseData = delegate.deletePatient( diagnosisId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

