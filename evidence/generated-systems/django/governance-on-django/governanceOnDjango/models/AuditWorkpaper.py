from django.db import models

#======================================================================
# 
# Encapsulates data for model AuditWorkpaper
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditWorkpaper Declaration
#======================================================================
class AuditWorkpaper (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	workpaperRef = models.CharField(max_length=200, null=True)
	subject = models.CharField(max_length=200, null=True)
	workpaperUrl = URL
	engagement = models.ForeignKey('AuditEngagement', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	evidence = models.ManyToManyField('Evidence',  blank=True, related_name='+')
	findings = models.ManyToManyField('AuditFinding',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.workpaperRef
		str = str + self.subject
		str = str + self.workpaperUrl
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AuditWorkpaper";
    
	def objectType(self):
		return "AuditWorkpaper";
