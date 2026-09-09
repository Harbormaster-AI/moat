import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CatalogDelegate import CatalogDelegate

 #======================================================================
# 
# Encapsulates data for View Catalog
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CatalogView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Catalog index.")

def get(request, catalogId ):
	delegate = CatalogDelegate()
	responseData = delegate.get( catalogId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	catalog = json.loads(request.body)
	delegate = CatalogDelegate()
	responseData = delegate.createFromJson( catalog )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	catalog = json.loads(request.body)
	delegate = CatalogDelegate()
	responseData = delegate.save( catalog )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, catalogId ):
	delegate = CatalogDelegate()
	responseData = delegate.delete( catalogId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CatalogDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChannel( request, catalogId, ChannelId ):
	delegate = CatalogDelegate()
	responseData = delegate.saveChannel( catalogId, ChannelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChannel( request, catalogId ):
	delegate = CatalogDelegate()
	responseData = delegate.deleteChannel( catalogId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCategories( request, catalogId, CategoriesIds ):
	delegate = CatalogDelegate()
	responseData = delegate.addCategories( catalogId, CategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCategories( request, catalogId, CategoriesIds ):
	delegate = CatalogDelegate()
	responseData = delegate.removeCategories( catalogId, CategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

