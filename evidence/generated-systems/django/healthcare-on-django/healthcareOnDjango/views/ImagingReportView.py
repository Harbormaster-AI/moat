import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

 #======================================================================
# 
# Encapsulates data for View ImagingReport
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingReportView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ImagingReport index.")

def get(request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.get( imagingReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	imagingReport = json.loads(request.body)
	delegate = ImagingReportDelegate()
	responseData = delegate.createFromJson( imagingReport )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	imagingReport = json.loads(request.body)
	delegate = ImagingReportDelegate()
	responseData = delegate.save( imagingReport )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.delete( imagingReportId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ImagingReportDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignImagingOrder( request, imagingReportId, ImagingOrderId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.saveImagingOrder( imagingReportId, ImagingOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignImagingOrder( request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.deleteImagingOrder( imagingReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClinician( request, imagingReportId, ClinicianId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.saveClinician( imagingReportId, ClinicianId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClinician( request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.deleteClinician( imagingReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, imagingReportId, EncounterId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.saveEncounter( imagingReportId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.deleteEncounter( imagingReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignImagingCenter( request, imagingReportId, ImagingCenterId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.saveImagingCenter( imagingReportId, ImagingCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignImagingCenter( request, imagingReportId ):
	delegate = ImagingReportDelegate()
	responseData = delegate.deleteImagingCenter( imagingReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

