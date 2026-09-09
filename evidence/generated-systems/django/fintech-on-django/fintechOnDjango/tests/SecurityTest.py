import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Security import Security
from fintechOnDjango.delegates.SecurityDelegate import SecurityDelegate

 #======================================================================
# 
# Encapsulates data for model Security
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SecurityTest Declaration
#======================================================================
class SecurityTest (TestCase) :
	def test_crud(self) :
		security = Security()
		security.symbol = "default symbol field value"
		security.isin = "default isin field value"
		security.cusip = "default cusip field value"
		security.currency = "default currency field value"
		security.securityType = "default securityType field value"
		
		delegate = SecurityDelegate()
		responseObj = delegate.create(security)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


