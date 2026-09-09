import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

 #======================================================================
# 
# Encapsulates data for View Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Organization index.")

def get(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.get( organizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.createFromJson( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.save( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.delete( organizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrganizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDepartments( request, organizationId, DepartmentsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addDepartments( organizationId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDepartments( request, organizationId, DepartmentsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeDepartments( organizationId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLocations( request, organizationId, LocationsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addLocations( organizationId, LocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLocations( request, organizationId, LocationsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeLocations( organizationId, LocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addJobFamilies( request, organizationId, JobFamiliesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addJobFamilies( organizationId, JobFamiliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeJobFamilies( request, organizationId, JobFamiliesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeJobFamilies( organizationId, JobFamiliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBenefitPlans( request, organizationId, BenefitPlansIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addBenefitPlans( organizationId, BenefitPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBenefitPlans( request, organizationId, BenefitPlansIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeBenefitPlans( organizationId, BenefitPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCostCenters( request, organizationId, CostCentersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addCostCenters( organizationId, CostCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCostCenters( request, organizationId, CostCentersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeCostCenters( organizationId, CostCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayrollCalendars( request, organizationId, PayrollCalendarsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addPayrollCalendars( organizationId, PayrollCalendarsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayrollCalendars( request, organizationId, PayrollCalendarsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removePayrollCalendars( organizationId, PayrollCalendarsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

