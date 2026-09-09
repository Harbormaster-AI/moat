import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Nonconformance import Nonconformance
from manufacturingOnDjango.delegates.NonconformanceDelegate import NonconformanceDelegate

 #======================================================================
# 
# Encapsulates data for model Nonconformance
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NonconformanceTest Declaration
#======================================================================
class NonconformanceTest (TestCase) :
	def test_crud(self) :
		nonconformance = Nonconformance()
		nonconformance.ncNumber = "default ncNumber field value"
		nonconformance.description = "default description field value"
		nonconformance.containmentAction = "default containmentAction field value"
		nonconformance.ncType = "default ncType field value"
		nonconformance.severity = "default severity field value"
		nonconformance.status = "default status field value"
		
		delegate = NonconformanceDelegate()
		responseObj = delegate.create(nonconformance)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


