import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.QualitySpecificationDelegate import QualitySpecificationDelegate

 #======================================================================
# 
# Encapsulates data for View QualitySpecification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualitySpecificationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the QualitySpecification index.")

def get(request, qualitySpecificationId ):
	delegate = QualitySpecificationDelegate()
	responseData = delegate.get( qualitySpecificationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	qualitySpecification = json.loads(request.body)
	delegate = QualitySpecificationDelegate()
	responseData = delegate.createFromJson( qualitySpecification )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	qualitySpecification = json.loads(request.body)
	delegate = QualitySpecificationDelegate()
	responseData = delegate.save( qualitySpecification )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, qualitySpecificationId ):
	delegate = QualitySpecificationDelegate()
	responseData = delegate.delete( qualitySpecificationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QualitySpecificationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, qualitySpecificationId, ItemId ):
	delegate = QualitySpecificationDelegate()
	responseData = delegate.saveItem( qualitySpecificationId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, qualitySpecificationId ):
	delegate = QualitySpecificationDelegate()
	responseData = delegate.deleteItem( qualitySpecificationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

