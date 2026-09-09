import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

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

def assignOrganization( request, departmentId, OrganizationId ):
	delegate = DepartmentDelegate()
	responseData = delegate.saveOrganization( departmentId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.deleteOrganization( departmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManager( request, departmentId, ManagerId ):
	delegate = DepartmentDelegate()
	responseData = delegate.saveManager( departmentId, ManagerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManager( request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.deleteManager( departmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCostCenter( request, departmentId, CostCenterId ):
	delegate = DepartmentDelegate()
	responseData = delegate.saveCostCenter( departmentId, CostCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCostCenter( request, departmentId ):
	delegate = DepartmentDelegate()
	responseData = delegate.deleteCostCenter( departmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPositions( request, departmentId, PositionsIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.addPositions( departmentId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePositions( request, departmentId, PositionsIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.removePositions( departmentId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmployees( request, departmentId, EmployeesIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.addEmployees( departmentId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmployees( request, departmentId, EmployeesIds ):
	delegate = DepartmentDelegate()
	responseData = delegate.removeEmployees( departmentId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

