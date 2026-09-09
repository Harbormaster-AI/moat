from django.db import models

#======================================================================
# 
# Encapsulates data for model TypeCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TypeCertificate Declaration
#======================================================================
class TypeCertificate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	certificateNumber = models.CharField(max_length=200, null=True)
	authority = models.CharField(max_length=200, null=True)
	program = models.ForeignKey('AircraftProgram', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.certificateNumber
		str = str + self.authority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TypeCertificate";
    
	def objectType(self):
		return "TypeCertificate";
