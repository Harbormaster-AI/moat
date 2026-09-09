import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

 #======================================================================
# 
# Encapsulates data for View Approval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApprovalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Approval index.")

def get(request, approvalId ):
	delegate = ApprovalDelegate()
	responseData = delegate.get( approvalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	approval = json.loads(request.body)
	delegate = ApprovalDelegate()
	responseData = delegate.createFromJson( approval )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	approval = json.loads(request.body)
	delegate = ApprovalDelegate()
	responseData = delegate.save( approval )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, approvalId ):
	delegate = ApprovalDelegate()
	responseData = delegate.delete( approvalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ApprovalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApprover( request, approvalId, ApproverId ):
	delegate = ApprovalDelegate()
	responseData = delegate.saveApprover( approvalId, ApproverId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApprover( request, approvalId ):
	delegate = ApprovalDelegate()
	responseData = delegate.deleteApprover( approvalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTimesheet( request, approvalId, TimesheetId ):
	delegate = ApprovalDelegate()
	responseData = delegate.saveTimesheet( approvalId, TimesheetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTimesheet( request, approvalId ):
	delegate = ApprovalDelegate()
	responseData = delegate.deleteTimesheet( approvalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLeaveRequest( request, approvalId, LeaveRequestId ):
	delegate = ApprovalDelegate()
	responseData = delegate.saveLeaveRequest( approvalId, LeaveRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLeaveRequest( request, approvalId ):
	delegate = ApprovalDelegate()
	responseData = delegate.deleteLeaveRequest( approvalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

