class TimeEntrysController < ApplicationController
  def index
    @timeEntrys = TimeEntry.all
  end
 
  def show
    @timeEntry = TimeEntry.find(params[:id])
  end
 
  def new
    @timeEntry = TimeEntry.new
  end
 
  def edit
    @timeEntry = TimeEntry.find(params[:id])
  end
 
  def create
    @timeEntry = TimeEntry.new(timeEntry_params)
 
    if @timeEntry.save
      redirect_to timeEntrys_path
    else
      render 'new'
    end
  end
 
  def update
    @timeEntry = TimeEntry.find(params[:id])
 
    if @timeEntry.update(timeEntry_params)
      redirect_to timeEntrys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @timeEntry = TimeEntry.find(params[:id])
    @timeEntry.destroy
    redirect_to timeEntrys_path
  end

 
  private
    def timeEntry_params
      params.require(:timeEntry).permit(:entryDate, :hoursWorked, :EntryType)
    end
end