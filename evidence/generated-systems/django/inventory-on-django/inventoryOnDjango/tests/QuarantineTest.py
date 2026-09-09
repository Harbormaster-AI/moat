import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.Quarantine import Quarantine
from inventoryOnDjango.delegates.QuarantineDelegate import QuarantineDelegate

 #======================================================================
# 
# Encapsulates data for model Quarantine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuarantineTest Declaration
#======================================================================
class QuarantineTest (TestCase) :
	def test_crud(self) :
		quarantine = Quarantine()
		quarantine.reason = "default reason field value"
		quarantine.startedAt = datetime.datetime.now()
		quarantine.releasedAt = datetime.datetime.now()
		quarantine.disposition = "default disposition field value"
		
		delegate = QuarantineDelegate()
		responseObj = delegate.create(quarantine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


