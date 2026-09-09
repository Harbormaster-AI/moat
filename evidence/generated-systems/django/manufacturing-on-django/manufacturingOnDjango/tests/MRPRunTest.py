import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.MRPRun import MRPRun
from manufacturingOnDjango.delegates.MRPRunDelegate import MRPRunDelegate

 #======================================================================
# 
# Encapsulates data for model MRPRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MRPRunTest Declaration
#======================================================================
class MRPRunTest (TestCase) :
	def test_crud(self) :
		mRPRun = MRPRun()
		mRPRun.runNumber = "default runNumber field value"
		mRPRun.runDateTime = "default runDateTime field value"
		mRPRun.planningHorizonDays = 22
		mRPRun.status = "default status field value"
		
		delegate = MRPRunDelegate()
		responseObj = delegate.create(mRPRun)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


