class RetentionSchedulesController < ApplicationController
  def index
    @retentionSchedules = RetentionSchedule.all
  end
 
  def show
    @retentionSchedule = RetentionSchedule.find(params[:id])
  end
 
  def new
    @retentionSchedule = RetentionSchedule.new
  end
 
  def edit
    @retentionSchedule = RetentionSchedule.find(params[:id])
  end
 
  def create
    @retentionSchedule = RetentionSchedule.new(retentionSchedule_params)
 
    if @retentionSchedule.save
      redirect_to retentionSchedules_path
    else
      render 'new'
    end
  end
 
  def update
    @retentionSchedule = RetentionSchedule.find(params[:id])
 
    if @retentionSchedule.update(retentionSchedule_params)
      redirect_to retentionSchedules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @retentionSchedule = RetentionSchedule.find(params[:id])
    @retentionSchedule.destroy
    redirect_to retentionSchedules_path
  end

 
  private
    def retentionSchedule_params
      params.require(:retentionSchedule).permit(:name, :retentionPeriodMonths, :RetentionTrigger, :DispositionAction, :Status)
    end
end