from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSource import DataSource
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.models.Measure import Measure
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.QualityRule import QualityRule
from analyticsOnDjango.models.LineageNode import LineageNode
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSetDelegate Declaration
#======================================================================
class DataSetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataSetId ):
		try:	
			dataSet = DataSet.objects.filter(id=dataSetId)
			return dataSet.first();
		except DataSet.DoesNotExist:
			raise ProcessingError("DataSet with id " + str(dataSetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataSet):
		for model in serializers.deserialize("json", dataSet):
			model.save()
			return model;

	def create(self, dataSet):
		dataSet.save()
		return dataSet;

	def saveFromJson(self, dataSet):
		for model in serializers.deserialize("json", dataSet):
			model.save()
			return dataSet;
	
	def save(self, dataSet):
		dataSet.save()
		return dataSet;
	
	def delete(self, dataSetId ):
		errMsg = "Failed to delete DataSet from db using id " + str(dataSetId)
		
		try:
			dataSet = DataSet.objects.get(id=dataSetId)
			dataSet.delete()
			return True
		except DataSet.DoesNotExist:
			raise ProcessingError("DataSet with id " + str(dataSetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataSet.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataSet from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, dataSetId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on DataSet"

		try:
			# get the DataSet from db
			dataSet = self.get( dataSetId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			dataSet.workspace = analyticsWorkspace
			
			#save it
			dataSet.save()

			# reload and return the appropriate version					
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, dataSetId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on DataSet"

		try:
			# get the DataSet from db
			dataSet = self.get( dataSetId ).first()	
			
			# assign to None for unassignment
			dataSet.analyticsWorkspace = None			

			#save it
			dataSet.save()

			# reload and return the appropriate version					
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineageNode( self, dataSetId, lineageNodeId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to assign element " + str(lineageNodeId) + " for LineageNode on DataSet"

		try:
			# get the DataSet from db
			dataSet = self.get( dataSetId ).first()	
			
			# get the LineageNode from db
			lineageNode = LineageNodeDelegate().get(lineageNodeId).first();
			
			# assign the LineageNode		
			dataSet.lineageNode = lineageNode
			
			#save it
			dataSet.save()

			# reload and return the appropriate version					
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineageNode( self, dataSetId ):
		errMsg = "Failed to unassign element " + str(lineageNodeId) + " for LineageNode on DataSet"

		try:
			# get the DataSet from db
			dataSet = self.get( dataSetId ).first()	
			
			# assign to None for unassignment
			dataSet.lineageNode = None			

			#save it
			dataSet.save()

			# reload and return the appropriate version					
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Exception:
			return None;
		
	def addSources( self, dataSetId, sourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to add elements " + str(sourcesIds) + " for Sources on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = sourcesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				dataSet.sources.add(dataSource)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSources( self, dataSetId, sourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to remove elements " + str(sourcesIds) + " for Sources on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = sourcesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				dataSet.sources.remove(dataSource)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPipelines( self, dataSetId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to add elements " + str(pipelinesIds) + " for Pipelines on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				dataSet.pipelines.add(dataPipeline)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePipelines( self, dataSetId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to remove elements " + str(pipelinesIds) + " for Pipelines on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				dataSet.pipelines.remove(dataPipeline)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSemanticModels( self, dataSetId, semanticModelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to add elements " + str(semanticModelsIds) + " for SemanticModels on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = semanticModelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SemanticModel		
				semanticModel = SemanticModelDelegate().get(id).first();	
				# add the SemanticModel
				dataSet.semanticModels.add(semanticModel)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSemanticModels( self, dataSetId, semanticModelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to remove elements " + str(semanticModelsIds) + " for SemanticModels on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = semanticModelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SemanticModel		
				semanticModel = SemanticModelDelegate().get(id).first();	
				# add the SemanticModel
				dataSet.semanticModels.remove(semanticModel)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDimensions( self, dataSetId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to add elements " + str(dimensionsIds) + " for Dimensions on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				dataSet.dimensions.add(dimension)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDimensions( self, dataSetId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to remove elements " + str(dimensionsIds) + " for Dimensions on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				dataSet.dimensions.remove(dimension)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMeasures( self, dataSetId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to add elements " + str(measuresIds) + " for Measures on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				dataSet.measures.add(measure)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMeasures( self, dataSetId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to remove elements " + str(measuresIds) + " for Measures on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				dataSet.measures.remove(measure)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMetrics( self, dataSetId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to add elements " + str(metricsIds) + " for Metrics on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				dataSet.metrics.add(metric)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMetrics( self, dataSetId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to remove elements " + str(metricsIds) + " for Metrics on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				dataSet.metrics.remove(metric)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQualityRules( self, dataSetId, qualityRulesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

		errMsg = "Failed to add elements " + str(qualityRulesIds) + " for QualityRules on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = qualityRulesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the QualityRule		
				qualityRule = QualityRuleDelegate().get(id).first();	
				# add the QualityRule
				dataSet.qualityRules.add(qualityRule)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQualityRules( self, dataSetId, qualityRulesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

		errMsg = "Failed to remove elements " + str(qualityRulesIds) + " for QualityRules on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = qualityRulesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the QualityRule		
				qualityRule = QualityRuleDelegate().get(id).first();	
				# add the QualityRule
				dataSet.qualityRules.remove(qualityRule)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTags( self, dataSetId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to add elements " + str(tagsIds) + " for Tags on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				dataSet.tags.add(tag)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTags( self, dataSetId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to remove elements " + str(tagsIds) + " for Tags on DataSet"

		try:
			# get the DataSet
			dataSet = self.get( dataSetId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				dataSet.tags.remove(tag)
				
			# save it		
			dataSet.save()
			
			# reload and return the appropriate version
			return self.get( dataSetId );
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(dataSetId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
