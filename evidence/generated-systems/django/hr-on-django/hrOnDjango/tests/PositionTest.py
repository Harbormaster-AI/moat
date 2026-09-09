import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Position import Position
from hrOnDjango.delegates.PositionDelegate import PositionDelegate

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
		position.positionCode = "default positionCode field value"
		position.fte = "default fte field value"
		position.status = "default status field value"
		position.workLocationType = "default workLocationType field value"
		
		delegate = PositionDelegate()
		responseObj = delegate.create(position)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


