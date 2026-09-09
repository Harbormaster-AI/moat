import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.KPI import KPI
from advertisingOnDjango.delegates.KPIDelegate import KPIDelegate

 #======================================================================
# 
# Encapsulates data for model KPI
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KPITest Declaration
#======================================================================
class KPITest (TestCase) :
	def test_crud(self) :
		kPI = KPI()
		kPI.targetValue = "default targetValue field value"
		kPI.metricType = "default metricType field value"
		
		delegate = KPIDelegate()
		responseObj = delegate.create(kPI)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


