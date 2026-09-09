import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

 #======================================================================
# 
# Encapsulates data for model Agency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgencyTest Declaration
#======================================================================
class AgencyTest (TestCase) :
	def test_crud(self) :
		agency = Agency()
		agency.name = "default name field value"
		agency.legalName = "default legalName field value"
		agency.headquartersCountry = "default headquartersCountry field value"
		agency.website = "default website field value"
		
		delegate = AgencyDelegate()
		responseObj = delegate.create(agency)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


