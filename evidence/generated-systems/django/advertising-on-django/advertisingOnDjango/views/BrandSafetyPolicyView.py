import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.BrandSafetyPolicyDelegate import BrandSafetyPolicyDelegate

 #======================================================================
# 
# Encapsulates data for View BrandSafetyPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandSafetyPolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BrandSafetyPolicy index.")

def get(request, brandSafetyPolicyId ):
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.get( brandSafetyPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	brandSafetyPolicy = json.loads(request.body)
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.createFromJson( brandSafetyPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	brandSafetyPolicy = json.loads(request.body)
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.save( brandSafetyPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, brandSafetyPolicyId ):
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.delete( brandSafetyPolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTargetingProfiles( request, brandSafetyPolicyId, TargetingProfilesIds ):
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.addTargetingProfiles( brandSafetyPolicyId, TargetingProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTargetingProfiles( request, brandSafetyPolicyId, TargetingProfilesIds ):
	delegate = BrandSafetyPolicyDelegate()
	responseData = delegate.removeTargetingProfiles( brandSafetyPolicyId, TargetingProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

