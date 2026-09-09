import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

 #======================================================================
# 
# Encapsulates data for View Facility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FacilityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Facility index.")

def get(request, facilityId ):
	delegate = FacilityDelegate()
	responseData = delegate.get( facilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	facility = json.loads(request.body)
	delegate = FacilityDelegate()
	responseData = delegate.createFromJson( facility )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	facility = json.loads(request.body)
	delegate = FacilityDelegate()
	responseData = delegate.save( facility )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, facilityId ):
	delegate = FacilityDelegate()
	responseData = delegate.delete( facilityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FacilityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignHealthSystem( request, facilityId, HealthSystemId ):
	delegate = FacilityDelegate()
	responseData = delegate.saveHealthSystem( facilityId, HealthSystemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignHealthSystem( request, facilityId ):
	delegate = FacilityDelegate()
	responseData = delegate.deleteHealthSystem( facilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDepartments( request, facilityId, DepartmentsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addDepartments( facilityId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDepartments( request, facilityId, DepartmentsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removeDepartments( facilityId, DepartmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCareTeams( request, facilityId, CareTeamsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addCareTeams( facilityId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCareTeams( request, facilityId, CareTeamsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removeCareTeams( facilityId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLaboratories( request, facilityId, LaboratoriesIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addLaboratories( facilityId, LaboratoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLaboratories( request, facilityId, LaboratoriesIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removeLaboratories( facilityId, LaboratoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingCenters( request, facilityId, ImagingCentersIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addImagingCenters( facilityId, ImagingCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingCenters( request, facilityId, ImagingCentersIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removeImagingCenters( facilityId, ImagingCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPharmacies( request, facilityId, PharmaciesIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addPharmacies( facilityId, PharmaciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePharmacies( request, facilityId, PharmaciesIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removePharmacies( facilityId, PharmaciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, facilityId, InventoryItemsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.addInventoryItems( facilityId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, facilityId, InventoryItemsIds ):
	delegate = FacilityDelegate()
	responseData = delegate.removeInventoryItems( facilityId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

