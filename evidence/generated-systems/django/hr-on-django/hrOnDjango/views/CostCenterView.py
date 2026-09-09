import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

 #======================================================================
# 
# Encapsulates data for View CostCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CostCenterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CostCenter index.")

def get(request, costCenterId ):
	delegate = CostCenterDelegate()
	responseData = delegate.get( costCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	costCenter = json.loads(request.body)
	delegate = CostCenterDelegate()
	responseData = delegate.createFromJson( costCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	costCenter = json.loads(request.body)
	delegate = CostCenterDelegate()
	responseData = delegate.save( costCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, costCenterId ):
	delegate = CostCenterDelegate()
	responseData = delegate.delete( costCenterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CostCenterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, costCenterId, OrganizationId ):
	delegate = CostCenterDelegate()
	responseData = delegate.saveOrganization( costCenterId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, costCenterId ):
	delegate = CostCenterDelegate()
	responseData = delegate.deleteOrganization( costCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDepartments( request, costCenterId, DepartmentsIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.addDepartments( costCenterId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDepartments( request, costCenterId, DepartmentsIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.removeDepartments( costCenterId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPositions( request, costCenterId, PositionsIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.addPositions( costCenterId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePositions( request, costCenterId, PositionsIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.removePositions( costCenterId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmployees( request, costCenterId, EmployeesIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.addEmployees( costCenterId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmployees( request, costCenterId, EmployeesIds ):
	delegate = CostCenterDelegate()
	responseData = delegate.removeEmployees( costCenterId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

