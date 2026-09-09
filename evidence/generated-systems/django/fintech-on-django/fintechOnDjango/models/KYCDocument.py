from django.db import models
from fintechOnDjango.models.KYCDocumentType import KYCDocumentType
from fintechOnDjango.models.DocumentStatus import DocumentStatus

#======================================================================
# 
# Encapsulates data for model KYCDocument
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCDocument Declaration
#======================================================================
class KYCDocument (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reference = DocumentReference
	issuedCountry = models.CharField(max_length=200, null=True)
	expirationDate = models.DateField(null=True)
	kycProfile = models.ForeignKey('KYCProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	documentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in KYCDocumentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reference
		str = str + self.issuedCountry
		str = str + self.expirationDate
		str = str + self.documentType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "KYCDocument";
    
	def objectType(self):
		return "KYCDocument";
