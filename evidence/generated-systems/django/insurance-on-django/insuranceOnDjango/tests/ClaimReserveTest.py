import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.ClaimReserve import ClaimReserve
from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

 #======================================================================
# 
# Encapsulates data for model ClaimReserve
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimReserveTest Declaration
#======================================================================
class ClaimReserveTest (TestCase) :
	def test_crud(self) :
		claimReserve = ClaimReserve()
		claimReserve.amount = "default amount field value"
		claimReserve.setDate = datetime.datetime.now()
		claimReserve.reserveType = "default reserveType field value"
		claimReserve.status = "default status field value"
		
		delegate = ClaimReserveDelegate()
		responseObj = delegate.create(claimReserve)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


