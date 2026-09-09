import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Consent import Consent
from fintechOnDjango.delegates.ConsentDelegate import ConsentDelegate

 #======================================================================
# 
# Encapsulates data for model Consent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConsentTest Declaration
#======================================================================
class ConsentTest (TestCase) :
	def test_crud(self) :
		consent = Consent()
		consent.grantedAt = "default grantedAt field value"
		consent.expiresAt = "default expiresAt field value"
		consent.scope = "default scope field value"
		consent.consentType = "default consentType field value"
		consent.status = "default status field value"
		
		delegate = ConsentDelegate()
		responseObj = delegate.create(consent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


