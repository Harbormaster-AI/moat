import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

 #======================================================================
# 
# Encapsulates data for View ComplianceProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceProgramView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ComplianceProgram index.")

def get(request, complianceProgramId ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.get( complianceProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	complianceProgram = json.loads(request.body)
	delegate = ComplianceProgramDelegate()
	responseData = delegate.createFromJson( complianceProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	complianceProgram = json.loads(request.body)
	delegate = ComplianceProgramDelegate()
	responseData = delegate.save( complianceProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, complianceProgramId ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.delete( complianceProgramId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, complianceProgramId, OrganizationId ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.saveOrganization( complianceProgramId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, complianceProgramId ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.deleteOrganization( complianceProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRequirements( request, complianceProgramId, RequirementsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.addRequirements( complianceProgramId, RequirementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRequirements( request, complianceProgramId, RequirementsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.removeRequirements( complianceProgramId, RequirementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, complianceProgramId, ControlsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.addControls( complianceProgramId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, complianceProgramId, ControlsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.removeControls( complianceProgramId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAttestations( request, complianceProgramId, AttestationsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.addAttestations( complianceProgramId, AttestationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAttestations( request, complianceProgramId, AttestationsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.removeAttestations( complianceProgramId, AttestationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRegulations( request, complianceProgramId, RegulationsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.addRegulations( complianceProgramId, RegulationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRegulations( request, complianceProgramId, RegulationsIds ):
	delegate = ComplianceProgramDelegate()
	responseData = delegate.removeRegulations( complianceProgramId, RegulationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

