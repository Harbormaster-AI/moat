import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

 #======================================================================
# 
# Encapsulates data for View ImagingOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ImagingOrder index.")

def get(request, imagingOrderId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.get( imagingOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	imagingOrder = json.loads(request.body)
	delegate = ImagingOrderDelegate()
	responseData = delegate.createFromJson( imagingOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	imagingOrder = json.loads(request.body)
	delegate = ImagingOrderDelegate()
	responseData = delegate.save( imagingOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, imagingOrderId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.delete( imagingOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ImagingOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, imagingOrderId, OrderId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.saveOrder( imagingOrderId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, imagingOrderId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.deleteOrder( imagingOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignImagingCenter( request, imagingOrderId, ImagingCenterId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.saveImagingCenter( imagingOrderId, ImagingCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignImagingCenter( request, imagingOrderId ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.deleteImagingCenter( imagingOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, imagingOrderId, ReportsIds ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.addReports( imagingOrderId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, imagingOrderId, ReportsIds ):
	delegate = ImagingOrderDelegate()
	responseData = delegate.removeReports( imagingOrderId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

