import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AirworthinessDirective import AirworthinessDirective
from aerospaceOnDjango.delegates.AirworthinessDirectiveDelegate import AirworthinessDirectiveDelegate

 #======================================================================
# 
# Encapsulates data for model AirworthinessDirective
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AirworthinessDirectiveTest Declaration
#======================================================================
class AirworthinessDirectiveTest (TestCase) :
	def test_crud(self) :
		airworthinessDirective = AirworthinessDirective()
		airworthinessDirective.directiveNumber = "default directiveNumber field value"
		airworthinessDirective.title = "default title field value"
		
		delegate = AirworthinessDirectiveDelegate()
		responseObj = delegate.create(airworthinessDirective)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


