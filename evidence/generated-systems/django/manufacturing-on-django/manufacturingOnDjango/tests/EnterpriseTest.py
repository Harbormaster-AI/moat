import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

 #======================================================================
# 
# Encapsulates data for model Enterprise
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EnterpriseTest Declaration
#======================================================================
class EnterpriseTest (TestCase) :
	def test_crud(self) :
		enterprise = Enterprise()
		enterprise.name = "default name field value"
		enterprise.legalName = "default legalName field value"
		enterprise.registrationCountry = "default registrationCountry field value"
		enterprise.website = "default website field value"
		enterprise.taxId = "default taxId field value"
		
		delegate = EnterpriseDelegate()
		responseObj = delegate.create(enterprise)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


