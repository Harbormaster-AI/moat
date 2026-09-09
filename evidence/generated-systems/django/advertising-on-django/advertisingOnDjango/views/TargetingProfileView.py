import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

 #======================================================================
# 
# Encapsulates data for View TargetingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TargetingProfileView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TargetingProfile index.")

def get(request, targetingProfileId ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.get( targetingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	targetingProfile = json.loads(request.body)
	delegate = TargetingProfileDelegate()
	responseData = delegate.createFromJson( targetingProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	targetingProfile = json.loads(request.body)
	delegate = TargetingProfileDelegate()
	responseData = delegate.save( targetingProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, targetingProfileId ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.delete( targetingProfileId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TargetingProfileDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBrandSafetyPolicy( request, targetingProfileId, BrandSafetyPolicyId ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.saveBrandSafetyPolicy( targetingProfileId, BrandSafetyPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBrandSafetyPolicy( request, targetingProfileId ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.deleteBrandSafetyPolicy( targetingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAudienceSegments( request, targetingProfileId, AudienceSegmentsIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.addAudienceSegments( targetingProfileId, AudienceSegmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAudienceSegments( request, targetingProfileId, AudienceSegmentsIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.removeAudienceSegments( targetingProfileId, AudienceSegmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGeoRegions( request, targetingProfileId, GeoRegionsIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.addGeoRegions( targetingProfileId, GeoRegionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGeoRegions( request, targetingProfileId, GeoRegionsIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.removeGeoRegions( targetingProfileId, GeoRegionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContentCategories( request, targetingProfileId, ContentCategoriesIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.addContentCategories( targetingProfileId, ContentCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContentCategories( request, targetingProfileId, ContentCategoriesIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.removeContentCategories( targetingProfileId, ContentCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDeviceCriteria( request, targetingProfileId, DeviceCriteriaIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.addDeviceCriteria( targetingProfileId, DeviceCriteriaIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDeviceCriteria( request, targetingProfileId, DeviceCriteriaIds ):
	delegate = TargetingProfileDelegate()
	responseData = delegate.removeDeviceCriteria( targetingProfileId, DeviceCriteriaIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

