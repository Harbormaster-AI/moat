import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.DispositionReview import DispositionReview
from governanceOnDjango.delegates.DispositionReviewDelegate import DispositionReviewDelegate

 #======================================================================
# 
# Encapsulates data for model DispositionReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionReviewTest Declaration
#======================================================================
class DispositionReviewTest (TestCase) :
	def test_crud(self) :
		dispositionReview = DispositionReview()
		dispositionReview.reviewDate = datetime.datetime.now()
		dispositionReview.reviewer = "default reviewer field value"
		dispositionReview.notes = "default notes field value"
		dispositionReview.outcome = "default outcome field value"
		
		delegate = DispositionReviewDelegate()
		responseObj = delegate.create(dispositionReview)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


