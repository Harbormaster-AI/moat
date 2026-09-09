import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.FeeScheduleDelegate import FeeScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View FeeSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FeeSchedule index.")

def get(request, feeScheduleId ):
	delegate = FeeScheduleDelegate()
	responseData = delegate.get( feeScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	feeSchedule = json.loads(request.body)
	delegate = FeeScheduleDelegate()
	responseData = delegate.createFromJson( feeSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	feeSchedule = json.loads(request.body)
	delegate = FeeScheduleDelegate()
	responseData = delegate.save( feeSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, feeScheduleId ):
	delegate = FeeScheduleDelegate()
	responseData = delegate.delete( feeScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FeeScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPricingPlan( request, feeScheduleId, PricingPlanId ):
	delegate = FeeScheduleDelegate()
	responseData = delegate.savePricingPlan( feeScheduleId, PricingPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPricingPlan( request, feeScheduleId ):
	delegate = FeeScheduleDelegate()
	responseData = delegate.deletePricingPlan( feeScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

