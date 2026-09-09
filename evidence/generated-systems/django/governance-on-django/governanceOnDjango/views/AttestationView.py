import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

 #======================================================================
# 
# Encapsulates data for View Attestation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AttestationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Attestation index.")

def get(request, attestationId ):
	delegate = AttestationDelegate()
	responseData = delegate.get( attestationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	attestation = json.loads(request.body)
	delegate = AttestationDelegate()
	responseData = delegate.createFromJson( attestation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	attestation = json.loads(request.body)
	delegate = AttestationDelegate()
	responseData = delegate.save( attestation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, attestationId ):
	delegate = AttestationDelegate()
	responseData = delegate.delete( attestationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AttestationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControl( request, attestationId, ControlId ):
	delegate = AttestationDelegate()
	responseData = delegate.saveControl( attestationId, ControlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControl( request, attestationId ):
	delegate = AttestationDelegate()
	responseData = delegate.deleteControl( attestationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, attestationId, PolicyId ):
	delegate = AttestationDelegate()
	responseData = delegate.savePolicy( attestationId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, attestationId ):
	delegate = AttestationDelegate()
	responseData = delegate.deletePolicy( attestationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignComplianceProgram( request, attestationId, ComplianceProgramId ):
	delegate = AttestationDelegate()
	responseData = delegate.saveComplianceProgram( attestationId, ComplianceProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignComplianceProgram( request, attestationId ):
	delegate = AttestationDelegate()
	responseData = delegate.deleteComplianceProgram( attestationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

