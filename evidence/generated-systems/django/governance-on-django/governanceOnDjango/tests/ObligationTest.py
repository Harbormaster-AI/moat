import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

 #======================================================================
# 
# Encapsulates data for model Obligation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObligationTest Declaration
#======================================================================
class ObligationTest (TestCase) :
	def test_crud(self) :
		obligation = Obligation()
		obligation.referenceNumber = "default referenceNumber field value"
		obligation.descriptionText = "default descriptionText field value"
		obligation.obligationType = "default obligationType field value"
		obligation.reviewFrequency = "default reviewFrequency field value"
		
		delegate = ObligationDelegate()
		responseObj = delegate.create(obligation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


