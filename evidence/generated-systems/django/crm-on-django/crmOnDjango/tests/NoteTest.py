import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Note import Note
from crmOnDjango.delegates.NoteDelegate import NoteDelegate

 #======================================================================
# 
# Encapsulates data for model Note
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NoteTest Declaration
#======================================================================
class NoteTest (TestCase) :
	def test_crud(self) :
		note = Note()
		note.title = "default title field value"
		note.content = "default content field value"
		note.createdAt = "default createdAt field value"
		note.updatedAt = "default updatedAt field value"
		
		delegate = NoteDelegate()
		responseObj = delegate.create(note)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


