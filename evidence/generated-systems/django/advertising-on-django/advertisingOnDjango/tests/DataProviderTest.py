import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.DataProvider import DataProvider
from advertisingOnDjango.delegates.DataProviderDelegate import DataProviderDelegate

 #======================================================================
# 
# Encapsulates data for model DataProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProviderTest Declaration
#======================================================================
class DataProviderTest (TestCase) :
	def test_crud(self) :
		dataProvider = DataProvider()
		dataProvider.name = "default name field value"
		dataProvider.website = "default website field value"
		dataProvider.providerType = "default providerType field value"
		
		delegate = DataProviderDelegate()
		responseObj = delegate.create(dataProvider)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


