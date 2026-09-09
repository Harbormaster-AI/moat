import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.UoMConversion import UoMConversion
from inventoryOnDjango.delegates.UoMConversionDelegate import UoMConversionDelegate

 #======================================================================
# 
# Encapsulates data for model UoMConversion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UoMConversionTest Declaration
#======================================================================
class UoMConversionTest (TestCase) :
	def test_crud(self) :
		uoMConversion = UoMConversion()
		uoMConversion.factor = "default factor field value"
		uoMConversion.precision = 22
		uoMConversion.fromUnit = "default fromUnit field value"
		uoMConversion.toUnit = "default toUnit field value"
		
		delegate = UoMConversionDelegate()
		responseObj = delegate.create(uoMConversion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


