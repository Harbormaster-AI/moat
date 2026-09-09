import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

 #======================================================================
# 
# Encapsulates data for View LeaveRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveRequestView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LeaveRequest index.")

def get(request, leaveRequestId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.get( leaveRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	leaveRequest = json.loads(request.body)
	delegate = LeaveRequestDelegate()
	responseData = delegate.createFromJson( leaveRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	leaveRequest = json.loads(request.body)
	delegate = LeaveRequestDelegate()
	responseData = delegate.save( leaveRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, leaveRequestId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.delete( leaveRequestId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LeaveRequestDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, leaveRequestId, EmployeeId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.saveEmployee( leaveRequestId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, leaveRequestId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.deleteEmployee( leaveRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLeavePolicy( request, leaveRequestId, LeavePolicyId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.saveLeavePolicy( leaveRequestId, LeavePolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLeavePolicy( request, leaveRequestId ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.deleteLeavePolicy( leaveRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApprovals( request, leaveRequestId, ApprovalsIds ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.addApprovals( leaveRequestId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApprovals( request, leaveRequestId, ApprovalsIds ):
	delegate = LeaveRequestDelegate()
	responseData = delegate.removeApprovals( leaveRequestId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

