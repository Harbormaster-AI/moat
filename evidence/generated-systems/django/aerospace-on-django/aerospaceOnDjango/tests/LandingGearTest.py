import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.LandingGear import LandingGear
from aerospaceOnDjango.delegates.LandingGearDelegate import LandingGearDelegate

 #======================================================================
# 
# Encapsulates data for model LandingGear
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LandingGearTest Declaration
#======================================================================
class LandingGearTest (TestCase) :
	def test_crud(self) :
		landingGear = LandingGear()
		landingGear.supplierPartNumber = "default supplierPartNumber field value"
		landingGear.gearType = "default gearType field value"
		
		delegate = LandingGearDelegate()
		responseObj = delegate.create(landingGear)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


