from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.ContentCategory import ContentCategory
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ContentCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContentCategoryDelegate Declaration
#======================================================================
class ContentCategoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, contentCategoryId ):
		try:	
			contentCategory = ContentCategory.objects.filter(id=contentCategoryId)
			return contentCategory.first();
		except ContentCategory.DoesNotExist:
			raise ProcessingError("ContentCategory with id " + str(contentCategoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, contentCategory):
		for model in serializers.deserialize("json", contentCategory):
			model.save()
			return model;

	def create(self, contentCategory):
		contentCategory.save()
		return contentCategory;

	def saveFromJson(self, contentCategory):
		for model in serializers.deserialize("json", contentCategory):
			model.save()
			return contentCategory;
	
	def save(self, contentCategory):
		contentCategory.save()
		return contentCategory;
	
	def delete(self, contentCategoryId ):
		errMsg = "Failed to delete ContentCategory from db using id " + str(contentCategoryId)
		
		try:
			contentCategory = ContentCategory.objects.get(id=contentCategoryId)
			contentCategory.delete()
			return True
		except ContentCategory.DoesNotExist:
			raise ProcessingError("ContentCategory with id " + str(contentCategoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ContentCategory.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ContentCategory from db")
		except Exception:
			return None;
		
