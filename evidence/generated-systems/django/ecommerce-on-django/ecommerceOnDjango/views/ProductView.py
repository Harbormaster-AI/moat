import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

 #======================================================================
# 
# Encapsulates data for View Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Product index.")

def get(request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.get( productId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	product = json.loads(request.body)
	delegate = ProductDelegate()
	responseData = delegate.createFromJson( product )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	product = json.loads(request.body)
	delegate = ProductDelegate()
	responseData = delegate.save( product )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.delete( productId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBrand( request, productId, BrandId ):
	delegate = ProductDelegate()
	responseData = delegate.saveBrand( productId, BrandId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBrand( request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.deleteBrand( productId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSeller( request, productId, SellerId ):
	delegate = ProductDelegate()
	responseData = delegate.saveSeller( productId, SellerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSeller( request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.deleteSeller( productId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCategories( request, productId, CategoriesIds ):
	delegate = ProductDelegate()
	responseData = delegate.addCategories( productId, CategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCategories( request, productId, CategoriesIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeCategories( productId, CategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, productId, VariantsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addVariants( productId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, productId, VariantsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeVariants( productId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMediaAssets( request, productId, MediaAssetsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addMediaAssets( productId, MediaAssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMediaAssets( request, productId, MediaAssetsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeMediaAssets( productId, MediaAssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReviews( request, productId, ReviewsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addReviews( productId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReviews( request, productId, ReviewsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeReviews( productId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

