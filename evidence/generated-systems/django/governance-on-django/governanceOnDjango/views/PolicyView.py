import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

 #======================================================================
# 
# Encapsulates data for View Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Policy index.")

def get(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.get( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.createFromJson( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.save( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.delete( policyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, policyId, OrganizationId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveOrganization( policyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteOrganization( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwners( request, policyId, OwnersIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addOwners( policyId, OwnersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwners( request, policyId, OwnersIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeOwners( policyId, OwnersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedRequirements( request, policyId, RelatedRequirementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addRelatedRequirements( policyId, RelatedRequirementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedRequirements( request, policyId, RelatedRequirementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeRelatedRequirements( policyId, RelatedRequirementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, policyId, ControlsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addControls( policyId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, policyId, ControlsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeControls( policyId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcedures( request, policyId, ProceduresIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addProcedures( policyId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcedures( request, policyId, ProceduresIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeProcedures( policyId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExceptions( request, policyId, ExceptionsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addExceptions( policyId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExceptions( request, policyId, ExceptionsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeExceptions( policyId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAttestations( request, policyId, AttestationsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addAttestations( policyId, AttestationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAttestations( request, policyId, AttestationsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeAttestations( policyId, AttestationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

