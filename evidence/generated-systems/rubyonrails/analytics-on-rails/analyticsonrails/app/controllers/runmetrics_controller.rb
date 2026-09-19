class RunMetricsController < ApplicationController
  def index
    @runMetrics = RunMetric.all
  end
 
  def show
    @runMetric = RunMetric.find(params[:id])
  end
 
  def new
    @runMetric = RunMetric.new
  end
 
  def edit
    @runMetric = RunMetric.find(params[:id])
  end
 
  def create
    @runMetric = RunMetric.new(runMetric_params)
 
    if @runMetric.save
      redirect_to runMetrics_path
    else
      render 'new'
    end
  end
 
  def update
    @runMetric = RunMetric.find(params[:id])
 
    if @runMetric.update(runMetric_params)
      redirect_to runMetrics_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @runMetric = RunMetric.find(params[:id])
    @runMetric.destroy
    redirect_to runMetrics_path
  end

 
  private
    def runMetric_params
      params.require(:runMetric).permit(:name, :value)
    end
end