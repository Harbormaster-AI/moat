import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.UsageLimitDelegate import UsageLimitDelegate

 #======================================================================
# 
# Encapsulates data for View UsageLimit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UsageLimitView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the UsageLimit index.")

def get(request, usageLimitId ):
	delegate = UsageLimitDelegate()
	responseData = delegate.get( usageLimitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	usageLimit = json.loads(request.body)
	delegate = UsageLimitDelegate()
	responseData = delegate.createFromJson( usageLimit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	usageLimit = json.loads(request.body)
	delegate = UsageLimitDelegate()
	responseData = delegate.save( usageLimit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, usageLimitId ):
	delegate = UsageLimitDelegate()
	responseData = delegate.delete( usageLimitId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = UsageLimitDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPricingPlan( request, usageLimitId, PricingPlanId ):
	delegate = UsageLimitDelegate()
	responseData = delegate.savePricingPlan( usageLimitId, PricingPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPricingPlan( request, usageLimitId ):
	delegate = UsageLimitDelegate()
	responseData = delegate.deletePricingPlan( usageLimitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

