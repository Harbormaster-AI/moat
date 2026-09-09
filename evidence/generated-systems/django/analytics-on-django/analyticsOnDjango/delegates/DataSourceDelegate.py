from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.DataSource import DataSource
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSourceDelegate Declaration
#======================================================================
class DataSourceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataSourceId ):
		try:	
			dataSource = DataSource.objects.filter(id=dataSourceId)
			return dataSource.first();
		except DataSource.DoesNotExist:
			raise ProcessingError("DataSource with id " + str(dataSourceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataSource):
		for model in serializers.deserialize("json", dataSource):
			model.save()
			return model;

	def create(self, dataSource):
		dataSource.save()
		return dataSource;

	def saveFromJson(self, dataSource):
		for model in serializers.deserialize("json", dataSource):
			model.save()
			return dataSource;
	
	def save(self, dataSource):
		dataSource.save()
		return dataSource;
	
	def delete(self, dataSourceId ):
		errMsg = "Failed to delete DataSource from db using id " + str(dataSourceId)
		
		try:
			dataSource = DataSource.objects.get(id=dataSourceId)
			dataSource.delete()
			return True
		except DataSource.DoesNotExist:
			raise ProcessingError("DataSource with id " + str(dataSourceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataSource.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataSource from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, dataSourceId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on DataSource"

		try:
			# get the DataSource from db
			dataSource = self.get( dataSourceId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			dataSource.workspace = analyticsWorkspace
			
			#save it
			dataSource.save()

			# reload and return the appropriate version					
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, dataSourceId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on DataSource"

		try:
			# get the DataSource from db
			dataSource = self.get( dataSourceId ).first()	
			
			# assign to None for unassignment
			dataSource.analyticsWorkspace = None			

			#save it
			dataSource.save()

			# reload and return the appropriate version					
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except Exception:
			return None;
		
	def addProducedDatasets( self, dataSourceId, producedDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(producedDatasetsIds) + " for ProducedDatasets on DataSource"

		try:
			# get the DataSource
			dataSource = self.get( dataSourceId ).first()
				
			# split on a comma with no spaces
			idList = producedDatasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataSource.producedDatasets.add(dataSet)
				
			# save it		
			dataSource.save()
			
			# reload and return the appropriate version
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducedDatasets( self, dataSourceId, producedDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(producedDatasetsIds) + " for ProducedDatasets on DataSource"

		try:
			# get the DataSource
			dataSource = self.get( dataSourceId ).first()
				
			# split on a comma with no spaces
			idList = producedDatasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataSource.producedDatasets.remove(dataSet)
				
			# save it		
			dataSource.save()
			
			# reload and return the appropriate version
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPipelines( self, dataSourceId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to add elements " + str(pipelinesIds) + " for Pipelines on DataSource"

		try:
			# get the DataSource
			dataSource = self.get( dataSourceId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				dataSource.pipelines.add(dataPipeline)
				
			# save it		
			dataSource.save()
			
			# reload and return the appropriate version
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePipelines( self, dataSourceId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to remove elements " + str(pipelinesIds) + " for Pipelines on DataSource"

		try:
			# get the DataSource
			dataSource = self.get( dataSourceId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				dataSource.pipelines.remove(dataPipeline)
				
			# save it		
			dataSource.save()
			
			# reload and return the appropriate version
			return self.get( dataSourceId );
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource with id " + str(dataSourceId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
