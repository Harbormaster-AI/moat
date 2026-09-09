import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Review import Review
from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

 #======================================================================
# 
# Encapsulates data for model Review
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReviewTest Declaration
#======================================================================
class ReviewTest (TestCase) :
	def test_crud(self) :
		review = Review()
		review.rating = 22
		review.title = "default title field value"
		review.content = "default content field value"
		review.createdAt = datetime.datetime.now()
		review.status = "default status field value"
		
		delegate = ReviewDelegate()
		responseObj = delegate.create(review)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


