import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.delegates.TagDelegate import TagDelegate

 #======================================================================
# 
# Encapsulates data for model Tag
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TagTest Declaration
#======================================================================
class TagTest (TestCase) :
	def test_crud(self) :
		tag = Tag()
		tag.name = "default name field value"
		tag.category = "default category field value"
		
		delegate = TagDelegate()
		responseObj = delegate.create(tag)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


