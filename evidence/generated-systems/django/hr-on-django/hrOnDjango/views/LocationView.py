import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.LocationDelegate import LocationDelegate

 #======================================================================
# 
# Encapsulates data for View Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Location index.")

def get(request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.get( locationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	location = json.loads(request.body)
	delegate = LocationDelegate()
	responseData = delegate.createFromJson( location )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	location = json.loads(request.body)
	delegate = LocationDelegate()
	responseData = delegate.save( location )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.delete( locationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LocationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, locationId, OrganizationId ):
	delegate = LocationDelegate()
	responseData = delegate.saveOrganization( locationId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.deleteOrganization( locationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDepartments( request, locationId, DepartmentsIds ):
	delegate = LocationDelegate()
	responseData = delegate.addDepartments( locationId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDepartments( request, locationId, DepartmentsIds ):
	delegate = LocationDelegate()
	responseData = delegate.removeDepartments( locationId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPositions( request, locationId, PositionsIds ):
	delegate = LocationDelegate()
	responseData = delegate.addPositions( locationId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePositions( request, locationId, PositionsIds ):
	delegate = LocationDelegate()
	responseData = delegate.removePositions( locationId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmployees( request, locationId, EmployeesIds ):
	delegate = LocationDelegate()
	responseData = delegate.addEmployees( locationId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmployees( request, locationId, EmployeesIds ):
	delegate = LocationDelegate()
	responseData = delegate.removeEmployees( locationId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

