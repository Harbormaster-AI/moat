import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.CompliancePolicyDelegate import CompliancePolicyDelegate

 #======================================================================
# 
# Encapsulates data for View CompliancePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompliancePolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CompliancePolicy index.")

def get(request, compliancePolicyId ):
	delegate = CompliancePolicyDelegate()
	responseData = delegate.get( compliancePolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	compliancePolicy = json.loads(request.body)
	delegate = CompliancePolicyDelegate()
	responseData = delegate.createFromJson( compliancePolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	compliancePolicy = json.loads(request.body)
	delegate = CompliancePolicyDelegate()
	responseData = delegate.save( compliancePolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, compliancePolicyId ):
	delegate = CompliancePolicyDelegate()
	responseData = delegate.delete( compliancePolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CompliancePolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstitution( request, compliancePolicyId, InstitutionId ):
	delegate = CompliancePolicyDelegate()
	responseData = delegate.saveInstitution( compliancePolicyId, InstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstitution( request, compliancePolicyId ):
	delegate = CompliancePolicyDelegate()
	responseData = delegate.deleteInstitution( compliancePolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

