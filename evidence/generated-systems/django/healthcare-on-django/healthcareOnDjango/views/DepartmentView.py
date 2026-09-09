import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

 #======================================================================
# 
# Encapsulates data for View Department
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DepartmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Department index.")

def get(request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.get( departmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	department = json.loads(request.body)
	delegate = DepartmentDelegate()
	responseData = delegate.createFromJson( department )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	department = json.loads(request.body)
	delegate = DepartmentDelegate()
	responseData = delegate.save( department )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.delete( departmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DepartmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, departmentId, FacilityId ):
	delegate = DepartmentDelegate()
	responseData = delegate.saveFacility( departmentId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.deleteFacility( departmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCareTeams( request, departmentId, CareTeamsIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.addCareTeams( departmentId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCareTeams( request, departmentId, CareTeamsIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.removeCareTeams( departmentId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

