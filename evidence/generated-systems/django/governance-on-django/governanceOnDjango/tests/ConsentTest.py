import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Consent import Consent
from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

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
		consent.subjectIdentifier = "default subjectIdentifier field value"
		consent.captureDate = datetime.datetime.now()
		consent.expiryDate = datetime.datetime.now()
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


