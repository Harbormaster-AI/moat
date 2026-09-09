import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.LeavePolicyDelegate import LeavePolicyDelegate

 #======================================================================
# 
# Encapsulates data for View LeavePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeavePolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LeavePolicy index.")

def get(request, leavePolicyId ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.get( leavePolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	leavePolicy = json.loads(request.body)
	delegate = LeavePolicyDelegate()
	responseData = delegate.createFromJson( leavePolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	leavePolicy = json.loads(request.body)
	delegate = LeavePolicyDelegate()
	responseData = delegate.save( leavePolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, leavePolicyId ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.delete( leavePolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LeavePolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, leavePolicyId, OrganizationId ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.saveOrganization( leavePolicyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, leavePolicyId ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.deleteOrganization( leavePolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLeaveRequests( request, leavePolicyId, LeaveRequestsIds ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.addLeaveRequests( leavePolicyId, LeaveRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLeaveRequests( request, leavePolicyId, LeaveRequestsIds ):
	delegate = LeavePolicyDelegate()
	responseData = delegate.removeLeaveRequests( leavePolicyId, LeaveRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

