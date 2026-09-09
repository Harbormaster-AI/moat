import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Competency import Competency
from hrOnDjango.delegates.CompetencyDelegate import CompetencyDelegate

 #======================================================================
# 
# Encapsulates data for model Competency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyTest Declaration
#======================================================================
class CompetencyTest (TestCase) :
	def test_crud(self) :
		competency = Competency()
		competency.name = "default name field value"
		competency.category = "default category field value"
		
		delegate = CompetencyDelegate()
		responseObj = delegate.create(competency)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


