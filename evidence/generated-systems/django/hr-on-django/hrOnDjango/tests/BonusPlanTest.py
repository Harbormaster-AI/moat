import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.BonusPlan import BonusPlan
from hrOnDjango.delegates.BonusPlanDelegate import BonusPlanDelegate

 #======================================================================
# 
# Encapsulates data for model BonusPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BonusPlanTest Declaration
#======================================================================
class BonusPlanTest (TestCase) :
	def test_crud(self) :
		bonusPlan = BonusPlan()
		bonusPlan.name = "default name field value"
		bonusPlan.targetPercentage = "default targetPercentage field value"
		
		delegate = BonusPlanDelegate()
		responseObj = delegate.create(bonusPlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


