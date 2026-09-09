import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

 #======================================================================
# 
# Encapsulates data for model Candidate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CandidateTest Declaration
#======================================================================
class CandidateTest (TestCase) :
	def test_crud(self) :
		candidate = Candidate()
		candidate.name = "default name field value"
		candidate.email = "default email field value"
		candidate.phone = "default phone field value"
		candidate.source = "default source field value"
		
		delegate = CandidateDelegate()
		responseObj = delegate.create(candidate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


