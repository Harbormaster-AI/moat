import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.BOMItemDelegate import BOMItemDelegate

 #======================================================================
# 
# Encapsulates data for View BOMItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BOMItem index.")

def get(request, bOMItemId ):
	delegate = BOMItemDelegate()
	responseData = delegate.get( bOMItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	bOMItem = json.loads(request.body)
	delegate = BOMItemDelegate()
	responseData = delegate.createFromJson( bOMItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	bOMItem = json.loads(request.body)
	delegate = BOMItemDelegate()
	responseData = delegate.save( bOMItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, bOMItemId ):
	delegate = BOMItemDelegate()
	responseData = delegate.delete( bOMItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BOMItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBom( request, bOMItemId, BomId ):
	delegate = BOMItemDelegate()
	responseData = delegate.saveBom( bOMItemId, BomId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBom( request, bOMItemId ):
	delegate = BOMItemDelegate()
	responseData = delegate.deleteBom( bOMItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignComponent( request, bOMItemId, ComponentId ):
	delegate = BOMItemDelegate()
	responseData = delegate.saveComponent( bOMItemId, ComponentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignComponent( request, bOMItemId ):
	delegate = BOMItemDelegate()
	responseData = delegate.deleteComponent( bOMItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

