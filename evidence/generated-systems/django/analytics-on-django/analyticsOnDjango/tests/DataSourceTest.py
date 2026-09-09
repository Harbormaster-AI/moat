import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.DataSource import DataSource
from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

 #======================================================================
# 
# Encapsulates data for model DataSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSourceTest Declaration
#======================================================================
class DataSourceTest (TestCase) :
	def test_crud(self) :
		dataSource = DataSource()
		dataSource.name = "default name field value"
		dataSource.connection = "default connection field value"
		dataSource.streaming = False
		dataSource.sourceType = "default sourceType field value"
		dataSource.format = "default format field value"
		
		delegate = DataSourceDelegate()
		responseObj = delegate.create(dataSource)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


