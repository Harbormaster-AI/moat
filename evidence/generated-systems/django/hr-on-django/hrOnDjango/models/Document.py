from django.db import models
from hrOnDjango.models.DocumentType import DocumentType

#======================================================================
# 
# Encapsulates data for model Document
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Document Declaration
#======================================================================
class Document (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	fileUrl = models.CharField(max_length=200, null=True)
	uploadedDate = models.DateField(null=True)
	candidate = models.ForeignKey('Candidate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	documentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.fileUrl
		str = str + self.uploadedDate
		str = str + self.documentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Document";
    
	def objectType(self):
		return "Document";
