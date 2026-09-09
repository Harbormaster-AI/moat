import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.BIQuery import BIQuery
from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

 #======================================================================
# 
# Encapsulates data for model BIQuery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BIQueryTest Declaration
#======================================================================
class BIQueryTest (TestCase) :
	def test_crud(self) :
		bIQuery = BIQuery()
		bIQuery.name = "default name field value"
		bIQuery.text = "default text field value"
		bIQuery.dialect = "default dialect field value"
		
		delegate = BIQueryDelegate()
		responseObj = delegate.create(bIQuery)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


