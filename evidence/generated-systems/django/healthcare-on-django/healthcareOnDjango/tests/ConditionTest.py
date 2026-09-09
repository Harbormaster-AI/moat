import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Condition import Condition
from healthcareOnDjango.delegates.ConditionDelegate import ConditionDelegate

 #======================================================================
# 
# Encapsulates data for model Condition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConditionTest Declaration
#======================================================================
class ConditionTest (TestCase) :
	def test_crud(self) :
		condition = Condition()
		condition.code = "default code field value"
		condition.onsetDate = datetime.datetime.now()
		condition.abatementDate = datetime.datetime.now()
		condition.clinicalStatus = "default clinicalStatus field value"
		condition.verificationStatus = "default verificationStatus field value"
		
		delegate = ConditionDelegate()
		responseObj = delegate.create(condition)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


