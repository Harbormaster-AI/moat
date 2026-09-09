from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Experiment import Experiment
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.ExperimentVariant import ExperimentVariant
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Experiment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentDelegate Declaration
#======================================================================
class ExperimentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, experimentId ):
		try:	
			experiment = Experiment.objects.filter(id=experimentId)
			return experiment.first();
		except Experiment.DoesNotExist:
			raise ProcessingError("Experiment with id " + str(experimentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, experiment):
		for model in serializers.deserialize("json", experiment):
			model.save()
			return model;

	def create(self, experiment):
		experiment.save()
		return experiment;

	def saveFromJson(self, experiment):
		for model in serializers.deserialize("json", experiment):
			model.save()
			return experiment;
	
	def save(self, experiment):
		experiment.save()
		return experiment;
	
	def delete(self, experimentId ):
		errMsg = "Failed to delete Experiment from db using id " + str(experimentId)
		
		try:
			experiment = Experiment.objects.get(id=experimentId)
			experiment.delete()
			return True
		except Experiment.DoesNotExist:
			raise ProcessingError("Experiment with id " + str(experimentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Experiment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Experiment from db")
		except Exception:
			return None;
		
	def assignCampaign( self, experimentId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on Experiment"

		try:
			# get the Experiment from db
			experiment = self.get( experimentId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			experiment.campaign = campaign
			
			#save it
			experiment.save()

			# reload and return the appropriate version					
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, experimentId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on Experiment"

		try:
			# get the Experiment from db
			experiment = self.get( experimentId ).first()	
			
			# assign to None for unassignment
			experiment.campaign = None			

			#save it
			experiment.save()

			# reload and return the appropriate version					
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Exception:
			return None;
		
	def addVariants( self, experimentId, variantsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ExperimentVariantDelegate import ExperimentVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ExperimentVariant		
				experimentVariant = ExperimentVariantDelegate().get(id).first();	
				# add the ExperimentVariant
				experiment.variants.add(experimentVariant)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, experimentId, variantsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ExperimentVariantDelegate import ExperimentVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on Experiment"

		try:
			# get the Experiment
			experiment = self.get( experimentId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ExperimentVariant		
				experimentVariant = ExperimentVariantDelegate().get(id).first();	
				# add the ExperimentVariant
				experiment.variants.remove(experimentVariant)
				
			# save it		
			experiment.save()
			
			# reload and return the appropriate version
			return self.get( experimentId );
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except ExperimentVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ExperimentVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
