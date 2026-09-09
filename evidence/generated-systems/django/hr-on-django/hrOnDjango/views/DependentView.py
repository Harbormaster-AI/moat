import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.DependentDelegate import DependentDelegate

 #======================================================================
# 
# Encapsulates data for View Dependent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DependentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Dependent index.")

def get(request, dependentId ):
	delegate = DependentDelegate()
	responseData = delegate.get( dependentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dependent = json.loads(request.body)
	delegate = DependentDelegate()
	responseData = delegate.createFromJson( dependent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dependent = json.loads(request.body)
	delegate = DependentDelegate()
	responseData = delegate.save( dependent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dependentId ):
	delegate = DependentDelegate()
	responseData = delegate.delete( dependentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DependentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBenefitEnrollment( request, dependentId, BenefitEnrollmentId ):
	delegate = DependentDelegate()
	responseData = delegate.saveBenefitEnrollment( dependentId, BenefitEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBenefitEnrollment( request, dependentId ):
	delegate = DependentDelegate()
	responseData = delegate.deleteBenefitEnrollment( dependentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, dependentId, EmployeeId ):
	delegate = DependentDelegate()
	responseData = delegate.saveEmployee( dependentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, dependentId ):
	delegate = DependentDelegate()
	responseData = delegate.deleteEmployee( dependentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

