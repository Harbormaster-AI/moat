import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.BusinessUnit import BusinessUnit
from manufacturingOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

 #======================================================================
# 
# Encapsulates data for model BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnitTest Declaration
#======================================================================
class BusinessUnitTest (TestCase) :
	def test_crud(self) :
		businessUnit = BusinessUnit()
		businessUnit.name = "default name field value"
		businessUnit.code = "default code field value"
		businessUnit.category = "default category field value"
		
		delegate = BusinessUnitDelegate()
		responseObj = delegate.create(businessUnit)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


