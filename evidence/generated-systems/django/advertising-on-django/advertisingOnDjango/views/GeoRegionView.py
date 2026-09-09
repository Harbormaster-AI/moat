import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

 #======================================================================
# 
# Encapsulates data for View GeoRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GeoRegionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GeoRegion index.")

def get(request, geoRegionId ):
	delegate = GeoRegionDelegate()
	responseData = delegate.get( geoRegionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	geoRegion = json.loads(request.body)
	delegate = GeoRegionDelegate()
	responseData = delegate.createFromJson( geoRegion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	geoRegion = json.loads(request.body)
	delegate = GeoRegionDelegate()
	responseData = delegate.save( geoRegion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, geoRegionId ):
	delegate = GeoRegionDelegate()
	responseData = delegate.delete( geoRegionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GeoRegionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParent( request, geoRegionId, ParentId ):
	delegate = GeoRegionDelegate()
	responseData = delegate.saveParent( geoRegionId, ParentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParent( request, geoRegionId ):
	delegate = GeoRegionDelegate()
	responseData = delegate.deleteParent( geoRegionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChildren( request, geoRegionId, ChildrenIds ):
	delegate = GeoRegionDelegate()
	responseData = delegate.addChildren( geoRegionId, ChildrenIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChildren( request, geoRegionId, ChildrenIds ):
	delegate = GeoRegionDelegate()
	responseData = delegate.removeChildren( geoRegionId, ChildrenIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

