import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Territory import Territory
from crmOnDjango.delegates.TerritoryDelegate import TerritoryDelegate

 #======================================================================
# 
# Encapsulates data for model Territory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerritoryTest Declaration
#======================================================================
class TerritoryTest (TestCase) :
	def test_crud(self) :
		territory = Territory()
		territory.name = "default name field value"
		territory.region = "default region field value"
		territory.territoryType = "default territoryType field value"
		
		delegate = TerritoryDelegate()
		responseObj = delegate.create(territory)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


