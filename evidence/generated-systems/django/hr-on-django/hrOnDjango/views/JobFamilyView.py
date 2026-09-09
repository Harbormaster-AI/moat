import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.JobFamilyDelegate import JobFamilyDelegate

 #======================================================================
# 
# Encapsulates data for View JobFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobFamilyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the JobFamily index.")

def get(request, jobFamilyId ):
	delegate = JobFamilyDelegate()
	responseData = delegate.get( jobFamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	jobFamily = json.loads(request.body)
	delegate = JobFamilyDelegate()
	responseData = delegate.createFromJson( jobFamily )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	jobFamily = json.loads(request.body)
	delegate = JobFamilyDelegate()
	responseData = delegate.save( jobFamily )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, jobFamilyId ):
	delegate = JobFamilyDelegate()
	responseData = delegate.delete( jobFamilyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = JobFamilyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, jobFamilyId, OrganizationId ):
	delegate = JobFamilyDelegate()
	responseData = delegate.saveOrganization( jobFamilyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, jobFamilyId ):
	delegate = JobFamilyDelegate()
	responseData = delegate.deleteOrganization( jobFamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addJobProfiles( request, jobFamilyId, JobProfilesIds ):
	delegate = JobFamilyDelegate()
	responseData = delegate.addJobProfiles( jobFamilyId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeJobProfiles( request, jobFamilyId, JobProfilesIds ):
	delegate = JobFamilyDelegate()
	responseData = delegate.removeJobProfiles( jobFamilyId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

