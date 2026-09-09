import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

 #======================================================================
# 
# Encapsulates data for View AudienceSegment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AudienceSegmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AudienceSegment index.")

def get(request, audienceSegmentId ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.get( audienceSegmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	audienceSegment = json.loads(request.body)
	delegate = AudienceSegmentDelegate()
	responseData = delegate.createFromJson( audienceSegment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	audienceSegment = json.loads(request.body)
	delegate = AudienceSegmentDelegate()
	responseData = delegate.save( audienceSegment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, audienceSegmentId ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.delete( audienceSegmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProvider( request, audienceSegmentId, ProviderId ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.saveProvider( audienceSegmentId, ProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProvider( request, audienceSegmentId ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.deleteProvider( audienceSegmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, audienceSegmentId, CampaignsIds ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.addCampaigns( audienceSegmentId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, audienceSegmentId, CampaignsIds ):
	delegate = AudienceSegmentDelegate()
	responseData = delegate.removeCampaigns( audienceSegmentId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

