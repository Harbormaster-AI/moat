import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Category import Category
from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

 #======================================================================
# 
# Encapsulates data for model Category
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CategoryTest Declaration
#======================================================================
class CategoryTest (TestCase) :
	def test_crud(self) :
		category = Category()
		category.name = "default name field value"
		category.slug = "default slug field value"
		category.position = 22
		category.asActive = False
		
		delegate = CategoryDelegate()
		responseObj = delegate.create(category)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


