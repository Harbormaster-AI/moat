import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.ReportDelegate import ReportDelegate

 #======================================================================
# 
# Encapsulates data for View Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Report index.")

def get(request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.get( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	report = json.loads(request.body)
	delegate = ReportDelegate()
	responseData = delegate.createFromJson( report )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	report = json.loads(request.body)
	delegate = ReportDelegate()
	responseData = delegate.save( report )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.delete( reportId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReportDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdAccount( request, reportId, AdAccountId ):
	delegate = ReportDelegate()
	responseData = delegate.saveAdAccount( reportId, AdAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdAccount( request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.deleteAdAccount( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, reportId, CampaignId ):
	delegate = ReportDelegate()
	responseData = delegate.saveCampaign( reportId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.deleteCampaign( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineItem( request, reportId, LineItemId ):
	delegate = ReportDelegate()
	responseData = delegate.saveLineItem( reportId, LineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineItem( request, reportId ):
	delegate = ReportDelegate()
	responseData = delegate.deleteLineItem( reportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

