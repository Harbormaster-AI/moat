import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

 #======================================================================
# 
# Encapsulates data for model Placement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlacementTest Declaration
#======================================================================
class PlacementTest (TestCase) :
	def test_crud(self) :
		placement = Placement()
		placement.name = "default name field value"
		placement.flight = "default flight field value"
		placement.goalImpressions = 22
		
		delegate = PlacementDelegate()
		responseObj = delegate.create(placement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


