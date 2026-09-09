import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.CompetencyRating import CompetencyRating
from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

 #======================================================================
# 
# Encapsulates data for model CompetencyRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyRatingTest Declaration
#======================================================================
class CompetencyRatingTest (TestCase) :
	def test_crud(self) :
		competencyRating = CompetencyRating()
		competencyRating.comment = "default comment field value"
		competencyRating.rating = "default rating field value"
		
		delegate = CompetencyRatingDelegate()
		responseObj = delegate.create(competencyRating)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


