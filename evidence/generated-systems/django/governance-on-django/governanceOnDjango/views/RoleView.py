import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RoleDelegate import RoleDelegate

 #======================================================================
# 
# Encapsulates data for View Role
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Role index.")

def get(request, roleId ):
	delegate = RoleDelegate()
	responseData = delegate.get( roleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	role = json.loads(request.body)
	delegate = RoleDelegate()
	responseData = delegate.createFromJson( role )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	role = json.loads(request.body)
	delegate = RoleDelegate()
	responseData = delegate.save( role )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, roleId ):
	delegate = RoleDelegate()
	responseData = delegate.delete( roleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RoleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssignments( request, roleId, AssignmentsIds ):
	delegate = RoleDelegate()
	responseData = delegate.addAssignments( roleId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssignments( request, roleId, AssignmentsIds ):
	delegate = RoleDelegate()
	responseData = delegate.removeAssignments( roleId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

