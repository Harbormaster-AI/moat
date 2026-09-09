import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Organization import Organization
from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

 #======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationTest Declaration
#======================================================================
class OrganizationTest (TestCase) :
	def test_crud(self) :
		organization = Organization()
		organization.name = "default name field value"
		organization.defaultCurrency = "default defaultCurrency field value"
		organization.defaultLocale = "default defaultLocale field value"
		organization.website = "default website field value"
		
		delegate = OrganizationDelegate()
		responseObj = delegate.create(organization)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


