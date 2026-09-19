class ActivitysController < ApplicationController
  def index
    @activitys = Activity.all
  end
 
  def show
    @activity = Activity.find(params[:id])
  end
 
  def new
    @activity = Activity.new
  end
 
  def edit
    @activity = Activity.find(params[:id])
  end
 
  def create
    @activity = Activity.new(activity_params)
 
    if @activity.save
      redirect_to activitys_path
    else
      render 'new'
    end
  end
 
  def update
    @activity = Activity.find(params[:id])
 
    if @activity.update(activity_params)
      redirect_to activitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @activity = Activity.find(params[:id])
    @activity.destroy
    redirect_to activitys_path
  end

 
  private
    def activity_params
      params.require(:activity).permit(:subject, :dueDate, :startAt, :endAt, :location, :ActivityType, :Status, :Priority)
    end
end