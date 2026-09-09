import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Adjuster import Adjuster
from insuranceOnDjango.delegates.AdjusterDelegate import AdjusterDelegate

 #======================================================================
# 
# Encapsulates data for model Adjuster
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjusterTest Declaration
#======================================================================
class AdjusterTest (TestCase) :
	def test_crud(self) :
		adjuster = Adjuster()
		adjuster.firstName = "default firstName field value"
		adjuster.lastName = "default lastName field value"
		adjuster.licenseNumber = "default licenseNumber field value"
		adjuster.adjusterType = "default adjusterType field value"
		
		delegate = AdjusterDelegate()
		responseObj = delegate.create(adjuster)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


