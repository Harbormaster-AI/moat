from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.models.Visualization import Visualization
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Metric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MetricDelegate Declaration
#======================================================================
class MetricDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, metricId ):
		try:	
			metric = Metric.objects.filter(id=metricId)
			return metric.first();
		except Metric.DoesNotExist:
			raise ProcessingError("Metric with id " + str(metricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, metric):
		for model in serializers.deserialize("json", metric):
			model.save()
			return model;

	def create(self, metric):
		metric.save()
		return metric;

	def saveFromJson(self, metric):
		for model in serializers.deserialize("json", metric):
			model.save()
			return metric;
	
	def save(self, metric):
		metric.save()
		return metric;
	
	def delete(self, metricId ):
		errMsg = "Failed to delete Metric from db using id " + str(metricId)
		
		try:
			metric = Metric.objects.get(id=metricId)
			metric.delete()
			return True
		except Metric.DoesNotExist:
			raise ProcessingError("Metric with id " + str(metricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Metric.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Metric from db")
		except Exception:
			return None;
		
	def assignSemanticModel( self, metricId, semanticModelId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to assign element " + str(semanticModelId) + " for SemanticModel on Metric"

		try:
			# get the Metric from db
			metric = self.get( metricId ).first()	
			
			# get the SemanticModel from db
			semanticModel = SemanticModelDelegate().get(semanticModelId).first();
			
			# assign the SemanticModel		
			metric.semanticModel = semanticModel
			
			#save it
			metric.save()

			# reload and return the appropriate version					
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSemanticModel( self, metricId ):
		errMsg = "Failed to unassign element " + str(semanticModelId) + " for SemanticModel on Metric"

		try:
			# get the Metric from db
			metric = self.get( metricId ).first()	
			
			# assign to None for unassignment
			metric.semanticModel = None			

			#save it
			metric.save()

			# reload and return the appropriate version					
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, metricId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				metric.datasets.add(dataSet)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, metricId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				metric.datasets.remove(dataSet)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGlossaryTerms( self, metricId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to add elements " + str(glossaryTermsIds) + " for GlossaryTerms on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				metric.glossaryTerms.add(businessGlossaryTerm)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGlossaryTerms( self, metricId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to remove elements " + str(glossaryTermsIds) + " for GlossaryTerms on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				metric.glossaryTerms.remove(businessGlossaryTerm)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAlerts( self, metricId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				metric.alerts.add(alert)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, metricId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				metric.alerts.remove(alert)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addVisualizations( self, metricId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to add elements " + str(visualizationsIds) + " for Visualizations on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				metric.visualizations.add(visualization)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVisualizations( self, metricId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to remove elements " + str(visualizationsIds) + " for Visualizations on Metric"

		try:
			# get the Metric
			metric = self.get( metricId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				metric.visualizations.remove(visualization)
				
			# save it		
			metric.save()
			
			# reload and return the appropriate version
			return self.get( metricId );
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
