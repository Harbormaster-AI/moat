import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

 #======================================================================
# 
# Encapsulates data for View Insurer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Insurer index.")

def get(request, insurerId ):
	delegate = InsurerDelegate()
	responseData = delegate.get( insurerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insurer = json.loads(request.body)
	delegate = InsurerDelegate()
	responseData = delegate.createFromJson( insurer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insurer = json.loads(request.body)
	delegate = InsurerDelegate()
	responseData = delegate.save( insurer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insurerId ):
	delegate = InsurerDelegate()
	responseData = delegate.delete( insurerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsurerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, insurerId, ProductsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.addProducts( insurerId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, insurerId, ProductsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.removeProducts( insurerId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDistributionPartners( request, insurerId, DistributionPartnersIds ):
	delegate = InsurerDelegate()
	responseData = delegate.addDistributionPartners( insurerId, DistributionPartnersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDistributionPartners( request, insurerId, DistributionPartnersIds ):
	delegate = InsurerDelegate()
	responseData = delegate.removeDistributionPartners( insurerId, DistributionPartnersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, insurerId, PoliciesIds ):
	delegate = InsurerDelegate()
	responseData = delegate.addPolicies( insurerId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, insurerId, PoliciesIds ):
	delegate = InsurerDelegate()
	responseData = delegate.removePolicies( insurerId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, insurerId, ClaimsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.addClaims( insurerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, insurerId, ClaimsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.removeClaims( insurerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReinsuranceAgreements( request, insurerId, ReinsuranceAgreementsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.addReinsuranceAgreements( insurerId, ReinsuranceAgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReinsuranceAgreements( request, insurerId, ReinsuranceAgreementsIds ):
	delegate = InsurerDelegate()
	responseData = delegate.removeReinsuranceAgreements( insurerId, ReinsuranceAgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

