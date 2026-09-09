import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

 #======================================================================
# 
# Encapsulates data for View Review
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReviewView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Review index.")

def get(request, reviewId ):
	delegate = ReviewDelegate()
	responseData = delegate.get( reviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	review = json.loads(request.body)
	delegate = ReviewDelegate()
	responseData = delegate.createFromJson( review )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	review = json.loads(request.body)
	delegate = ReviewDelegate()
	responseData = delegate.save( review )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, reviewId ):
	delegate = ReviewDelegate()
	responseData = delegate.delete( reviewId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReviewDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, reviewId, ProductId ):
	delegate = ReviewDelegate()
	responseData = delegate.saveProduct( reviewId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, reviewId ):
	delegate = ReviewDelegate()
	responseData = delegate.deleteProduct( reviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, reviewId, CustomerId ):
	delegate = ReviewDelegate()
	responseData = delegate.saveCustomer( reviewId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, reviewId ):
	delegate = ReviewDelegate()
	responseData = delegate.deleteCustomer( reviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, reviewId, OrderId ):
	delegate = ReviewDelegate()
	responseData = delegate.saveOrder( reviewId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, reviewId ):
	delegate = ReviewDelegate()
	responseData = delegate.deleteOrder( reviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

