import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Position import Position
from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

 #======================================================================
# 
# Encapsulates data for model Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionTest Declaration
#======================================================================
class PositionTest (TestCase) :
	def test_crud(self) :
		position = Position()
		position.quantity = "default quantity field value"
		position.averageCost = "default averageCost field value"
		position.marketValue = "default marketValue field value"
		
		delegate = PositionDelegate()
		responseObj = delegate.create(position)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


