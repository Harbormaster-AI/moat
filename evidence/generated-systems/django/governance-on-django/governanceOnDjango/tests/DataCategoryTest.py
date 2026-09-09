import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.DataCategory import DataCategory
from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

 #======================================================================
# 
# Encapsulates data for model DataCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataCategoryTest Declaration
#======================================================================
class DataCategoryTest (TestCase) :
	def test_crud(self) :
		dataCategory = DataCategory()
		dataCategory.name = "default name field value"
		dataCategory.description = "default description field value"
		dataCategory.classification = "default classification field value"
		
		delegate = DataCategoryDelegate()
		responseObj = delegate.create(dataCategory)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


