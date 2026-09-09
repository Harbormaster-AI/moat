from django.db import models

#======================================================================
# 
# Encapsulates data for model CreativeFile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeFile Declaration
#======================================================================
class CreativeFile (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	uri = URL
	fileSizeKB = models.IntegerField(null=True)
	mimeType = models.CharField(max_length=200, null=True)
	checksum = models.CharField(max_length=200, null=True)
	creativeAsset = models.ForeignKey('CreativeAsset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.uri
		str = str + self.fileSizeKB
		str = str + self.mimeType
		str = str + self.checksum
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CreativeFile";
    
	def objectType(self):
		return "CreativeFile";
