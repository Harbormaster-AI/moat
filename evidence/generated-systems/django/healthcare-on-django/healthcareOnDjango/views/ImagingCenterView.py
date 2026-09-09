import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

 #======================================================================
# 
# Encapsulates data for View ImagingCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingCenterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ImagingCenter index.")

def get(request, imagingCenterId ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.get( imagingCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	imagingCenter = json.loads(request.body)
	delegate = ImagingCenterDelegate()
	responseData = delegate.createFromJson( imagingCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	imagingCenter = json.loads(request.body)
	delegate = ImagingCenterDelegate()
	responseData = delegate.save( imagingCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, imagingCenterId ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.delete( imagingCenterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ImagingCenterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, imagingCenterId, FacilityId ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.saveFacility( imagingCenterId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, imagingCenterId ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.deleteFacility( imagingCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingOrders( request, imagingCenterId, ImagingOrdersIds ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.addImagingOrders( imagingCenterId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingOrders( request, imagingCenterId, ImagingOrdersIds ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.removeImagingOrders( imagingCenterId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingReports( request, imagingCenterId, ImagingReportsIds ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.addImagingReports( imagingCenterId, ImagingReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingReports( request, imagingCenterId, ImagingReportsIds ):
	delegate = ImagingCenterDelegate()
	responseData = delegate.removeImagingReports( imagingCenterId, ImagingReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

