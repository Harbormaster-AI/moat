import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.SubrogationRecovery import SubrogationRecovery
from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

 #======================================================================
# 
# Encapsulates data for model SubrogationRecovery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubrogationRecoveryTest Declaration
#======================================================================
class SubrogationRecoveryTest (TestCase) :
	def test_crud(self) :
		subrogationRecovery = SubrogationRecovery()
		subrogationRecovery.recoveryReference = "default recoveryReference field value"
		subrogationRecovery.amount = "default amount field value"
		subrogationRecovery.recoveryDate = datetime.datetime.now()
		subrogationRecovery.status = "default status field value"
		
		delegate = SubrogationRecoveryDelegate()
		responseObj = delegate.create(subrogationRecovery)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


