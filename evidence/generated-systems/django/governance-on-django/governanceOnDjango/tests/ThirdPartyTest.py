import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

 #======================================================================
# 
# Encapsulates data for model ThirdParty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyTest Declaration
#======================================================================
class ThirdPartyTest (TestCase) :
	def test_crud(self) :
		thirdParty = ThirdParty()
		thirdParty.name = "default name field value"
		thirdParty.country = "default country field value"
		thirdParty.contactEmail = "default contactEmail field value"
		thirdParty.thirdPartyType = "default thirdPartyType field value"
		thirdParty.criticality = "default criticality field value"
		
		delegate = ThirdPartyDelegate()
		responseObj = delegate.create(thirdParty)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


