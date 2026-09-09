import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.ContentCategory import ContentCategory
from advertisingOnDjango.delegates.ContentCategoryDelegate import ContentCategoryDelegate

 #======================================================================
# 
# Encapsulates data for model ContentCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContentCategoryTest Declaration
#======================================================================
class ContentCategoryTest (TestCase) :
	def test_crud(self) :
		contentCategory = ContentCategory()
		contentCategory.code = "default code field value"
		contentCategory.name = "default name field value"
		
		delegate = ContentCategoryDelegate()
		responseObj = delegate.create(contentCategory)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


