class TrainingRunsController < ApplicationController
  def index
    @trainingRuns = TrainingRun.all
  end
 
  def show
    @trainingRun = TrainingRun.find(params[:id])
  end
 
  def new
    @trainingRun = TrainingRun.new
  end
 
  def edit
    @trainingRun = TrainingRun.find(params[:id])
  end
 
  def create
    @trainingRun = TrainingRun.new(trainingRun_params)
 
    if @trainingRun.save
      redirect_to trainingRuns_path
    else
      render 'new'
    end
  end
 
  def update
    @trainingRun = TrainingRun.find(params[:id])
 
    if @trainingRun.update(trainingRun_params)
      redirect_to trainingRuns_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @trainingRun = TrainingRun.find(params[:id])
    @trainingRun.destroy
    redirect_to trainingRuns_path
  end

 
  private
    def trainingRun_params
      params.require(:trainingRun).permit(:runLabel, :startedAt, :completedAt, :Status)
    end
end