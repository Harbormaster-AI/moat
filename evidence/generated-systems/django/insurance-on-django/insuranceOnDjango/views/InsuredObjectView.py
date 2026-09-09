import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

 #======================================================================
# 
# Encapsulates data for View InsuredObject
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuredObjectView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InsuredObject index.")

def get(request, insuredObjectId ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.get( insuredObjectId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insuredObject = json.loads(request.body)
	delegate = InsuredObjectDelegate()
	responseData = delegate.createFromJson( insuredObject )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insuredObject = json.loads(request.body)
	delegate = InsuredObjectDelegate()
	responseData = delegate.save( insuredObject )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insuredObjectId ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.delete( insuredObjectId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsuredObjectDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, insuredObjectId, PolicyId ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.savePolicy( insuredObjectId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, insuredObjectId ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.deletePolicy( insuredObjectId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoverages( request, insuredObjectId, CoveragesIds ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.addCoverages( insuredObjectId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoverages( request, insuredObjectId, CoveragesIds ):
	delegate = InsuredObjectDelegate()
	responseData = delegate.removeCoverages( insuredObjectId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

