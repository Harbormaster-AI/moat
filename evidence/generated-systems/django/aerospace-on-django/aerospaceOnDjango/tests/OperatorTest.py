import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

 #======================================================================
# 
# Encapsulates data for model Operator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperatorTest Declaration
#======================================================================
class OperatorTest (TestCase) :
	def test_crud(self) :
		operator = Operator()
		operator.name = "default name field value"
		operator.icaoDesignator = "default icaoDesignator field value"
		operator.operatorType = "default operatorType field value"
		
		delegate = OperatorDelegate()
		responseObj = delegate.create(operator)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


