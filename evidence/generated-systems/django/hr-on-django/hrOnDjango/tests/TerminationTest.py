import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Termination import Termination
from hrOnDjango.delegates.TerminationDelegate import TerminationDelegate

 #======================================================================
# 
# Encapsulates data for model Termination
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminationTest Declaration
#======================================================================
class TerminationTest (TestCase) :
	def test_crud(self) :
		termination = Termination()
		termination.terminationNumber = "default terminationNumber field value"
		termination.terminationDate = datetime.datetime.now()
		termination.notes = "default notes field value"
		termination.eligibleForRehire = False
		termination.reason = "default reason field value"
		termination.type = "default type field value"
		
		delegate = TerminationDelegate()
		responseObj = delegate.create(termination)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


