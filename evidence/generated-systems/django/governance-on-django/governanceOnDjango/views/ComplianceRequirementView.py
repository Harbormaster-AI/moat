import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

 #======================================================================
# 
# Encapsulates data for View ComplianceRequirement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceRequirementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ComplianceRequirement index.")

def get(request, complianceRequirementId ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.get( complianceRequirementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	complianceRequirement = json.loads(request.body)
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.createFromJson( complianceRequirement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	complianceRequirement = json.loads(request.body)
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.save( complianceRequirement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, complianceRequirementId ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.delete( complianceRequirementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignComplianceProgram( request, complianceRequirementId, ComplianceProgramId ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.saveComplianceProgram( complianceRequirementId, ComplianceProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignComplianceProgram( request, complianceRequirementId ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.deleteComplianceProgram( complianceRequirementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, complianceRequirementId, PoliciesIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.addPolicies( complianceRequirementId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, complianceRequirementId, PoliciesIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.removePolicies( complianceRequirementId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, complianceRequirementId, ControlsIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.addControls( complianceRequirementId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, complianceRequirementId, ControlsIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.removeControls( complianceRequirementId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObligations( request, complianceRequirementId, ObligationsIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.addObligations( complianceRequirementId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObligations( request, complianceRequirementId, ObligationsIds ):
	delegate = ComplianceRequirementDelegate()
	responseData = delegate.removeObligations( complianceRequirementId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

