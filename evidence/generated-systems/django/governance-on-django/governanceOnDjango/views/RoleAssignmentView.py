import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for View RoleAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleAssignmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RoleAssignment index.")

def get(request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.get( roleAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	roleAssignment = json.loads(request.body)
	delegate = RoleAssignmentDelegate()
	responseData = delegate.createFromJson( roleAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	roleAssignment = json.loads(request.body)
	delegate = RoleAssignmentDelegate()
	responseData = delegate.save( roleAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.delete( roleAssignmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPerson( request, roleAssignmentId, PersonId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.savePerson( roleAssignmentId, PersonId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPerson( request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.deletePerson( roleAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRole( request, roleAssignmentId, RoleId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.saveRole( roleAssignmentId, RoleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRole( request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.deleteRole( roleAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignGovernanceBody( request, roleAssignmentId, GovernanceBodyId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.saveGovernanceBody( roleAssignmentId, GovernanceBodyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignGovernanceBody( request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.deleteGovernanceBody( roleAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, roleAssignmentId, OrganizationId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.saveOrganization( roleAssignmentId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, roleAssignmentId ):
	delegate = RoleAssignmentDelegate()
	responseData = delegate.deleteOrganization( roleAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

