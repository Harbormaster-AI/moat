import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.AuditWorkpaper import AuditWorkpaper
from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

 #======================================================================
# 
# Encapsulates data for model AuditWorkpaper
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditWorkpaperTest Declaration
#======================================================================
class AuditWorkpaperTest (TestCase) :
	def test_crud(self) :
		auditWorkpaper = AuditWorkpaper()
		auditWorkpaper.workpaperRef = "default workpaperRef field value"
		auditWorkpaper.subject = "default subject field value"
		auditWorkpaper.workpaperUrl = "default workpaperUrl field value"
		
		delegate = AuditWorkpaperDelegate()
		responseObj = delegate.create(auditWorkpaper)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


