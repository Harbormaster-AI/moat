import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.ContentCategoryDelegate import ContentCategoryDelegate

 #======================================================================
# 
# Encapsulates data for View ContentCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContentCategoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ContentCategory index.")

def get(request, contentCategoryId ):
	delegate = ContentCategoryDelegate()
	responseData = delegate.get( contentCategoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	contentCategory = json.loads(request.body)
	delegate = ContentCategoryDelegate()
	responseData = delegate.createFromJson( contentCategory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	contentCategory = json.loads(request.body)
	delegate = ContentCategoryDelegate()
	responseData = delegate.save( contentCategory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, contentCategoryId ):
	delegate = ContentCategoryDelegate()
	responseData = delegate.delete( contentCategoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ContentCategoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

