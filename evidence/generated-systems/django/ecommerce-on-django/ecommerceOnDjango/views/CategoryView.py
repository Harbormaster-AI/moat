import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

 #======================================================================
# 
# Encapsulates data for View Category
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CategoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Category index.")

def get(request, categoryId ):
	delegate = CategoryDelegate()
	responseData = delegate.get( categoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	category = json.loads(request.body)
	delegate = CategoryDelegate()
	responseData = delegate.createFromJson( category )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	category = json.loads(request.body)
	delegate = CategoryDelegate()
	responseData = delegate.save( category )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, categoryId ):
	delegate = CategoryDelegate()
	responseData = delegate.delete( categoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CategoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCatalog( request, categoryId, CatalogId ):
	delegate = CategoryDelegate()
	responseData = delegate.saveCatalog( categoryId, CatalogId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCatalog( request, categoryId ):
	delegate = CategoryDelegate()
	responseData = delegate.deleteCatalog( categoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentCategory( request, categoryId, ParentCategoryId ):
	delegate = CategoryDelegate()
	responseData = delegate.saveParentCategory( categoryId, ParentCategoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentCategory( request, categoryId ):
	delegate = CategoryDelegate()
	responseData = delegate.deleteParentCategory( categoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubcategories( request, categoryId, SubcategoriesIds ):
	delegate = CategoryDelegate()
	responseData = delegate.addSubcategories( categoryId, SubcategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubcategories( request, categoryId, SubcategoriesIds ):
	delegate = CategoryDelegate()
	responseData = delegate.removeSubcategories( categoryId, SubcategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, categoryId, ProductsIds ):
	delegate = CategoryDelegate()
	responseData = delegate.addProducts( categoryId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, categoryId, ProductsIds ):
	delegate = CategoryDelegate()
	responseData = delegate.removeProducts( categoryId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

