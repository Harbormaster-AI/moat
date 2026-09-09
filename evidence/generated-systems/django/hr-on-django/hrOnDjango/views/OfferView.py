import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.OfferDelegate import OfferDelegate

 #======================================================================
# 
# Encapsulates data for View Offer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OfferView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Offer index.")

def get(request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.get( offerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	offer = json.loads(request.body)
	delegate = OfferDelegate()
	responseData = delegate.createFromJson( offer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	offer = json.loads(request.body)
	delegate = OfferDelegate()
	responseData = delegate.save( offer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.delete( offerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OfferDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRequisition( request, offerId, RequisitionId ):
	delegate = OfferDelegate()
	responseData = delegate.saveRequisition( offerId, RequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRequisition( request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.deleteRequisition( offerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCandidate( request, offerId, CandidateId ):
	delegate = OfferDelegate()
	responseData = delegate.saveCandidate( offerId, CandidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCandidate( request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.deleteCandidate( offerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApprovedBy( request, offerId, ApprovedById ):
	delegate = OfferDelegate()
	responseData = delegate.saveApprovedBy( offerId, ApprovedById )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApprovedBy( request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.deleteApprovedBy( offerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContract( request, offerId, ContractId ):
	delegate = OfferDelegate()
	responseData = delegate.saveContract( offerId, ContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContract( request, offerId ):
	delegate = OfferDelegate()
	responseData = delegate.deleteContract( offerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

