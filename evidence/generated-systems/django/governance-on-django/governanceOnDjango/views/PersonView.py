import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.PersonDelegate import PersonDelegate

 #======================================================================
# 
# Encapsulates data for View Person
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PersonView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Person index.")

def get(request, personId ):
	delegate = PersonDelegate()
	responseData = delegate.get( personId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	person = json.loads(request.body)
	delegate = PersonDelegate()
	responseData = delegate.createFromJson( person )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	person = json.loads(request.body)
	delegate = PersonDelegate()
	responseData = delegate.save( person )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, personId ):
	delegate = PersonDelegate()
	responseData = delegate.delete( personId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PersonDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRoleAssignments( request, personId, RoleAssignmentsIds ):
	delegate = PersonDelegate()
	responseData = delegate.addRoleAssignments( personId, RoleAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRoleAssignments( request, personId, RoleAssignmentsIds ):
	delegate = PersonDelegate()
	responseData = delegate.removeRoleAssignments( personId, RoleAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwnedPolicies( request, personId, OwnedPoliciesIds ):
	delegate = PersonDelegate()
	responseData = delegate.addOwnedPolicies( personId, OwnedPoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwnedPolicies( request, personId, OwnedPoliciesIds ):
	delegate = PersonDelegate()
	responseData = delegate.removeOwnedPolicies( personId, OwnedPoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCorrectiveActions( request, personId, CorrectiveActionsIds ):
	delegate = PersonDelegate()
	responseData = delegate.addCorrectiveActions( personId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCorrectiveActions( request, personId, CorrectiveActionsIds ):
	delegate = PersonDelegate()
	responseData = delegate.removeCorrectiveActions( personId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

