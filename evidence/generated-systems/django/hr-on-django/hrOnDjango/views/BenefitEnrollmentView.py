import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

 #======================================================================
# 
# Encapsulates data for View BenefitEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitEnrollmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BenefitEnrollment index.")

def get(request, benefitEnrollmentId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.get( benefitEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	benefitEnrollment = json.loads(request.body)
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.createFromJson( benefitEnrollment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	benefitEnrollment = json.loads(request.body)
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.save( benefitEnrollment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, benefitEnrollmentId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.delete( benefitEnrollmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBenefitPlan( request, benefitEnrollmentId, BenefitPlanId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.saveBenefitPlan( benefitEnrollmentId, BenefitPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBenefitPlan( request, benefitEnrollmentId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.deleteBenefitPlan( benefitEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, benefitEnrollmentId, EmployeeId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.saveEmployee( benefitEnrollmentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, benefitEnrollmentId ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.deleteEmployee( benefitEnrollmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDependents( request, benefitEnrollmentId, DependentsIds ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.addDependents( benefitEnrollmentId, DependentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDependents( request, benefitEnrollmentId, DependentsIds ):
	delegate = BenefitEnrollmentDelegate()
	responseData = delegate.removeDependents( benefitEnrollmentId, DependentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

