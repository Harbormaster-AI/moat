from django.db import models

#======================================================================
# 
# Encapsulates data for model Certification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Certification Declaration
#======================================================================
class Certification (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	issuer = models.CharField(max_length=200, null=True)
	validFrom = models.DateField(null=True)
	validTo = models.DateField(null=True)
	credentialId = models.CharField(max_length=200, null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	course = models.ForeignKey('TrainingCourse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.issuer
		str = str + self.validFrom
		str = str + self.validTo
		str = str + self.credentialId
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Certification";
    
	def objectType(self):
		return "Certification";
