from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.Feature import Feature
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FeatureSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureSetDelegate Declaration
#======================================================================
class FeatureSetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, featureSetId ):
		try:	
			featureSet = FeatureSet.objects.filter(id=featureSetId)
			return featureSet.first();
		except FeatureSet.DoesNotExist:
			raise ProcessingError("FeatureSet with id " + str(featureSetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, featureSet):
		for model in serializers.deserialize("json", featureSet):
			model.save()
			return model;

	def create(self, featureSet):
		featureSet.save()
		return featureSet;

	def saveFromJson(self, featureSet):
		for model in serializers.deserialize("json", featureSet):
			model.save()
			return featureSet;
	
	def save(self, featureSet):
		featureSet.save()
		return featureSet;
	
	def delete(self, featureSetId ):
		errMsg = "Failed to delete FeatureSet from db using id " + str(featureSetId)
		
		try:
			featureSet = FeatureSet.objects.get(id=featureSetId)
			featureSet.delete()
			return True
		except FeatureSet.DoesNotExist:
			raise ProcessingError("FeatureSet with id " + str(featureSetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FeatureSet.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FeatureSet from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, featureSetId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on FeatureSet"

		try:
			# get the FeatureSet from db
			featureSet = self.get( featureSetId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			featureSet.workspace = analyticsWorkspace
			
			#save it
			featureSet.save()

			# reload and return the appropriate version					
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, featureSetId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on FeatureSet"

		try:
			# get the FeatureSet from db
			featureSet = self.get( featureSetId ).first()	
			
			# assign to None for unassignment
			featureSet.analyticsWorkspace = None			

			#save it
			featureSet.save()

			# reload and return the appropriate version					
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Exception:
			return None;
		
	def addFeatures( self, featureSetId, featuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

		errMsg = "Failed to add elements " + str(featuresIds) + " for Features on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = featuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Feature		
				feature = FeatureDelegate().get(id).first();	
				# add the Feature
				featureSet.features.add(feature)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatures( self, featureSetId, featuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

		errMsg = "Failed to remove elements " + str(featuresIds) + " for Features on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = featuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Feature		
				feature = FeatureDelegate().get(id).first();	
				# add the Feature
				featureSet.features.remove(feature)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, featureSetId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				featureSet.datasets.add(dataSet)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, featureSetId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				featureSet.datasets.remove(dataSet)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, featureSetId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				featureSet.models.add(model)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, featureSetId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				featureSet.models.remove(model)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModelVersions( self, featureSetId, modelVersionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to add elements " + str(modelVersionsIds) + " for ModelVersions on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = modelVersionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				featureSet.modelVersions.add(modelVersion)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModelVersions( self, featureSetId, modelVersionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to remove elements " + str(modelVersionsIds) + " for ModelVersions on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = modelVersionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				featureSet.modelVersions.remove(modelVersion)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTags( self, featureSetId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to add elements " + str(tagsIds) + " for Tags on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				featureSet.tags.add(tag)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTags( self, featureSetId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to remove elements " + str(tagsIds) + " for Tags on FeatureSet"

		try:
			# get the FeatureSet
			featureSet = self.get( featureSetId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				featureSet.tags.remove(tag)
				
			# save it		
			featureSet.save()
			
			# reload and return the appropriate version
			return self.get( featureSetId );
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
