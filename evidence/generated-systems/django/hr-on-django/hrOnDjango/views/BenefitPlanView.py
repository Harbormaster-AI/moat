import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.BenefitPlanDelegate import BenefitPlanDelegate

 #======================================================================
# 
# Encapsulates data for View BenefitPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitPlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BenefitPlan index.")

def get(request, benefitPlanId ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.get( benefitPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	benefitPlan = json.loads(request.body)
	delegate = BenefitPlanDelegate()
	responseData = delegate.createFromJson( benefitPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	benefitPlan = json.loads(request.body)
	delegate = BenefitPlanDelegate()
	responseData = delegate.save( benefitPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, benefitPlanId ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.delete( benefitPlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BenefitPlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, benefitPlanId, OrganizationId ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.saveOrganization( benefitPlanId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, benefitPlanId ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.deleteOrganization( benefitPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEnrollments( request, benefitPlanId, EnrollmentsIds ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.addEnrollments( benefitPlanId, EnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEnrollments( request, benefitPlanId, EnrollmentsIds ):
	delegate = BenefitPlanDelegate()
	responseData = delegate.removeEnrollments( benefitPlanId, EnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

