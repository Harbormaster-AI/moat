from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.ExperimentVariant import ExperimentVariant
from advertisingOnDjango.models.Experiment import Experiment
from advertisingOnDjango.models.CreativeVariation import CreativeVariation
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ExperimentVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentVariantDelegate Declaration
#======================================================================
class ExperimentVariantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, experimentVariantId ):
		try:	
			experimentVariant = ExperimentVariant.objects.filter(id=experimentVariantId)
			return experimentVariant.first();
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError("ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, experimentVariant):
		for model in serializers.deserialize("json", experimentVariant):
			model.save()
			return model;

	def create(self, experimentVariant):
		experimentVariant.save()
		return experimentVariant;

	def saveFromJson(self, experimentVariant):
		for model in serializers.deserialize("json", experimentVariant):
			model.save()
			return experimentVariant;
	
	def save(self, experimentVariant):
		experimentVariant.save()
		return experimentVariant;
	
	def delete(self, experimentVariantId ):
		errMsg = "Failed to delete ExperimentVariant from db using id " + str(experimentVariantId)
		
		try:
			experimentVariant = ExperimentVariant.objects.get(id=experimentVariantId)
			experimentVariant.delete()
			return True
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError("ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ExperimentVariant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ExperimentVariant from db")
		except Exception:
			return None;
		
	def assignExperiment( self, experimentVariantId, experimentId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to assign element " + str(experimentId) + " for Experiment on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# get the Experiment from db
			experiment = ExperimentDelegate().get(experimentId).first();
			
			# assign the Experiment		
			experimentVariant.experiment = experiment
			
			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignExperiment( self, experimentVariantId ):
		errMsg = "Failed to unassign element " + str(experimentId) + " for Experiment on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# assign to None for unassignment
			experimentVariant.experiment = None			

			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCreativeVariation( self, experimentVariantId, creativeVariationId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeVariationDelegate import CreativeVariationDelegate

		errMsg = "Failed to assign element " + str(creativeVariationId) + " for CreativeVariation on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# get the CreativeVariation from db
			creativeVariation = CreativeVariationDelegate().get(creativeVariationId).first();
			
			# assign the CreativeVariation		
			experimentVariant.creativeVariation = creativeVariation
			
			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except CreativeVariation.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeVariation with id " + str(creativeVariationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreativeVariation( self, experimentVariantId ):
		errMsg = "Failed to unassign element " + str(creativeVariationId) + " for CreativeVariation on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# assign to None for unassignment
			experimentVariant.creativeVariation = None			

			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineItem( self, experimentVariantId, lineItemId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to assign element " + str(lineItemId) + " for LineItem on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# get the LineItem from db
			lineItem = LineItemDelegate().get(lineItemId).first();
			
			# assign the LineItem		
			experimentVariant.lineItem = lineItem
			
			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineItem( self, experimentVariantId ):
		errMsg = "Failed to unassign element " + str(lineItemId) + " for LineItem on ExperimentVariant"

		try:
			# get the ExperimentVariant from db
			experimentVariant = self.get( experimentVariantId ).first()	
			
			# assign to None for unassignment
			experimentVariant.lineItem = None			

			#save it
			experimentVariant.save()

			# reload and return the appropriate version					
			return self.get( experimentVariantId );
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant with id " + str(experimentVariantId) + " does not exist.")
		except Exception:
			return None;
		
