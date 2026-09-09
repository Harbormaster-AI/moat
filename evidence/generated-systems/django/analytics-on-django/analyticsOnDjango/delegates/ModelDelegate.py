from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Model
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelDelegate Declaration
#======================================================================
class ModelDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, modelId ):
		try:	
			model = Model.objects.filter(id=modelId)
			return model.first();
		except Model.DoesNotExist:
			raise ProcessingError("Model with id " + str(modelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, model):
		for model in serializers.deserialize("json", model):
			model.save()
			return model;

	def create(self, model):
		model.save()
		return model;

	def saveFromJson(self, model):
		for model in serializers.deserialize("json", model):
			model.save()
			return model;
	
	def save(self, model):
		model.save()
		return model;
	
	def delete(self, modelId ):
		errMsg = "Failed to delete Model from db using id " + str(modelId)
		
		try:
			model = Model.objects.get(id=modelId)
			model.delete()
			return True
		except Model.DoesNotExist:
			raise ProcessingError("Model with id " + str(modelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Model.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Model from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, modelId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on Model"

		try:
			# get the Model from db
			model = self.get( modelId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			model.workspace = analyticsWorkspace
			
			#save it
			model.save()

			# reload and return the appropriate version					
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, modelId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on Model"

		try:
			# get the Model from db
			model = self.get( modelId ).first()	
			
			# assign to None for unassignment
			model.analyticsWorkspace = None			

			#save it
			model.save()

			# reload and return the appropriate version					
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Exception:
			return None;
		
	def addVersions( self, modelId, versionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to add elements " + str(versionsIds) + " for Versions on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = versionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				model.versions.add(modelVersion)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVersions( self, modelId, versionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to remove elements " + str(versionsIds) + " for Versions on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = versionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				model.versions.remove(modelVersion)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatureSets( self, modelId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to add elements " + str(featureSetsIds) + " for FeatureSets on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				model.featureSets.add(featureSet)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatureSets( self, modelId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to remove elements " + str(featureSetsIds) + " for FeatureSets on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				model.featureSets.remove(featureSet)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExperiments( self, modelId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to add elements " + str(experimentsIds) + " for Experiments on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				model.experiments.add(experiment)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExperiments( self, modelId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to remove elements " + str(experimentsIds) + " for Experiments on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				model.experiments.remove(experiment)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTags( self, modelId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to add elements " + str(tagsIds) + " for Tags on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				model.tags.add(tag)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTags( self, modelId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to remove elements " + str(tagsIds) + " for Tags on Model"

		try:
			# get the Model
			model = self.get( modelId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				model.tags.remove(tag)
				
			# save it		
			model.save()
			
			# reload and return the appropriate version
			return self.get( modelId );
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
