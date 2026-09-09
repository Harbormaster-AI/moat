import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.MROFacility import MROFacility
from aerospaceOnDjango.delegates.MROFacilityDelegate import MROFacilityDelegate

 #======================================================================
# 
# Encapsulates data for model MROFacility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MROFacilityTest Declaration
#======================================================================
class MROFacilityTest (TestCase) :
	def test_crud(self) :
		mROFacility = MROFacility()
		mROFacility.name = "default name field value"
		mROFacility.approvalScope = "default approvalScope field value"
		mROFacility.address = "default address field value"
		
		delegate = MROFacilityDelegate()
		responseObj = delegate.create(mROFacility)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


