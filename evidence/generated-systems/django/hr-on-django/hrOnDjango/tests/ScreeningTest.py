import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Screening import Screening
from hrOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

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
		screening.name = "default name field value"
		screening.completedDate = datetime.datetime.now()
		screening.status = "default status field value"
		
		delegate = ScreeningDelegate()
		responseObj = delegate.create(screening)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


