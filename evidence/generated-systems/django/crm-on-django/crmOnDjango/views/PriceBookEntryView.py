import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

 #======================================================================
# 
# Encapsulates data for View PriceBookEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookEntryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PriceBookEntry index.")

def get(request, priceBookEntryId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.get( priceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	priceBookEntry = json.loads(request.body)
	delegate = PriceBookEntryDelegate()
	responseData = delegate.createFromJson( priceBookEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	priceBookEntry = json.loads(request.body)
	delegate = PriceBookEntryDelegate()
	responseData = delegate.save( priceBookEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, priceBookEntryId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.delete( priceBookEntryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBook( request, priceBookEntryId, PriceBookId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.savePriceBook( priceBookEntryId, PriceBookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBook( request, priceBookEntryId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.deletePriceBook( priceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, priceBookEntryId, ProductId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.saveProduct( priceBookEntryId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, priceBookEntryId ):
	delegate = PriceBookEntryDelegate()
	responseData = delegate.deleteProduct( priceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

