class Case_sController < ApplicationController
  def index
    @case_s = Case_.all
  end
 
  def show
    @case_ = Case_.find(params[:id])
  end
 
  def new
    @case_ = Case_.new
  end
 
  def edit
    @case_ = Case_.find(params[:id])
  end
 
  def create
    @case_ = Case_.new(case__params)
 
    if @case_.save
      redirect_to case_s_path
    else
      render 'new'
    end
  end
 
  def update
    @case_ = Case_.find(params[:id])
 
    if @case_.update(case__params)
      redirect_to case_s_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @case_ = Case_.find(params[:id])
    @case_.destroy
    redirect_to case_s_path
  end

 
  private
    def case__params
      params.require(:case_).permit(:caseNumber, :subject, :description, :slaDue, :Status, :Priority, :Origin, :Severity)
    end
end