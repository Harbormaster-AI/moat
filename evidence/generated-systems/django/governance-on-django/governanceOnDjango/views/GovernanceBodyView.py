import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.GovernanceBodyDelegate import GovernanceBodyDelegate

 #======================================================================
# 
# Encapsulates data for View GovernanceBody
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceBodyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GovernanceBody index.")

def get(request, governanceBodyId ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.get( governanceBodyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	governanceBody = json.loads(request.body)
	delegate = GovernanceBodyDelegate()
	responseData = delegate.createFromJson( governanceBody )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	governanceBody = json.loads(request.body)
	delegate = GovernanceBodyDelegate()
	responseData = delegate.save( governanceBody )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, governanceBodyId ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.delete( governanceBodyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, governanceBodyId, OrganizationId ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.saveOrganization( governanceBodyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, governanceBodyId ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.deleteOrganization( governanceBodyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRoleAssignments( request, governanceBodyId, RoleAssignmentsIds ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.addRoleAssignments( governanceBodyId, RoleAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRoleAssignments( request, governanceBodyId, RoleAssignmentsIds ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.removeRoleAssignments( governanceBodyId, RoleAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, governanceBodyId, PoliciesIds ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.addPolicies( governanceBodyId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, governanceBodyId, PoliciesIds ):
	delegate = GovernanceBodyDelegate()
	responseData = delegate.removePolicies( governanceBodyId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

