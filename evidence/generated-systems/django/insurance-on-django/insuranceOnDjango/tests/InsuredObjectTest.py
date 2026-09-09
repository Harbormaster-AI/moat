import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

 #======================================================================
# 
# Encapsulates data for model InsuredObject
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuredObjectTest Declaration
#======================================================================
class InsuredObjectTest (TestCase) :
	def test_crud(self) :
		insuredObject = InsuredObject()
		insuredObject.description = "default description field value"
		insuredObject.serialOrId = "default serialOrId field value"
		insuredObject.primaryAddress = "default primaryAddress field value"
		insuredObject.objectType = "default objectType field value"
		
		delegate = InsuredObjectDelegate()
		responseObj = delegate.create(insuredObject)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


