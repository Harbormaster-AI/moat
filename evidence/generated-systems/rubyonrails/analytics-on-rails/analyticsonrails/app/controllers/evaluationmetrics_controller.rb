class EvaluationMetricsController < ApplicationController
  def index
    @evaluationMetrics = EvaluationMetric.all
  end
 
  def show
    @evaluationMetric = EvaluationMetric.find(params[:id])
  end
 
  def new
    @evaluationMetric = EvaluationMetric.new
  end
 
  def edit
    @evaluationMetric = EvaluationMetric.find(params[:id])
  end
 
  def create
    @evaluationMetric = EvaluationMetric.new(evaluationMetric_params)
 
    if @evaluationMetric.save
      redirect_to evaluationMetrics_path
    else
      render 'new'
    end
  end
 
  def update
    @evaluationMetric = EvaluationMetric.find(params[:id])
 
    if @evaluationMetric.update(evaluationMetric_params)
      redirect_to evaluationMetrics_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @evaluationMetric = EvaluationMetric.find(params[:id])
    @evaluationMetric.destroy
    redirect_to evaluationMetrics_path
  end

 
  private
    def evaluationMetric_params
      params.require(:evaluationMetric).permit(:name, :value)
    end
end