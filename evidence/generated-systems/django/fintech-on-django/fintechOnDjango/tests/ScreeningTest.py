import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Screening import Screening
from fintechOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

 #======================================================================
# 
# Encapsulates data for model Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningTest Declaration
#======================================================================
class ScreeningTest (TestCase) :
	def test_crud(self) :
		screening = Screening()
		screening.score = "default score field value"
		screening.screenedAt = "default screenedAt field value"
		screening.screeningType = "default screeningType field value"
		screening.status = "default status field value"
		
		delegate = ScreeningDelegate()
		responseObj = delegate.create(screening)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


