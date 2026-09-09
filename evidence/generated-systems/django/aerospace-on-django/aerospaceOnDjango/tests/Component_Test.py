import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Component_ import Component_
from aerospaceOnDjango.delegates.Component_Delegate import Component_Delegate

 #======================================================================
# 
# Encapsulates data for model Component_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Component_Test Declaration
#======================================================================
class Component_Test (TestCase) :
	def test_crud(self) :
		component_ = Component_()
		component_.partNumber = "default partNumber field value"
		component_.name = "default name field value"
		component_.componentCategory = "default componentCategory field value"
		component_.serializationMethod = "default serializationMethod field value"
		
		delegate = Component_Delegate()
		responseObj = delegate.create(component_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


