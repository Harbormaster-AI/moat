import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Operation import Operation
from manufacturingOnDjango.delegates.OperationDelegate import OperationDelegate

 #======================================================================
# 
# Encapsulates data for model Operation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperationTest Declaration
#======================================================================
class OperationTest (TestCase) :
	def test_crud(self) :
		operation = Operation()
		operation.operationNumber = "default operationNumber field value"
		operation.name = "default name field value"
		operation.setupTime = "default setupTime field value"
		operation.standardCycleTime = "default standardCycleTime field value"
		operation.operationType = "default operationType field value"
		
		delegate = OperationDelegate()
		responseObj = delegate.create(operation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


