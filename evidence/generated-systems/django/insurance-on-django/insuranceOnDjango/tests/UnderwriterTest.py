import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Underwriter import Underwriter
from insuranceOnDjango.delegates.UnderwriterDelegate import UnderwriterDelegate

 #======================================================================
# 
# Encapsulates data for model Underwriter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwriterTest Declaration
#======================================================================
class UnderwriterTest (TestCase) :
	def test_crud(self) :
		underwriter = Underwriter()
		underwriter.firstName = "default firstName field value"
		underwriter.lastName = "default lastName field value"
		underwriter.employeeId = "default employeeId field value"
		underwriter.authorityLimit = "default authorityLimit field value"
		
		delegate = UnderwriterDelegate()
		responseObj = delegate.create(underwriter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


