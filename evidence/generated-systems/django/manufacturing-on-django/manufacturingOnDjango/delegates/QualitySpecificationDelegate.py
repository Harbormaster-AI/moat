from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.QualitySpecification import QualitySpecification
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model QualitySpecification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualitySpecificationDelegate Declaration
#======================================================================
class QualitySpecificationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, qualitySpecificationId ):
		try:	
			qualitySpecification = QualitySpecification.objects.filter(id=qualitySpecificationId)
			return qualitySpecification.first();
		except QualitySpecification.DoesNotExist:
			raise ProcessingError("QualitySpecification with id " + str(qualitySpecificationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, qualitySpecification):
		for model in serializers.deserialize("json", qualitySpecification):
			model.save()
			return model;

	def create(self, qualitySpecification):
		qualitySpecification.save()
		return qualitySpecification;

	def saveFromJson(self, qualitySpecification):
		for model in serializers.deserialize("json", qualitySpecification):
			model.save()
			return qualitySpecification;
	
	def save(self, qualitySpecification):
		qualitySpecification.save()
		return qualitySpecification;
	
	def delete(self, qualitySpecificationId ):
		errMsg = "Failed to delete QualitySpecification from db using id " + str(qualitySpecificationId)
		
		try:
			qualitySpecification = QualitySpecification.objects.get(id=qualitySpecificationId)
			qualitySpecification.delete()
			return True
		except QualitySpecification.DoesNotExist:
			raise ProcessingError("QualitySpecification with id " + str(qualitySpecificationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = QualitySpecification.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all QualitySpecification from db")
		except Exception:
			return None;
		
	def assignItem( self, qualitySpecificationId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on QualitySpecification"

		try:
			# get the QualitySpecification from db
			qualitySpecification = self.get( qualitySpecificationId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			qualitySpecification.item = item
			
			#save it
			qualitySpecification.save()

			# reload and return the appropriate version					
			return self.get( qualitySpecificationId );
		except QualitySpecification.DoesNotExist:
			raise ProcessingError(errMsg + " : QualitySpecification with id " + str(qualitySpecificationId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, qualitySpecificationId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on QualitySpecification"

		try:
			# get the QualitySpecification from db
			qualitySpecification = self.get( qualitySpecificationId ).first()	
			
			# assign to None for unassignment
			qualitySpecification.item = None			

			#save it
			qualitySpecification.save()

			# reload and return the appropriate version					
			return self.get( qualitySpecificationId );
		except QualitySpecification.DoesNotExist:
			raise ProcessingError(errMsg + " : QualitySpecification with id " + str(qualitySpecificationId) + " does not exist.")
		except Exception:
			return None;
		
