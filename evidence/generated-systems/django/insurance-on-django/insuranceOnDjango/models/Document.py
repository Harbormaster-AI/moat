from django.db import models
from insuranceOnDjango.models.DocumentType import DocumentType

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
	fileName = models.CharField(max_length=200, null=True)
	uploadedDate = models.DateField(null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	application = models.ForeignKey('Application', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	documentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.fileName
		str = str + self.uploadedDate
		str = str + self.documentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Document";
    
	def objectType(self):
		return "Document";
