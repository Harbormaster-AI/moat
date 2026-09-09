import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

 #======================================================================
# 
# Encapsulates data for View CareTeam
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTeamView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CareTeam index.")

def get(request, careTeamId ):
	delegate = CareTeamDelegate()
	responseData = delegate.get( careTeamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	careTeam = json.loads(request.body)
	delegate = CareTeamDelegate()
	responseData = delegate.createFromJson( careTeam )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	careTeam = json.loads(request.body)
	delegate = CareTeamDelegate()
	responseData = delegate.save( careTeam )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, careTeamId ):
	delegate = CareTeamDelegate()
	responseData = delegate.delete( careTeamId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CareTeamDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDepartment( request, careTeamId, DepartmentId ):
	delegate = CareTeamDelegate()
	responseData = delegate.saveDepartment( careTeamId, DepartmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDepartment( request, careTeamId ):
	delegate = CareTeamDelegate()
	responseData = delegate.deleteDepartment( careTeamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClinicians( request, careTeamId, CliniciansIds ):
	delegate = CareTeamDelegate()
	responseData = delegate.addClinicians( careTeamId, CliniciansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClinicians( request, careTeamId, CliniciansIds ):
	delegate = CareTeamDelegate()
	responseData = delegate.removeClinicians( careTeamId, CliniciansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPatients( request, careTeamId, PatientsIds ):
	delegate = CareTeamDelegate()
	responseData = delegate.addPatients( careTeamId, PatientsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePatients( request, careTeamId, PatientsIds ):
	delegate = CareTeamDelegate()
	responseData = delegate.removePatients( careTeamId, PatientsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

