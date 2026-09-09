import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.SalaryComponentDelegate import SalaryComponentDelegate

 #======================================================================
# 
# Encapsulates data for View SalaryComponent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalaryComponentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SalaryComponent index.")

def get(request, salaryComponentId ):
	delegate = SalaryComponentDelegate()
	responseData = delegate.get( salaryComponentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	salaryComponent = json.loads(request.body)
	delegate = SalaryComponentDelegate()
	responseData = delegate.createFromJson( salaryComponent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	salaryComponent = json.loads(request.body)
	delegate = SalaryComponentDelegate()
	responseData = delegate.save( salaryComponent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, salaryComponentId ):
	delegate = SalaryComponentDelegate()
	responseData = delegate.delete( salaryComponentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SalaryComponentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCompensationPackage( request, salaryComponentId, CompensationPackageId ):
	delegate = SalaryComponentDelegate()
	responseData = delegate.saveCompensationPackage( salaryComponentId, CompensationPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCompensationPackage( request, salaryComponentId ):
	delegate = SalaryComponentDelegate()
	responseData = delegate.deleteCompensationPackage( salaryComponentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

