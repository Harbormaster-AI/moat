from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.models.EvaluationMetric import EvaluationMetric
from analyticsOnDjango.models.InferenceEndpoint import InferenceEndpoint
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ModelVersion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelVersionDelegate Declaration
#======================================================================
class ModelVersionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, modelVersionId ):
		try:	
			modelVersion = ModelVersion.objects.filter(id=modelVersionId)
			return modelVersion.first();
		except ModelVersion.DoesNotExist:
			raise ProcessingError("ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, modelVersion):
		for model in serializers.deserialize("json", modelVersion):
			model.save()
			return model;

	def create(self, modelVersion):
		modelVersion.save()
		return modelVersion;

	def saveFromJson(self, modelVersion):
		for model in serializers.deserialize("json", modelVersion):
			model.save()
			return modelVersion;
	
	def save(self, modelVersion):
		modelVersion.save()
		return modelVersion;
	
	def delete(self, modelVersionId ):
		errMsg = "Failed to delete ModelVersion from db using id " + str(modelVersionId)
		
		try:
			modelVersion = ModelVersion.objects.get(id=modelVersionId)
			modelVersion.delete()
			return True
		except ModelVersion.DoesNotExist:
			raise ProcessingError("ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ModelVersion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ModelVersion from db")
		except Exception:
			return None;
		
	def assignModel( self, modelVersionId, modelId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to assign element " + str(modelId) + " for Model on ModelVersion"

		try:
			# get the ModelVersion from db
			modelVersion = self.get( modelVersionId ).first()	
			
			# get the Model from db
			model = ModelDelegate().get(modelId).first();
			
			# assign the Model		
			modelVersion.model = model
			
			#save it
			modelVersion.save()

			# reload and return the appropriate version					
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model with id " + str(modelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModel( self, modelVersionId ):
		errMsg = "Failed to unassign element " + str(modelId) + " for Model on ModelVersion"

		try:
			# get the ModelVersion from db
			modelVersion = self.get( modelVersionId ).first()	
			
			# assign to None for unassignment
			modelVersion.model = None			

			#save it
			modelVersion.save()

			# reload and return the appropriate version					
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTrainingRun( self, modelVersionId, trainingRunId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to assign element " + str(trainingRunId) + " for TrainingRun on ModelVersion"

		try:
			# get the ModelVersion from db
			modelVersion = self.get( modelVersionId ).first()	
			
			# get the TrainingRun from db
			trainingRun = TrainingRunDelegate().get(trainingRunId).first();
			
			# assign the TrainingRun		
			modelVersion.trainingRun = trainingRun
			
			#save it
			modelVersion.save()

			# reload and return the appropriate version					
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTrainingRun( self, modelVersionId ):
		errMsg = "Failed to unassign element " + str(trainingRunId) + " for TrainingRun on ModelVersion"

		try:
			# get the ModelVersion from db
			modelVersion = self.get( modelVersionId ).first()	
			
			# assign to None for unassignment
			modelVersion.trainingRun = None			

			#save it
			modelVersion.save()

			# reload and return the appropriate version					
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
		
	def addEvaluationMetrics( self, modelVersionId, evaluationMetricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.EvaluationMetricDelegate import EvaluationMetricDelegate

		errMsg = "Failed to add elements " + str(evaluationMetricsIds) + " for EvaluationMetrics on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = evaluationMetricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EvaluationMetric		
				evaluationMetric = EvaluationMetricDelegate().get(id).first();	
				# add the EvaluationMetric
				modelVersion.evaluationMetrics.add(evaluationMetric)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEvaluationMetrics( self, modelVersionId, evaluationMetricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.EvaluationMetricDelegate import EvaluationMetricDelegate

		errMsg = "Failed to remove elements " + str(evaluationMetricsIds) + " for EvaluationMetrics on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = evaluationMetricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EvaluationMetric		
				evaluationMetric = EvaluationMetricDelegate().get(id).first();	
				# add the EvaluationMetric
				modelVersion.evaluationMetrics.remove(evaluationMetric)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDeployments( self, modelVersionId, deploymentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.InferenceEndpointDelegate import InferenceEndpointDelegate

		errMsg = "Failed to add elements " + str(deploymentsIds) + " for Deployments on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = deploymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InferenceEndpoint		
				inferenceEndpoint = InferenceEndpointDelegate().get(id).first();	
				# add the InferenceEndpoint
				modelVersion.deployments.add(inferenceEndpoint)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDeployments( self, modelVersionId, deploymentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.InferenceEndpointDelegate import InferenceEndpointDelegate

		errMsg = "Failed to remove elements " + str(deploymentsIds) + " for Deployments on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = deploymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InferenceEndpoint		
				inferenceEndpoint = InferenceEndpointDelegate().get(id).first();	
				# add the InferenceEndpoint
				modelVersion.deployments.remove(inferenceEndpoint)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except InferenceEndpoint.DoesNotExist:
			raise ProcessingError(errMsg + " : InferenceEndpoint does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatureSets( self, modelVersionId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to add elements " + str(featureSetsIds) + " for FeatureSets on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				modelVersion.featureSets.add(featureSet)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatureSets( self, modelVersionId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to remove elements " + str(featureSetsIds) + " for FeatureSets on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				modelVersion.featureSets.remove(featureSet)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, modelVersionId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				modelVersion.datasets.add(dataSet)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, modelVersionId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on ModelVersion"

		try:
			# get the ModelVersion
			modelVersion = self.get( modelVersionId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				modelVersion.datasets.remove(dataSet)
				
			# save it		
			modelVersion.save()
			
			# reload and return the appropriate version
			return self.get( modelVersionId );
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
