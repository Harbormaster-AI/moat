from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Feature import Feature
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Feature
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureDelegate Declaration
#======================================================================
class FeatureDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, featureId ):
		try:	
			feature = Feature.objects.filter(id=featureId)
			return feature.first();
		except Feature.DoesNotExist:
			raise ProcessingError("Feature with id " + str(featureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, feature):
		for model in serializers.deserialize("json", feature):
			model.save()
			return model;

	def create(self, feature):
		feature.save()
		return feature;

	def saveFromJson(self, feature):
		for model in serializers.deserialize("json", feature):
			model.save()
			return feature;
	
	def save(self, feature):
		feature.save()
		return feature;
	
	def delete(self, featureId ):
		errMsg = "Failed to delete Feature from db using id " + str(featureId)
		
		try:
			feature = Feature.objects.get(id=featureId)
			feature.delete()
			return True
		except Feature.DoesNotExist:
			raise ProcessingError("Feature with id " + str(featureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Feature.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Feature from db")
		except Exception:
			return None;
		
	def assignFeatureSet( self, featureId, featureSetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to assign element " + str(featureSetId) + " for FeatureSet on Feature"

		try:
			# get the Feature from db
			feature = self.get( featureId ).first()	
			
			# get the FeatureSet from db
			featureSet = FeatureSetDelegate().get(featureSetId).first();
			
			# assign the FeatureSet		
			feature.featureSet = featureSet
			
			#save it
			feature.save()

			# reload and return the appropriate version					
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet with id " + str(featureSetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFeatureSet( self, featureId ):
		errMsg = "Failed to unassign element " + str(featureSetId) + " for FeatureSet on Feature"

		try:
			# get the Feature from db
			feature = self.get( featureId ).first()	
			
			# assign to None for unassignment
			feature.featureSet = None			

			#save it
			feature.save()

			# reload and return the appropriate version					
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except Exception:
			return None;
		
	def addSourceDatasets( self, featureId, sourceDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(sourceDatasetsIds) + " for SourceDatasets on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = sourceDatasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				feature.sourceDatasets.add(dataSet)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSourceDatasets( self, featureId, sourceDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(sourceDatasetsIds) + " for SourceDatasets on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = sourceDatasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				feature.sourceDatasets.remove(dataSet)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, featureId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				feature.models.add(model)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, featureId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				feature.models.remove(model)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrainingRuns( self, featureId, trainingRunsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to add elements " + str(trainingRunsIds) + " for TrainingRuns on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = trainingRunsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingRun		
				trainingRun = TrainingRunDelegate().get(id).first();	
				# add the TrainingRun
				feature.trainingRuns.add(trainingRun)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrainingRuns( self, featureId, trainingRunsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to remove elements " + str(trainingRunsIds) + " for TrainingRuns on Feature"

		try:
			# get the Feature
			feature = self.get( featureId ).first()
				
			# split on a comma with no spaces
			idList = trainingRunsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingRun		
				trainingRun = TrainingRunDelegate().get(id).first();	
				# add the TrainingRun
				feature.trainingRuns.remove(trainingRun)
				
			# save it		
			feature.save()
			
			# reload and return the appropriate version
			return self.get( featureId );
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature with id " + str(featureId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
