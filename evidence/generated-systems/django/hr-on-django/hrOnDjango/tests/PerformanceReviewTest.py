import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PerformanceReview import PerformanceReview
from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

 #======================================================================
# 
# Encapsulates data for model PerformanceReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceReviewTest Declaration
#======================================================================
class PerformanceReviewTest (TestCase) :
	def test_crud(self) :
		performanceReview = PerformanceReview()
		performanceReview.reviewNumber = "default reviewNumber field value"
		performanceReview.reviewDate = datetime.datetime.now()
		performanceReview.reviewerComments = "default reviewerComments field value"
		performanceReview.rating = "default rating field value"
		performanceReview.status = "default status field value"
		
		delegate = PerformanceReviewDelegate()
		responseObj = delegate.create(performanceReview)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


