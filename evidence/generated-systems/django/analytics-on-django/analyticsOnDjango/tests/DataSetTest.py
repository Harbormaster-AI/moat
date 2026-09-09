import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

 #======================================================================
# 
# Encapsulates data for model DataSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSetTest Declaration
#======================================================================
class DataSetTest (TestCase) :
	def test_crud(self) :
		dataSet = DataSet()
		dataSet.name = "default name field value"
		dataSet.schemaVersion = "default schemaVersion field value"
		dataSet.refreshSchedule = "default refreshSchedule field value"
		dataSet.sensitive = False
		dataSet.dataFormat = "default dataFormat field value"
		
		delegate = DataSetDelegate()
		responseObj = delegate.create(dataSet)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


