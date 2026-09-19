class TimesheetsController < ApplicationController
  def index
    @timesheets = Timesheet.all
  end
 
  def show
    @timesheet = Timesheet.find(params[:id])
  end
 
  def new
    @timesheet = Timesheet.new
  end
 
  def edit
    @timesheet = Timesheet.find(params[:id])
  end
 
  def create
    @timesheet = Timesheet.new(timesheet_params)
 
    if @timesheet.save
      redirect_to timesheets_path
    else
      render 'new'
    end
  end
 
  def update
    @timesheet = Timesheet.find(params[:id])
 
    if @timesheet.update(timesheet_params)
      redirect_to timesheets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @timesheet = Timesheet.find(params[:id])
    @timesheet.destroy
    redirect_to timesheets_path
  end

 
  private
    def timesheet_params
      params.require(:timesheet).permit(:periodStart, :periodEnd, :submissionDate, :Status)
    end
end