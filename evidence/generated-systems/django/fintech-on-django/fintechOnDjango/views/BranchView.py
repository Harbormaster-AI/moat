import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.BranchDelegate import BranchDelegate

 #======================================================================
# 
# Encapsulates data for View Branch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BranchView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Branch index.")

def get(request, branchId ):
	delegate = BranchDelegate()
	responseData = delegate.get( branchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	branch = json.loads(request.body)
	delegate = BranchDelegate()
	responseData = delegate.createFromJson( branch )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	branch = json.loads(request.body)
	delegate = BranchDelegate()
	responseData = delegate.save( branch )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, branchId ):
	delegate = BranchDelegate()
	responseData = delegate.delete( branchId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BranchDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstitution( request, branchId, InstitutionId ):
	delegate = BranchDelegate()
	responseData = delegate.saveInstitution( branchId, InstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstitution( request, branchId ):
	delegate = BranchDelegate()
	responseData = delegate.deleteInstitution( branchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

