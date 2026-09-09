import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

 #======================================================================
# 
# Encapsulates data for View PricingPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PricingPlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PricingPlan index.")

def get(request, pricingPlanId ):
	delegate = PricingPlanDelegate()
	responseData = delegate.get( pricingPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	pricingPlan = json.loads(request.body)
	delegate = PricingPlanDelegate()
	responseData = delegate.createFromJson( pricingPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	pricingPlan = json.loads(request.body)
	delegate = PricingPlanDelegate()
	responseData = delegate.save( pricingPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, pricingPlanId ):
	delegate = PricingPlanDelegate()
	responseData = delegate.delete( pricingPlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PricingPlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProductOffering( request, pricingPlanId, ProductOfferingId ):
	delegate = PricingPlanDelegate()
	responseData = delegate.saveProductOffering( pricingPlanId, ProductOfferingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProductOffering( request, pricingPlanId ):
	delegate = PricingPlanDelegate()
	responseData = delegate.deleteProductOffering( pricingPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFeeSchedules( request, pricingPlanId, FeeSchedulesIds ):
	delegate = PricingPlanDelegate()
	responseData = delegate.addFeeSchedules( pricingPlanId, FeeSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFeeSchedules( request, pricingPlanId, FeeSchedulesIds ):
	delegate = PricingPlanDelegate()
	responseData = delegate.removeFeeSchedules( pricingPlanId, FeeSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLimits( request, pricingPlanId, LimitsIds ):
	delegate = PricingPlanDelegate()
	responseData = delegate.addLimits( pricingPlanId, LimitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLimits( request, pricingPlanId, LimitsIds ):
	delegate = PricingPlanDelegate()
	responseData = delegate.removeLimits( pricingPlanId, LimitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

